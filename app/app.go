package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the cli app for the application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Go CLI"
	app.Usage = "Search IPs and Hostnames"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  host,
			Value: "google.com.br",
		},
	}
	

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP",
			Flags: flags,
			Action: searchIP,
		},
		{
			Name:  "server",
			Usage: "Search Server Name",
			Flags: flags,
			Action: searchServerName,
		},
	}

	return app
}

func searchIP(c *cli.Context) {
	host := c.String(host)
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServerName(c *cli.Context) {
	host := c.String(host)
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range servers {
		fmt.Println(s.Host)
	}
}
