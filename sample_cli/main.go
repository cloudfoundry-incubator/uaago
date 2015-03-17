package main

import (
	"os"

	"github.com/cloudfoundry/uaago"
)

func main() {
	uaa := uaago.NewClient(os.Args[1])
	token, err := uaa.GetAuthToken(os.Args[2], os.Args[3])
	if err != nil {
		panic(err.Error())
	}

	println("TOKEN:", token)
}
