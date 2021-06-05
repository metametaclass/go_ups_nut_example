package main

import (
	"fmt"
	"os"

	nut "github.com/metametaclass/go.nut"
)

func main() {
	client, connectErr := nut.Connect(os.Args[1])
	if connectErr != nil {
		fmt.Println(connectErr)
		os.Exit(1)
	}
	defer func() {
		ok, err := client.Disconnect()
		if err != nil {
			fmt.Printf("disconnect error %s\n", err)
		} else if !ok {
			fmt.Printf("invalid LOGOUT answer")
		}
	}()
	_, authenticationError := client.Authenticate("admin", "admin")
	if authenticationError != nil {
		fmt.Printf("Authenticate: error %s\n", authenticationError)
		os.Exit(1)
	}
	upsList, listErr := client.GetUPSList()
	if listErr != nil {
		fmt.Printf("GetUPSList: error %s\n", listErr)
		os.Exit(1)
	}
	for _, u := range upsList {
		fmt.Printf("UPS:%s\n", u.Name)
		for _, v := range u.Variables {
			fmt.Printf("\t%s: %v\n", v.Name, v.Value)
		}
	}
}
