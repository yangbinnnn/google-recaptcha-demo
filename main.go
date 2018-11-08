package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yangbinnnn/google-recaptcha-demo/api"

	"github.com/labstack/echo"
)

const (
	version = "0.1"
)

var (
	// Build information should only be set by -ldflags.
	BuildDate    string
	BuildGitHash string

	h bool
	v bool
	e *echo.Echo
)

func cmd() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		fmt.Println("Version:", version)
		if len(BuildDate) > 0 {
			fmt.Println("BuildDate:", BuildDate)
			fmt.Println("BuildGitHash:", BuildGitHash)
		}
		os.Exit(0)
	}

	if h {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	// cmd
	cmd()

	// api
	api.InitApi()

	// this will block forever
	api.StartAPP()
}
