# gohelperfiles

## Prerequisities

- go version ```1.20```
- to assure getting the latest gohelper module version it is recommended to set the GOPROXY variable to ```direct``` only. Example
```go env -u GOPROXY; go env -w GOPROXY="direct"``` <br>
To reset back to default:
```go env -u GOPROXY;  go env -w GOPROXY="https://proxy.golang.org,direct"```

## Usage

- see ```examples/ecdsahelper/main.go``` for example using ecdsahelper function:
```
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
```
For usage of Encrypt/Decrypt see tests:

```
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
```
## Running tests

- run
```
go test
PASS
ok  	github.com/1Kosmos/gohelperfiles	0.012s

``` 
- to see individual results us `-v`

```
go test -v 
=== RUN   Test_Keypair
    BIDECDSA_test.go:25: Test_Keypair q10WnF32y2ViAbHJfv7kGwU+6vinvZPFOqerTQeyygw= vU4Z0bJomXZlhQ97q+8PQQIqwMAsQI/wjWmeG/aLbpIn+fjiVtvHvdq6FzOsNFVgblyVvT2H4R+/y0KjiL3hAA==
--- PASS: Test_Keypair (0.01s)
=== RUN   Test_Sharedkey
    BIDECDSA_test.go:44: Test_Sharedkey [253 164 53 156 41 123 108 237 220 112 104 237 229 235 213 134 104 25 50 87 105 209 146 113 117 88 237 34 199 249 36 0]
--- PASS: Test_Sharedkey (0.00s)
=== RUN   Test_EncryptDecrypt
    BIDECDSA_test.go:81: Test_EncryptDecrypt decrypted text equal to original text
--- PASS: Test_EncryptDecrypt (0.00s)
=== RUN   Test_Ecsahelper
    BIDECDSA_test.go:118: Test_Ecsahelper decrypted text equal to original text
--- PASS: Test_Ecsahelper (0.00s)
PASS
ok  	github.com/1Kosmos/gohelperfiles	0.016s
```