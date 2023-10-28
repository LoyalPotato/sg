package main

import (
	"fmt"
	"os"

	"github.com/LoyalPotato/stacked-guide/src/cmd"
)

//import "github.com/LoyalPotato/stacked-guide/src/cmd"

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
