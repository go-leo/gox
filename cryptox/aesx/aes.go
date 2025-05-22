package aesx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Cipher 加密解密接口
type Cipher interface {
	// Encrypt 加密
	Encrypt(plaintext, key []byte) ([]byte, error)

	// Decrypt 解密
	Decrypt(ciphertext, key []byte) ([]byte, error)
}

func ECB() Cipher {
	return ecb{}
}

func CBC() Cipher {
	return cbc{}
}

func CFB() Cipher {
	return cfb{}
}

func OFB() Cipher {
	return ofb{}
}

func CTR() Cipher {
	return ctr{}
}

type ecb struct {
}

// Encrypt 加密
func (ecb) Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充明文
	plaintext = padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	// 分块加密
	for i := 0; i < len(plaintext); i += block.BlockSize() {
		block.Encrypt(ciphertext[i:i+block.BlockSize()], plaintext[i:i+block.BlockSize()])
	}
	return ciphertext, nil
}

// Decrypt 解密
func (ecb) Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))
	// 分块解密
	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(plaintext[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
	}
	// 去除填充
	plaintext, err = unPadding(plaintext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

type cbc struct{}

// Encrypt CBC模式加密
func (cbc) Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = padding(plaintext, blockSize)
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintext)
	return ciphertext, nil
}

// Decrypt CBC模式解密
func (cbc) Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext, err = unPadding(plaintext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

type cfb struct{}

// Encrypt CFB模式加密
func (cfb) Encrypt(plaintext, key []byte) ([]byte, error) {
	return baseStream{newStream: cipher.NewCFBEncrypter}.Encrypt(plaintext, key)
}

// Decrypt CFB模式解密
func (cfb) Decrypt(ciphertext, key []byte) ([]byte, error) {
	return baseStream{newStream: cipher.NewCFBDecrypter}.Decrypt(ciphertext, key)
}

type ofb struct {
}

// Encrypt OFB模式加密
func (ofb) Encrypt(plaintext, key []byte) ([]byte, error) {
	return baseStream{newStream: cipher.NewOFB}.Encrypt(plaintext, key)
}

// Decrypt OFB模式解密
func (ofb) Decrypt(ciphertext, key []byte) ([]byte, error) {
	return baseStream{newStream: cipher.NewOFB}.Decrypt(ciphertext, key)
}

type ctr struct{}

// Encrypt OFB模式加密
func (ctr) Encrypt(plaintext, key []byte) ([]byte, error) {
	return baseStream{newStream: cipher.NewCTR}.Encrypt(plaintext, key)
}

// Decrypt OFB模式解密
func (ctr) Decrypt(ciphertext, key []byte) ([]byte, error) {
	return baseStream{newStream: cipher.NewCTR}.Decrypt(ciphertext, key)
}

type baseStream struct {
	newStream func(block cipher.Block, iv []byte) cipher.Stream
}

// Encrypt OFB模式加密
func (b baseStream) Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := b.newStream(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// Decrypt OFB模式解密
func (b baseStream) Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))
	stream := b.newStream(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// padding 填充数据
func padding(data []byte, blockSize int) []byte {
	size := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(size)}, size)...)
}

// unPadding 去除填充
func unPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	size := int(data[length-1])
	if size > length {
		return nil, fmt.Errorf("invalid size")
	}
	return data[:length-size], nil
}
