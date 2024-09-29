package access

import (
	"errors"
	"time"
)

var (
  ErrNotValidUserId = errors.New("UserID is not valid")
)

func GenerateAccessToken(userId string, duration time.Duration) (string, error) {
  if userId == "" {
    return "", ErrNotValidUserId
  }
  return "token-for-" + userId, nil
}
