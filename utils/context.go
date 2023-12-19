package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	// Attempt to retrieve the rawUserID from the context using the key "x-user-id"
	rawUserID, exists := c.Get("x-user-id")
	if !exists {
		return 0, errors.New("user id not found in context")
	}

	// Check if the retrieved rawUserID is a string type
	userIDStr, ok := rawUserID.(string)
	if !ok {
		return 0, errors.New("invalid user id format in context")
	}

	// Parse the userIDStr as an unsigned integer (base 10, 32-bit)
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0, err
	}

	// Convert the userID from uint64 to uint and return it along with no error
	return uint(userID), nil
}
