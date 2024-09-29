package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"crypto/rand"
)

var (
  ErrEmptyText = errors.New("Empty text to encrypt")
)

func generateKey(pass string) []byte {
  hash := sha256.Sum256([]byte(pass))
  return hash[:]
}

func Encrypt(text []byte, pass string) (string, error) {
  if len(text) == 0 {
    return "", ErrEmptyText
  }

  key := generateKey(pass)
  block, err := aes.NewCipher(key)
  if err != nil {
    return "", err
  }

  iv := make([]byte, aes.BlockSize)
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    return "", err
  }
  cipherText := make([]byte, len(text))
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText, text)

    return hex.EncodeToString(iv) + hex.EncodeToString(cipherText), nil
}

func Decrypt(text string, pass string) ([]byte, error) {
   key := generateKey(pass)
    cipherText, _ := hex.DecodeString(string(text))
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]

    plainText := make([]byte, len(cipherText))
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(plainText, cipherText)

    return plainText, nil
}
