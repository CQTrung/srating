package services

import (
	"context"
	"time"

	"srating/domain"
	"srating/utils"
)

type locationService struct {
	locationRepository domain.LocationRepository
	contextTimeout       time.Duration
}

func NewLocationService(locationRepository domain.LocationRepository, timeout time.Duration) domain.LocationService {
	return &locationService{
		locationRepository: locationRepository,
		contextTimeout:       timeout,
	}
}

func (u *locationService) CreateLocation(c context.Context, input *domain.Location) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.locationRepository.CreateLocation(ctx, input); err != nil {
		utils.LogError(err, "Failed to create location")
		return err
	}
	return nil
}

func (u *locationService) GetAllLocation(c context.Context, input domain.GetAllLocationRequest) (int64, int64, []*domain.Location, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if input.Limit < 0 {
		input.Limit = 10
	}
	if input.Page < 0 {
		input.Page = 1
	}
	total, totalCount, locations, err := u.locationRepository.GetAllLocation(ctx, input)
	if err != nil {
		utils.LogError(err, "Failed to get all location")
		return 0, 0, nil, err
	}
	return total, totalCount, locations, nil
}

func (u *locationService) GetLocationDetail(c context.Context, id uint) (*domain.Location, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	location, err := u.locationRepository.GetLocationDetail(ctx, id)
	if err != nil {
		utils.LogError(err, "Failed to get location detail")
		return nil, err
	}
	return location, nil
}

func (u *locationService) UpdateLocation(c context.Context, input *domain.Location) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := utils.Validate(input); err != nil {
		utils.LogError(err, "Failed to validate input")
		return err
	}
	if err := u.locationRepository.UpdateLocation(ctx, input); err != nil {
		utils.LogError(err, "Failed to update location")
		return err
	}
	return nil
}

func (u *locationService) DeleteLocation(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	if err := u.locationRepository.DeleteLocation(ctx, id); err != nil {
		utils.LogError(err, "Failed to delete location")
		return err
	}
	return nil
}
