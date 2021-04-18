package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadKey(tokenType string) (privKey, pubKey []byte, err error){
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed find path :%s", err)
		return nil,nil,err
	}

	privKey, err = ioutil.ReadFile(path + "/certs/" + tokenType+".rsa")
	//privKey, err = ioutil.ReadFile(path + "/certs/id_rsa")
	if err != nil {
		fmt.Printf("failed read private key :%s", err)
		return nil, nil, err
	}

	pubKey, err = ioutil.ReadFile(path + "/certs/" +tokenType+".rsa.pub")
	//pubKey, err = ioutil.ReadFile(path + "/certs/id_rsa.pub")
	if err != nil {
		fmt.Printf("failed read public key :%s", err)
		return nil, nil, err
	}

	//privKeyDecoded, _ := pem.Decode(privKey)
	//pubKeyDecoded, _ := pem.Decode(pubKey)
	return privKey, pubKey, nil
}
