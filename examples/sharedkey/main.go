package main

import (
	"fmt"

	BIDECDSA "github.com/1Kosmos/gohelperfiles"
)

func main() {
	peerPublicKey := "cmiLaXiIxrEvqDMmZOcer0GU7njlBk7jpwnDjvQbMc/HFOUn3PGWrLpCVchZI/YNpY614J/60cJQkhsL7vm9HA=="
	pv, _, _ := BIDECDSA.GenerateKeyPair()
	sharedKey := BIDECDSA.CreateSharedKey(pv, peerPublicKey)
	fmt.Printf("SharedKey %v", sharedKey)
}
