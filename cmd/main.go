package main

import "github.com/pedrolopesme/open-rba/internal/api"

func main() {
	api := api.NewAPI()
	api.Setup()
	api.Run()
}
