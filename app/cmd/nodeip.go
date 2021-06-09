package cmd

import (
	"fmt"

	"github.com/longhorn/longhorn-instance-manager/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func GetIP() cli.Command {
	return cli.Command{
		Name: "match",
		Action: func(c *cli.Context) {
			if err := match(c); err != nil {
				logrus.Fatalf("Error running match command: %v.", err)
			}
		},
	}
}

func match(c *cli.Context) error {
	url := c.GlobalString("url")
	netCard := c.GlobalString("interface")
	cli := client.NewProcessManagerClient(url)
	ip, err := cli.DataIPGet(netCard)
	if err != nil {
		return err
	}
	fmt.Println(*ip)
	return nil
}
