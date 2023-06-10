package runner

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xehap/jago/config"
)

func RunMain() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter client ID: ")
	clientID, _ := reader.ReadString('\n')

	fmt.Print("Enter client secret: ")
	clientSecret, _ := reader.ReadString('\n')

	cfg := config.NewConfig(clientID, clientSecret)

	log.Printf("Your Client ID is: %s and your Client Secret is %s", cfg.APIKey, cfg.APISecret)
}
