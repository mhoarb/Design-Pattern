package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := NewWaalletFacade("abc", 20)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 20, 100)
	if err != nil {
		log.Fatalf("Error : %s\n", err.Error())
	}
	fmt.Println()

	err = walletFacade.deductMoneyfromWallet("abc", 20, 1002)
	if err != nil {
		log.Fatalf("error : %s\n", err.Error())
	}

	fmt.Println("done")
}
