package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"os"
)

func CreateKey(KEY string)[]byte{

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("failed to generate private key: %s\n", err)
		return nil
	}

	realKey, err := jwk.New(privKey)
	if err != nil {
		fmt.Printf("failed to create JWK: %s\n", err)
		return nil
	}
	realKey.Set(jwk.KeyIDKey, KEY)
	pubKey, err := jwk.New(privKey.PublicKey)
	if err != nil {
		fmt.Printf("failed to create JWK: %s\n", err)
		return nil
	}

	// Remember, the key must have the proper "kid"
	pubKey.Set(jwk.KeyIDKey, KEY)


	// This key set contains two keys, the first one is the correct one
	keyset := &jwk.Set{Keys: []jwk.Key{pubKey}}
	keySetJSON, _ := json.MarshalIndent(keyset, " ", " ")

	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed find path :%s", err)
		return nil
	}

	err = ioutil.WriteFile(path + "/oauth/jwks.json", keySetJSON, 0644)
	if err != nil {
		fmt.Printf("failed write file :%s", err)
		return nil
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	privateKeyBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	return pem.EncodeToMemory(privateKeyBlock)
}