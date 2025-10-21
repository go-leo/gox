package rsax

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
)

func DecodeKeyHex(data []byte) (string, string, error) {
	block, _ := pem.Decode(data)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}

func DecodeKeyBase64(data []byte) (string, string, error) {
	block, _ := pem.Decode(data)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", "", err
	}
	privateKeyStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey))
	publicKeyStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	return privateKeyStr, publicKeyStr, nil
}
