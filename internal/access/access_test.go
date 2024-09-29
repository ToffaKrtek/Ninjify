package access

import (
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
  token, err := GenerateAccessToken("user123", time.Minute)
  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }

  if token == "" {
    t.Fatal("Expected a not emptyh token")
  }
}
