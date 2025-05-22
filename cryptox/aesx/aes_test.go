package aesx

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	// 示例密钥 (16, 24, 或 32 字节分别对应 -128, -192, 或 AES-256)
	key := []byte("0123456789abcdef0123456789abcdef") // 256-bit key
	plaintext := []byte("Hello, AES encryption in Golang!")

	fmt.Printf("原始数据: %s\n", plaintext)

	// ECB模式
	ecbCiphertext, err := ECB().Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("ECB加密错误:", err)
		return
	}
	fmt.Printf("ECB加密后: %s\n", hex.EncodeToString(ecbCiphertext))
	ecbDecrypted, err := ECB().Decrypt(ecbCiphertext, key)
	if err != nil {
		fmt.Println("ECB解密错误:", err)
		return
	}
	fmt.Printf("ECB解密后: %s\n", ecbDecrypted)

	// CBC模式
	cbcCiphertext, err := CBC().Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("CBC加密错误:", err)
		return
	}
	fmt.Printf("CBC加密后: %s\n", hex.EncodeToString(cbcCiphertext))
	cbcDecrypted, err := CBC().Decrypt(cbcCiphertext, key)
	if err != nil {
		fmt.Println("CBC解密错误:", err)
		return
	}
	fmt.Printf("CBC解密后: %s\n", cbcDecrypted)

	// CFB模式
	cfbCiphertext, err := CFB().Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("CFB加密错误:", err)
		return
	}
	fmt.Printf("CFB加密后: %s\n", hex.EncodeToString(cfbCiphertext))
	cfbDecrypted, err := CFB().Decrypt(cfbCiphertext, key)
	if err != nil {
		fmt.Println("CFB解密错误:", err)
		return
	}
	fmt.Printf("CFB解密后: %s\n", cfbDecrypted)

	// OFB模式
	ofbCiphertext, err := OFB().Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("OFB加密错误:", err)
		return
	}
	fmt.Printf("OFB加密后: %s\n", hex.EncodeToString(ofbCiphertext))
	ofbDecrypted, err := OFB().Decrypt(ofbCiphertext, key)
	if err != nil {
		fmt.Println("OFB解密错误:", err)
		return
	}
	fmt.Printf("OFB解密后: %s\n", ofbDecrypted)
}
