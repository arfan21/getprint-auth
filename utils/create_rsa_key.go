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

//CreateKey ... typeToken is token or refreshToken
func CreateKey(KEY string, typeToken string) error{
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("failed to generate private key: %s\n", err)
		return err
	}
	pubKey := &privKey.PublicKey

	//create jwks
	realKey, err := jwk.New(privKey)
	if err != nil {
		fmt.Printf("failed to create JWK: %s\n", err)
		return err
	}
	realKey.Set(jwk.KeyIDKey, KEY)
	JwksPubKey, err := jwk.New(pubKey)
	if err != nil {
		fmt.Printf("failed to create JWK: %s\n", err)
		return err
	}

	// Remember, the key must have the proper "kid"
	JwksPubKey.Set(jwk.KeyIDKey, KEY)


	// This key set contains two keys, the first one is the correct one
	keyset := &jwk.Set{Keys: []jwk.Key{JwksPubKey}}
	keySetJSON, _ := json.MarshalIndent(keyset, " ", " ")

	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed find path :%s", err)
		return err
	}
	//check directory /well-knows if not exist
	if _, err := os.Stat(path + "/well-knows"); os.IsNotExist(err) {
		os.Mkdir(path + "/well-knows", 0777)
	}
	//write jwks
	err = ioutil.WriteFile(path + "/well-knows/jwks-"+typeToken+".json", keySetJSON, 0666)
	if err != nil {
		fmt.Printf("failed write jwks file :%s", err)
		return err
	}
	//check directory /certs if not exist
	if _, err := os.Stat(path + "/certs"); os.IsNotExist(err) {
		os.Mkdir(path + "/certs", 0777)
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	privateKeyBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	//write privatepem
	privatePem, err := os.Create(path + "/certs/private-"+typeToken+".pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		return err
	}
	//encode block into privatePem file
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		return err
	}

	pubKeyBytes := x509.MarshalPKCS1PublicKey(pubKey)
	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	//write pubPem
	pubPem, err := os.Create(path + "/certs/public-"+typeToken+".pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		return err
	}
	//encode block into pubPem file
	err = pem.Encode(pubPem, pubKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		return err
	}

	return nil
}