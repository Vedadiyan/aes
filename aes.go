package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

type (
	AES struct {
		encryptor cipher.BlockMode
		decryptor cipher.BlockMode
	}
)

func New(key []byte, iv []byte) (*AES, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encryptor := cipher.NewCBCEncrypter(block, iv)
	decryptor := cipher.NewCBCDecrypter(block, iv)
	aes := AES{
		encryptor: encryptor,
		decryptor: decryptor,
	}
	return &aes, nil
}
func (aes *AES) Encrypt(bytes []byte) []byte {
	initLen := len(bytes)
	remainder := 256 - (initLen % 256)
	if remainder == 0 {
		remainder = 256
	}
	len := initLen + (remainder)
	dest := make([]byte, len)
	copy := append(bytes, dest[initLen+1:]...)
	copy = append(copy, byte(remainder-1))
	aes.encryptor.CryptBlocks(dest, copy)
	return dest
}

func (aes *AES) Decrypt(bytes []byte) []byte {
	dest := make([]byte, len(bytes))
	aes.decryptor.CryptBlocks(dest, bytes)
	remainder := int(dest[len(dest)-1])
	return dest[:len(dest)-remainder-1]
}
