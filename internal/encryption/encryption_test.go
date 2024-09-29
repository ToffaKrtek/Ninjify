package encryption

import "testing"

func TestEncryptDecrypt(t *testing.T) {
  pass := "mysecretpassword"
  originalText := []byte("Test text")

  encrypted, err := Encrypt(originalText, pass)
  if err != nil {
    t.Fatalf("Error in encrypt expected nil, got: %v", err)
  }
  decrypted, err := Decrypt(encrypted, pass)
    if err != nil {
        t.Fatalf("Failed to decrypt: %v", err)
    }

    if string(decrypted) != string(originalText) {
        t.Errorf("Decrypted text does not match original: got %s, want %s", decrypted, originalText)
    }
}
