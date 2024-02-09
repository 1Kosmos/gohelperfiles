/**
 * Copyright (c) 2018, 1Kosmos Inc. All rights reserved.
 * Licensed under 1Kosmos Open Source Public License version 1.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of this license at
 *    https://github.com/1Kosmos/1Kosmos_License/blob/main/LICENSE.txt
 */

package gohelperfiles

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

const (
	NoPadding Padding = iota
)

const (
	GCM CipherMode = iota
)

type CipherMode int
type Padding int

const ENCRYPT = "encrypt"
const DECRYPT = "decrypt"

type AES struct {
	CipherMode CipherMode
	Padding    Padding
}

func EcdsaHelper(method, text string, key []byte) (string, error) {
	if method == ENCRYPT {
		return Decrypt(text, key)
	} else if method == DECRYPT {
		return Encrypt(text, key)
	} else {
		return "", fmt.Errorf("EcdsaHelper invalid method %s", method)
	}
}

func Encrypt(plainText string, key []byte) (string, error) {
	aes, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	return encryptGcm(aes, plainText)
}

func Decrypt(cipherText string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	aes, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	return decryptGcm(aes, data)
}

func GenerateKeyPair() (string, string, error) {
	privKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}
	pubKeyBytes := privKey.PubKey().SerializeUncompressed()
	dBytes := privKey.ToECDSA().D.Bytes()
	return base64.StdEncoding.EncodeToString(dBytes),
		base64.StdEncoding.EncodeToString(pubKeyBytes[1:]), err
}

func CreateSharedKey(privKeyStr string, pubKeyStr string) []byte {
	pubBytes, _ := base64.StdEncoding.DecodeString(pubKeyStr)
	//make into uncompressed format
	tot := make([]byte, len(pubBytes)+1)
	tot[0] = 0x04
	copy(tot[1:], pubBytes)
	pubKey, err := secp256k1.ParsePubKey(tot)
	if err != nil {
		log.Printf("CreateSharedKey error: %v %v", err, len(pubBytes))
	}
	privBytes, _ := base64.StdEncoding.DecodeString(privKeyStr)
	privKey := secp256k1.PrivKeyFromBytes(privBytes)
	return secp256k1.GenerateSharedSecret(privKey, pubKey)
}

func encryptGcm(aes cipher.Block, plainText string) (string, error) {
	gcm, err := cipher.NewGCMWithNonceSize(aes, 16)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	plainTextBytes := []byte(plainText)
	cipherText := gcm.Seal(nil, nonce, plainTextBytes, nil)
	return packCipherData(cipherText, nonce, gcm.Overhead()), nil
}

func decryptGcm(aes cipher.Block, encrypted []byte) (string, error) {
	aesgcm, err := cipher.NewGCMWithNonceSize(aes, 16)
	if err != nil {
		return "", err
	}
	encryptedBytes, nonce := unpackCipherData(encrypted, aesgcm.NonceSize())
	decryptedBytes, err := aesgcm.Open(nil, nonce, encryptedBytes, nil)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes[:]), nil
}

func packCipherData(cipherText []byte, iv []byte, tagSize int) string {
	ivLen := len(iv)
	data := make([]byte, len(cipherText)+ivLen)
	copy(data[:], iv[0:ivLen])
	copy(data[ivLen:], cipherText)
	return base64.StdEncoding.EncodeToString(data)
}

func unpackCipherData(data []byte, ivSize int) ([]byte, []byte) {
	iv, encryptedBytes := data[:ivSize], data[ivSize:]
	return encryptedBytes, iv
}
