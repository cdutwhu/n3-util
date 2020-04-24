package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"regexp"
)

var (
	RExpMD5    = regexp.MustCompile("\"[A-Fa-f0-9]{32}\"")
	RExpSHA1   = regexp.MustCompile("\"[A-Fa-f0-9]{40}\"")
	RExpSHA256 = regexp.MustCompile("\"[A-Fa-f0-9]{64}\"")
)

// SHA1Str :
func SHA1Str(s string) string {
	return fSf("%x", sha1.Sum([]byte(s)))
}

// SHA256Str :
func SHA256Str(s string) string {
	return fSf("%x", sha256.Sum256([]byte(s)))
}

// MD5Str :
func MD5Str(s string) string {
	return fSf("%x", md5.Sum([]byte(s)))
}

// Encrypt :
func Encrypt(data []byte, password string) []byte {
	if password == "" {
		return data
	}
	block, _ := aes.NewCipher([]byte(MD5Str(password)))
	gcm, err := cipher.NewGCM(block)
	FailOnErr("%v", err)
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	FailOnErr("%v", err)
	return gcm.Seal(nonce, nonce, data, nil)
}

// Decrypt :
func Decrypt(data []byte, password string) ([]byte, error) {
	if password == "" {
		return data, nil
	}
	key := []byte(MD5Str(password))
	block, err := aes.NewCipher(key)
	FailOnErr("%v", err)
	gcm, err := cipher.NewGCM(block)
	FailOnErr("%v", err)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	// FailOnErr("%v", err)
	return plaintext, err
}
