package main

import (
	"fmt"
	"os"
    "flag"
	BIDECDSA "github.com/1Kosmos/gohelperfiles"
)

const CLEAR_TEXT = "Clear text to encrypt."

func main() {
	var clearText string
	flag.StringVar(&clearText, "text", CLEAR_TEXT, "clear text to test Ecdsahelper with")
	localPrivateKey, localPublicKey, err := BIDECDSA.GenerateKeyPair()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper error: %v\n", err)
		os.Exit(1)
	}
	remotePrivateKey, remotePublicKey, err := BIDECDSA.GenerateKeyPair()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper error: %v\n", err)
		os.Exit(1)
	}
	encSharedKey := BIDECDSA.CreateSharedKey(localPrivateKey, remotePublicKey)
	encryptedText, err := BIDECDSA.EcdsaHelper(BIDECDSA.ENCRYPT, CLEAR_TEXT, encSharedKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s error: %v\n", BIDECDSA.ENCRYPT, err)
		os.Exit(1)
	}
	decSharedKey := BIDECDSA.CreateSharedKey(remotePrivateKey, localPublicKey)
	decryptedText, err := BIDECDSA.EcdsaHelper(BIDECDSA.DECRYPT, encryptedText, decSharedKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s error: %v\n", BIDECDSA.DECRYPT, err)
		os.Exit(1)
	}
	if decryptedText == CLEAR_TEXT {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s == %s\n", decryptedText, CLEAR_TEXT)
	} else {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s != %s\n", decryptedText, CLEAR_TEXT)
	}
}
