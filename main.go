package main

import (
	"fmt"
	"os"

	"github.com/sakuraji-labs/sakuraji_log4j/backend"
	"github.com/urfave/cli"
)

func main() {
	log4j := cli.NewApp()
	log4j.Version = "1.0"
	log4j.Usage = "Discover and remediate Log4Shell vulnerability [CVE-2021-45105]"

	log4j.Commands = []cli.Command{
		{
			Name:  "discover",
			Usage: "Scan OS for Log4Shell vulnerability.",
			Action: func(console *cli.Context) {
				fmt.Println("Discovering Log4Shell vulnerability [CVE-2021-45105]")
				backend.Go()
			},
		},
	}
	log4j.Run(os.Args)
}
