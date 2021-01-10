package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type AES struct {
}

func (a AES) Encrypt(key, data []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encodeBytes := padding(data, aes.BlockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes)+len(iv))
	copy(crypted[:aes.BlockSize], iv)
	blockMode.CryptBlocks(crypted[aes.BlockSize:], encodeBytes)
	return crypted, nil
}

func padding(ciphertext []byte, blockSize int) []byte {
	dataSize := ((len(ciphertext)-1)/blockSize + 1) * blockSize
	if dataSize == len(ciphertext) {
		return ciphertext
	}
	newData := make([]byte, dataSize)
	copy(newData, ciphertext)
	return newData
}

func (a AES) Decrypt(key, data []byte) ([]byte, error) {
	if len(data) < aes.BlockSize {
		return nil, errors.New("data len must greater than 16")
	}
	iv := data[:aes.BlockSize]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(data))
	blockMode.CryptBlocks(result, data[aes.BlockSize:])
	return bytes.Trim(result, "\x00"), nil
}
