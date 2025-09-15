// 代码生成时间: 2025-09-16 06:56:22
// password_tool.go

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

// EncryptionKey 定义加密密钥
const EncryptionKey = "your_encryption_key"

// encryptData 加密数据
func encryptData(plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(EncryptionKey))
	if err != nil {
		return "", err
	}

	gf := cipher.NewGCM(block)
	nonce := make([]byte, gf.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gf.Seal(nonce, nonce, []byte(plaintext), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// decryptData 解密数据
func decryptData(ciphertext string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(EncryptionKey))
	if err != nil {
		return "", err
	}

	gf, cipher := cipher.NewGCM(block), ciphertextBytes
	nonceSize := gf.NonceSize()
	if len(cipher) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := cipher[:nonceSize], cipher[nonceSize:]
	plaintext, err := gf.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func main() {
	// 测试加密解密
	originalText := "Hello, World!"

	encryptedText, err := encryptData(originalText)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted text: ", encryptedText)

	decryptedText, err := decryptData(encryptedText)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted text: ", decryptedText)
}
