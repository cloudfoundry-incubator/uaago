package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-incubator/uaago"
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args[1:]) != 3 {
		fmt.Fprintf(os.Stderr, "Usage %s [URL] [CLIENT_ID] [EXISTING_REFRESH_TOKEN]\n", args[0])
		return 1
	}

	uaa, err := uaago.NewClient(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create client: %s\n", err.Error())
		return 1
	}

	refreshToken, accessToken, err := uaa.GetRefreshToken(args[2], args[3], true)
	if err != nil {
		if refreshToken == "" {
			fmt.Fprintf(os.Stderr, "Failed to get new refresh token: %s\n", err.Error())
		}
		if accessToken == "" {
			fmt.Fprintf(os.Stderr, "Failed to get access token: %s\n", err.Error())
		}
		return 1
	}

	fmt.Fprintf(os.Stdout, "REFRESH_TOKEN: %s\n", refreshToken)
	fmt.Fprintf(os.Stdout, "ACCESS_TOKEN: %s\n", accessToken)
	return 0
}
