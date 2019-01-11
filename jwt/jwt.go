package jwt

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

var (
	privateKey *ecdsa.PrivateKey
)

/*--------------------------------- Create New JWT Token for ios APN ---------------------------------*/
func CreateJWT(privateFilePath string,keyId string, teamId string) (string, error) {
	if privateKey == nil {
		newPrivateKey, err := readPrivateKey(privateFilePath)
		if err != nil {
			return "", err
		}
		privateKey = newPrivateKey
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"iat": time.Now().Unix(),
			"iss": teamId,
		})
	token.Header["alg"] = "ES256"
	token.Header["kid"] = keyId
	s, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return s, nil
}

/*--------------------------------- Read private key from provided config.json ---------------------------------*/
func readPrivateKey(filepath string) (*ecdsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(file)

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if privateKey, ok := key.(*ecdsa.PrivateKey); ok {
		return privateKey, nil
	}
	return nil, errors.New("")
}
