package rsax

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
)

func GenerateKeyHex(bits int) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}

func GenerateKeyBase64(bits int) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}
