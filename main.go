package main

import (
	"log"

	"github.com/joho/godotenv"
	cmd "github.com/kerokerogeorge/gacha-dapp.git/cmd"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	cmd.DeployGachaContract()
}
