package connection

// This is to simplify the connection to a host
// using password or pubkey

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	
	"code.google.com/p/go.crypto/ssh"
)


// RESOURCES:
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


func PubKeyClientConfig(usr, pathToKey string) (*ssh.ClientConfig) {
	
	var (
		k *keychain
		err error
	)
	
	if pathToKey == "" {
		pathToKey = filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa")
	} 
	
	k = &keychain{}
	
	// Add path to id_rsa file
	if err = k.loadPEM(pathToKey); err != nil {
		log.Fatal("Cannot load key: " + err.Error())
	}
	
	return &ssh.ClientConfig{
		// Change to your username
		User: usr,
		Auth: []ssh.AuthMethod{
			ssh.ClientAuthKeyring(k),
		},
	}
}


func PWClientConfig(usr, pw string) (*ssh.ClientConfig) {
	return &ssh.ClientConfig{
		User: usr,
		Auth: []ssh.AuthMethod{
			ssh.ClientAuthPassword(password(pw)),
		},
	}
}

// All of the following is taken from 
// http://kiyor.us/2013/12/29/golang-ssh-example/
// Thank you!!!:)

func strip(v string) string {
	return strings.TrimSpace(strings.Trim(v, "\n"))
}
 
type keychain struct {
	keys []ssh.Signer
}
 
func (k *keychain) Key(i int) (ssh.PublicKey, error) {
	if i < 0 || i >= len(k.keys) {
		return nil, nil
	}
	return k.keys[i].PublicKey(), nil
}
 
func (k *keychain) Sign(i int, rand io.Reader, data []byte) (sig []byte, err error) {
	return k.keys[i].Sign(rand, data)
}
 
func (k *keychain) add(key ssh.Signer) {
	k.keys = append(k.keys, key)
}
 
func (k *keychain) loadPEM(file string) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return err
	}
	k.add(key)
	return nil
}



// password implements the ClientPassword interface
type password string

func (p password) Password(user string) (string, error) {
	return string(p), nil
}



