/**
 * Copyright (c) 2018, 1Kosmos Inc. All rights reserved.
 * Licensed under 1Kosmos Open Source Public License version 1.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of this license at
 *    https://github.com/1Kosmos/1Kosmos_License/blob/main/LICENSE.txt
 */

package gohelperfiles

import (
	"testing"
)

const CLEAR_TEXT = "Clear text to encrypt/decrypt."
const DECRYPT_FAIL = "decrypted text not equal to original text"
const DECRYPT_SUCCESS = "decrypted text equal to original text"

func Test_Keypair(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_Keypair = %q, %v, want "", error`,
			"Test_Keypair failed", err)
	}
	t.Logf("Test_Keypair %v %v", privateKey, publicKey)
}

func Test_Sharedkey(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_Sharedkey = %q, %v, want "", error`,
			"Test_Sharedkey failed", err)
	}
	_, publicKey1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_Sharedkey = %q, %v, want "", error`,
			"Test_Sharedkey failed", err)
	}
	encSharedKey := CreateSharedKey(privateKey, publicKey1)
	if err != nil {
		t.Fatalf(`Test_Sharedkey = %q, %v, want "", error`,
			"Test_Sharedkey failed", err)
	}
	t.Logf("Test_Sharedkey %v", encSharedKey)
}

func Test_EncryptDecrypt(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_EncryptDecrypt = %q, %v, want "", error`,
			"Test_EncryptDecrypt failed", err)
	}
	privateKey1, publicKey1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_EncryptDecrypt = %q, %v, want "", error`,
			"Test_EncryptDecrypt failed", err)
	}
	sharedKey := CreateSharedKey(privateKey, publicKey1)
	if err != nil {
		t.Fatalf(`Test_EncryptDecrypt = %q, %v, want "", error`,
			"Test_EncryptDecrypt failed", err)
	}
	encryptedText, err := Encrypt(CLEAR_TEXT, sharedKey)
	if err != nil {
		t.Fatalf(`Test_EncryptDecrypt = %q, %v, want "", error`,
			"Test_EncryptDecrypt failed", err)
	}
	sharedKeyDecrypt := CreateSharedKey(privateKey1, publicKey)
	if err != nil {
		t.Fatalf(`Test_EncryptDecrypt = %q, %v, want "", error`,
			"Test_EncryptDecrypt failed", err)
	}
	decryptedText, err := Decrypt(encryptedText, sharedKeyDecrypt)
	if err != nil {
		t.Fatalf(`Test_EncryptDecrypt = %q, %v, want "", error`,
			"Test_EncryptDecrypt failed", err)
	}
	if decryptedText != CLEAR_TEXT {
		t.Fatalf(`Test_EncryptDecrypt = %s "`, DECRYPT_FAIL)
	}
	t.Logf("Test_EncryptDecrypt %s", DECRYPT_SUCCESS)
}

func Test_Ecsahelper(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_Ecsahelper = %q, %v, want "", error`,
			"Test_Ecsahelper failed", err)
	}
	privateKey1, publicKey1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf(`Test_Ecsahelper = %q, %v, want "", error`,
			"Test_Ecsahelper failed", err)
	}
	sharedKey := CreateSharedKey(privateKey, publicKey1)
	if err != nil {
		t.Fatalf(`Test_Ecsahelper = %q, %v, want "", error`,
			"Test_Ecsahelper failed", err)
	}
	encryptedText, err := EcdsaHelper(ENCRYPT, CLEAR_TEXT, sharedKey)
	if err != nil {
		t.Fatalf(`Test_Ecsahelper = %q, %v, want "", error`,
			"Test_Ecsahelper failed", err)
	}
	sharedKeyDecrypt := CreateSharedKey(privateKey1, publicKey)
	if err != nil {
		t.Fatalf(`Test_Ecsahelper = %q, %v, want "", error`,
			"Test_Ecsahelper failed", err)
	}
	decryptedText, err := EcdsaHelper(DECRYPT, encryptedText, sharedKeyDecrypt)
	if err != nil {
		t.Fatalf(`Test_Ecsahelper = %q, %v, want "", error`,
			"Test_Ecsahelper failed", err)
	}
	if decryptedText != CLEAR_TEXT {
		t.Fatalf(`Test_Ecsahelper = %s "`, DECRYPT_FAIL)
	}
	t.Logf("Test_Ecsahelper %s", DECRYPT_SUCCESS)
}
