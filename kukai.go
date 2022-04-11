package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"strings"
)

func decrypt_v2orV3(cipherData string, password string, salt []byte) ([]byte, error) {
	if len(password) == 0 || len(salt) == 0 {
		return nil, fmt.Errorf("must supply password or salt to decrypt")
	}
	split := strings.Split(cipherData, "==")
	if len(split) != 2 {
		return nil, fmt.Errorf("cipher not in correct format")
	}
	ciphertext, tag := split[0], split[1]

	ctb, err := hex.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	tb, err := hex.DecodeString(tag)
	if err != nil {
		return nil, err
	}
	// The tag is expected at the end of the ciphertext for authentication
	ctb = append(ctb, tb...)
	key, err := scrypt.Key([]byte(password), salt, 65536, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}
	decoded, err := gcm.Open(nil, salt, ctb, nil)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}
