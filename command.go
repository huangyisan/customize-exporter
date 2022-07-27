package main

import (
	"fmt"
	"os"
)

func showAppInfo() {
	fmt.Printf("Version:\t %s\n", Version)
	fmt.Printf("Build:\t %s\n", Build)

}

func command() {
	if version {
		showAppInfo()
		os.Exit(0)
	}
}
