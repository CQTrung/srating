package services

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"srating/domain"
	"srating/utils"
)

type mediaService struct {
	mediaRepository domain.MediaRepository
	contextTimeout  time.Duration
}

func NewMediaService(mediaRepository domain.MediaRepository, timeout time.Duration) domain.MediaService {
	return &mediaService{
		mediaRepository: mediaRepository,
		contextTimeout:  timeout,
	}
}

func (u *mediaService) Upload(ctx context.Context, input []*domain.UploadFileInput) ([]*domain.Media, error) {
	// Attempt to create the "assets" directory if it doesn't exist
	if err := os.MkdirAll("assets", os.ModePerm); err != nil {
		utils.LogError(err, "Error creating assets directory")
		return nil, err
	}

	// Create a slice to store information about uploaded media
	medias := make([]*domain.Media, len(input))

	// Use a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Iterate over each media input for concurrent processing
	for i, uploadFile := range input {
		// Validate the input media file
		if err := utils.Validate(uploadFile); err != nil {
			utils.LogError(err, "Error validating input")
			return nil, err
		}

		// Increment the wait group counter for the current goroutine
		wg.Add(1)

		// Launch a goroutine to process the current media input
		go func(fileIndex int, file *domain.UploadFileInput) {
			// Decrement the wait group counter when the goroutine completes
			defer wg.Done()

			// Open the source file for reading
			fileHeader := file.FileHeader
			src, err := fileHeader.Open()
			if err != nil {
				utils.LogError(err, "Error opening file")
				return
			}
			defer src.Close()

			// Generate a sanitized filename and construct the destination path
			filename := utils.SanitizeFilename(fileHeader.Filename)
			if !utils.IsImage(filename) {
				utils.LogError(err, "Invalid image format")
				return
			}
			dst := filepath.Join("assets", filename)

			// Create necessary directories for the destination path
			if err = os.MkdirAll(filepath.Dir(dst), 0o750); err != nil {
				utils.LogError(err, "Error creating directory")
				return
			}

			// Create the destination file and copy the contents
			out, err := os.Create(filepath.Clean(dst))
			if err != nil {
				utils.LogError(err, "Error creating destination file")
				return
			}
			defer out.Close()

			if _, err := io.Copy(out, src); err != nil {
				utils.LogError(err, "Error copying file")
				return
			}

			// Record media information and upload to repository
			medias[fileIndex] = &domain.Media{
				FileName: filename,
				URL:      "assets/" + filename,
			}
			if err := u.mediaRepository.Upload(ctx, medias[fileIndex]); err != nil {
				utils.LogError(err, "Error uploading file")
				medias[fileIndex] = nil
				return
			}
		}(i, uploadFile)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Return information about uploaded media and potential errors
	return medias, nil
}

func (u *mediaService) GetAll(ctx context.Context) ([]*domain.Media, error) {
	medias, err := u.mediaRepository.GetAll(ctx)
	if err != nil {
		utils.LogError(err, "Error getting all media")
		return nil, err
	}
	return medias, nil
}
