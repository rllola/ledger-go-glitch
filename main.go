package main

import (
		"fmt"
		//"time"
		"encoding/hex"
		"github.com/zondax/ledger-go"
	)

func main() {
	ledgerAdmin := ledger_go.NewLedgerAdmin()
	ledgerAdmin.ListDevices()
	
	ledger, err := ledgerAdmin.Connect(0)
	if err != nil {
		fmt.Printf("Couldn't connect to ledger\n")
	}
	
	message := []byte{0xE0, 0x01, 0, 0, 0}
	
	response, err := ledger.Exchange(message)
	if err != nil {
		fmt.Printf("Error while exchanging message\n")
	}
	
	fmt.Printf("Response: %x\n", response)

	// Jusqu'ici tout va bien
	messageFirstPart, err := hex.DecodeString("05020000142c000080da010080000000800000008000000080")
	if err != nil {
		fmt.Printf("Couldn't decode message\n")
	}

	response, err = ledger.Exchange(messageFirstPart)
	if err != nil {
		fmt.Printf("Error while exchanging message: %s\n", err.Error())
	}
		
	fmt.Printf("Response: %x\n", response)
	
	messageSecondPart, err := hex.DecodeString("05020200a3636f617369732d636f72652f636f6e73656e7375733a20747820666f7220636861696e2033353565316431653334316532626466363164663838336661353864393037663262653361666338366162333039326132363161396166373362316133353639a463666565a26367617319099c66616d6f756e744064626f6479a166616d6f756e74417b656e6f6e636507666d6574686f646c7374616b696e672e4275726e")
	if err != nil {
		fmt.Printf("Couldn't decode message\n")
	}

	response, err = ledger.Exchange(messageSecondPart)
	if err != nil {
		fmt.Printf("Error while exchanging message: %s\n", err.Error())
	}
			
	fmt.Printf("Response: %x\n", response)

	// Introducing delay
	//time.Sleep(2 * time.Second)

	// On recommence !
	response, err = ledger.Exchange(messageFirstPart)
	if err != nil {
		fmt.Printf("Error while exchanging message\n")
	}
			
	fmt.Printf("Response: %x\n", response)
	
	response, err = ledger.Exchange(messageSecondPart)
	if err != nil {
		fmt.Printf("Error while exchanging message: %s\n", err.Error())
	}
			
	fmt.Printf("Response: %x\n", response)

}
