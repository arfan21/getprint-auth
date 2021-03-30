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

	privKey, err = ioutil.ReadFile(path + "/certs/private-"+tokenType+".pem")
	if err != nil {
		fmt.Printf("failed read private key :%s", err)
		return nil, nil, err
	}

	pubKey, err = ioutil.ReadFile(path + "/certs/public-"+tokenType+".pem")
	if err != nil {
		fmt.Printf("failed read public key :%s", err)
		return nil, nil, err
	}
	return privKey, pubKey, nil
}
