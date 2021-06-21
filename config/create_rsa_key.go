package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lestrrat-go/jwx/jwk"
)

//CreateKey ... tokenType is token or refreshToken
func CreateKey(KEY string, tokenType string) error {
	bitSize := 1024

	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}
	// Extract public component.
	pub := key.Public()

	//create jwks
	jwksKey, err := jwk.New(key)
	if err != nil {
		fmt.Printf("failed to create JWK: %s\n", err)
		return err
	}
	jwksKey.Set(jwk.KeyIDKey, KEY)
	jwksPubKey, err := jwk.New(pub)
	if err != nil {
		fmt.Printf("failed to create JWK: %s\n", err)
		return err
	}

	// Remember, the key must have the proper "kid"
	jwksPubKey.Set(jwk.KeyIDKey, KEY)

	// This key set contains two keys, the first one is the correct one
	jwksPubKeySet := &jwk.Set{Keys: []jwk.Key{jwksPubKey}}
	jwksPubKeySetJSON, _ := json.MarshalIndent(jwksPubKeySet, " ", " ")

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyBytes, _ := x509.MarshalPKCS8PrivateKey(key)
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: keyBytes,
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubBytes, _ := x509.MarshalPKIXPublicKey(pub.(*rsa.PublicKey))
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		},
	)

	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed find path :%s", err)
		return err
	}

	//check directory /certs if not exist
	if _, err := os.Stat(path + "/certs"); os.IsNotExist(err) {
		os.Mkdir(path+"/certs", 0777)
	}

	//check directory /well-knows if not exist
	if _, err := os.Stat(path + "/well-knows"); os.IsNotExist(err) {
		os.Mkdir(path+"/well-knows", 0777)
	}

	jwksFilename := path + "/well-knows/jwks-" + tokenType + ".json"
	privateKeyFilename := path + "/certs/" + tokenType + ".rsa"
	publicKeyFilename := path + "/certs/" + tokenType + ".rsa.pub"

	if _, err := os.Stat(jwksFilename); os.IsNotExist(err) {
		//write jwks
		if err := ioutil.WriteFile(path+"/well-knows/jwks-"+tokenType+".json", jwksPubKeySetJSON, 0666); err != nil {
			fmt.Printf("failed write jwks file :%s", err)
			return err
		}
	}

	if _, err := os.Stat(privateKeyFilename); os.IsNotExist(err) {
		// Write private key to file.
		if err := ioutil.WriteFile(path+"/certs/"+tokenType+".rsa", keyPEM, 0700); err != nil {
			fmt.Printf("failed write private key :%s", err)
			return err
		}
	}

	if _, err := os.Stat(publicKeyFilename); os.IsNotExist(err) {
		// Write public key to file.
		if err := ioutil.WriteFile(path+"/certs/"+tokenType+".rsa.pub", pubPEM, 0755); err != nil {
			fmt.Printf("failed write public key :%s", err)
			return err
		}
	}

	return nil
}
