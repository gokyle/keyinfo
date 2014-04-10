package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
)

var verbose bool

func main() {
	flVerbose := flag.Bool("v", false, "dump key parameters")
	flag.Parse()
	verbose = *flVerbose

	for _, file := range flag.Args() {
		if err := dumpKey(file); err != nil {
			fmt.Printf("[!] file: %s err: %v\n", file, err)
		}
	}
}

func dumpKey(file string) (err error) {
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	block, _ := pem.Decode(fileData)
	if block == nil {
		err = fmt.Errorf("no valid PEM data found")
		return
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		err = dumpRSAKey(block.Bytes)
	case "PRIVATE KEY":
		err = dumpPKCS8(block.Bytes)
	case "EC PRIVATE KEY":
		err = dumpECKey(block.Bytes)
	default:
		err = fmt.Errorf("unsupported private key")
	}
	return
}

func dumpPKCS8(der []byte) error {
	priv, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return err
	}

	switch priv := priv.(type) {
	case *rsa.PrivateKey:
		printRSAKey(priv)
	case *ecdsa.PrivateKey:
		printECKey(priv)
	default:
		err = fmt.Errorf("unknown key type")
	}
	return err
}

func dumpRSAKey(der []byte) (err error) {
	key, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return
	}
	printRSAKey(key)
	return
}

func printRSAKey(key *rsa.PrivateKey) {
	fmt.Println("[+] Type: RSA")
	fmt.Println("[+] Size:", key.PublicKey.N.BitLen())
	if !verbose {
		return
	}
	fmt.Printf("D: %x\n", key.D.Bytes())
	fmt.Printf("N: %x\n", key.N.Bytes())
	for _, p := range key.Primes {
		fmt.Printf("Prime: %x\n", p.Bytes())
	}
}

func dumpECKey(der []byte) (err error) {
	key, err := x509.ParseECPrivateKey(der)
	if err != nil {
		return
	}
	printECKey(key)
	return
}

func printECKey(key *ecdsa.PrivateKey) {
	fmt.Println("[+] Type: EC")
	fmt.Printf("[+] Curve: ")
	switch key.Curve {
	case elliptic.P256():
		fmt.Println("P256")
	case elliptic.P384():
		fmt.Println("P384")
	case elliptic.P521():
		fmt.Println("P521")
	default:
		fmt.Println("unknown")
	}
}
