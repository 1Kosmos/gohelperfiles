/**
 * Copyright (c) 2018, 1Kosmos Inc. All rights reserved.
 * Licensed under 1Kosmos Open Source Public License version 1.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of this license at
 *    https://github.com/1Kosmos/1Kosmos_License/blob/main/LICENSE.txt
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/1Kosmos/gohelperfiles"
)

const CLEAR_TEXT = "Clear text to encrypt."

func main() {
	var clearText string
	flag.StringVar(&clearText, "text", CLEAR_TEXT, "clear text to test")
	flag.Parse()
	localPrivateKey, localPublicKey, err := gohelperfiles.GenerateKeyPair()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper error: %v\n", err)
		os.Exit(1)
	}
	remotePrivateKey, remotePublicKey, err := gohelperfiles.GenerateKeyPair()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper error: %v\n", err)
		os.Exit(1)
	}
	encSharedKey := gohelperfiles.CreateSharedKey(localPrivateKey, remotePublicKey)
	encryptedText, err := gohelperfiles.EcdsaHelper(gohelperfiles.ENCRYPT, clearText, encSharedKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s error: %v\n", gohelperfiles.ENCRYPT, err)
		os.Exit(1)
	}
	decSharedKey := gohelperfiles.CreateSharedKey(remotePrivateKey, localPublicKey)
	decryptedText, err := gohelperfiles.EcdsaHelper(gohelperfiles.DECRYPT, encryptedText, decSharedKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s error: %v\n", gohelperfiles.DECRYPT, err)
		os.Exit(1)
	}
	if decryptedText == clearText {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s == %s\n", decryptedText, clearText)
	} else {
		fmt.Fprintf(os.Stderr, "Ecdsahelper %s != %s\n", decryptedText, clearText)
	}
}
