package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type Rsa struct {
	// rsa 公钥
	priKey *rsa.PrivateKey

	// rsa 私钥
	pubKey *rsa.PublicKey
}

// NewRsa get a new instance of cryptor
func NewRsa(pub, priv string) (*Rsa, error) {
	if len(pub) == 0 && len(priv) == 0 {
		return nil, errors.New("pub and priv cannot both be empty")
	}

	e := &Rsa{}

	if len(pub) > 0 {
		block, _ := pem.Decode([]byte(pub))
		if block == nil {
			return nil, errors.New("public key error")
		}

		pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		e.pubKey = pubInterface.(*rsa.PublicKey)
	}

	if len(priv) > 0 {
		block, _ := pem.Decode([]byte(priv))
		if block == nil {
			return nil, errors.New("private key error")
		}
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		e.priKey = privateKey
	}

	return e, nil
}

// rsaEncrypt rsa encrypt.
func (e *Rsa) rsaEncrypt(text []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, e.pubKey, text)
}

// rsaDecrypt rsa decrypt.
func (e *Rsa) rsaDecrypt(text []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, e.priKey, text)
}

// Encrypt encrypt and base64 encoding
func (e *Rsa) Encrypt(plantText string) string {
	if plantText == "" {
		return plantText
	}

	cipherText, err := e.rsaEncrypt([]byte(plantText))
	if err != nil {
		return plantText
	}
	return base64.StdEncoding.EncodeToString(cipherText)
}

// Decrypt decrypt and base64 decode
func (e *Rsa) Decrypt(cipherText string) string {
	if cipherText == "" {
		return cipherText
	}

	cipher, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return cipherText
	}

	plantText, err := e.rsaDecrypt(cipher)
	if err != nil {
		return cipherText
	}

	return string(plantText)
}
