package connection

// This is to simplify the connection to a host
// using password or pubkey

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	
	"code.google.com/p/go.crypto/ssh"
)


// RESOURCES:
// * https://docs.google.com/document/d/1nF2wlkIwuA4AXryOvE2p0hgQUbsyRYklKSot4ahH3Aw/edit#
// * https://godoc.org/code.google.com/p/go.crypto/ssh
// * https://godoc.org/code.google.com/p/go.crypto/ssh
// * http://dave.cheney.net/tag/golang-3/page/2
// * http://play.golang.org/p/3z513UKrOY
// * https://code.google.com/p/go/issues/detail?id=7787
// * https://groups.google.com/forum/#!topic/golang-nuts/QRsZSPqwDhM
// * http://play.golang.org/p/kMhHvbl4SG
// * https://docs.google.com/document/d/1nF2wlkIwuA4AXryOvE2p0hgQUbsyRYklKSot4ahH3Aw/edit
// * http://godoc.org/code.google.com/p/go.crypto/ssh/agent
// * https://groups.google.com/d/topic/golang-nuts/0JfeQ-Qu37U/discussion
// ---
// * http://kiyor.us/2013/12/29/golang-ssh-example/
// ---


	
	
func SshSessionWithKey (server, usr, pathToKey string) (*ssh.Session, error) {
	var (
		config *ssh.ClientConfig 
		client *ssh.Client
		err error
	)
	
	if config, err = PubKeyClientConfig(usr, pathToKey); err != nil {
		log.Println("Can'r create config for connection: ", err)
		return &ssh.Session{}, err
	}
	
	log.Println("Try to connect to ", server)
	if client, err = ssh.Dial("tcp", server, config); err != nil {
		log.Println("Failed to dial: " + err.Error())
		return &ssh.Session{}, err
	}
	
	log.Println("Start new session")
	return client.NewSession()
}
	
func SshSessionWithPw (server, usr, pw string) (*ssh.Session, error) {
	var (
		config *ssh.ClientConfig 
		client *ssh.Client
		err error
	)
	
	if config, err = PwClientConfig(usr, pw); err != nil {
		log.Println("Can'r create config for connection: ", err)
		return &ssh.Session{}, err
	}
	
	log.Println("Try to connect to ", server)
	if client, err = ssh.Dial("tcp", server, config); err != nil {
		log.Println("Failed to dial: " + err.Error())
		return &ssh.Session{}, err
	}
	
	log.Println("Start new session")
	return client.NewSession()
}	
	
	

func PubKeyClientConfig(usr, pathToKey string) (*ssh.ClientConfig, error) {
	
	var (
		buf []byte
		key ssh.Signer
		err error
	)
	
	// Check for default key location
	if pathToKey == "" {
		pathToKey = filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa")
	} 
	
	// Load key
	if buf, err = ioutil.ReadFile(pathToKey); err != nil {
		log.Printf("Can't read PEM key file in %v, error: %v\n", pathToKey, err)
		return &ssh.ClientConfig{}, err
	}
	
	if key, err = ssh.ParsePrivateKey(buf); err != nil {
		log.Printf("Can't parse PEM key buffer to load key, error: %v\n", err)
		return &ssh.ClientConfig{}, err
	}	
	
	return &ssh.ClientConfig{
		// Change to your username
		User: usr,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}, nil
}


func PwClientConfig(usr, pw string) (*ssh.ClientConfig, error) {
	return &ssh.ClientConfig{
		User: usr,
		Auth: []ssh.AuthMethod{
// 			ssh.Password(password(pw)),
			ssh.Password(pw),
		},
	}, nil
}
