package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xehap/jago/config"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter client ID: ")
	clientID, _ := reader.ReadString('\n')

	fmt.Print("Enter client secret: ")
	clientSecret, _ := reader.ReadString('\n')

	cfg := config.NewConfig(clientID, clientSecret)

	log.Print("Your Client ID is:" + cfg.APIKey + " and your Client Secret is " + cfg.APISecret)
}
