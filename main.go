package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var config string
	var localFile string
	var tagName string

	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install xray",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Destination: &config,
				},
				&cli.StringFlag{
					Name:        "local",
					Aliases:     []string{"l"},
					Destination: &localFile,
				},
				&cli.StringFlag{
					Name:        "tag",
					Aliases:     []string{"t"},
					Usage:       "set xray tag name",
					Destination: &tagName,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				err = downloadBinaryFile(localFile, tagName)
				if err != nil {
					return err
				}
				err = installBinaryFile()
				if err != nil {
					return err
				}
				err = installConfig(config)
				if err != nil {
					return err
				}
				return installService()
			},
		},
		{
			Name:  "uninstall",
			Usage: "Remove config,cache and uninstall xray",
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				err = uninstallService()
				if err != nil {
					return err
				}
				return uninstallBinaryFile()
			},
		},
		{
			Name:  "update",
			Usage: "Update xray",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Destination: &config,
				},
				&cli.StringFlag{
					Name:        "local",
					Aliases:     []string{"l"},
					Destination: &localFile,
				},
				&cli.StringFlag{
					Name:        "tag",
					Aliases:     []string{"t"},
					Usage:       "set x tag name",
					Destination: &tagName,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				err = updateBinaryFile(localFile, tagName)
				if err != nil {
					return err
				}
				if config != "" {
					err = installConfig(config)
					if err != nil {
						return err
					}
				}
				return updateService()
			},
		},
		{
			Name:  "reload",
			Usage: "Reload service",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Destination: &config,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				if config != "" {
					err = installConfig(config)
					if err != nil {
						return err
					}
				}
				return reloadService()
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s\n", cmd.Root().Version)
	}

	cmd := &cli.Command{
		Usage:    "xray quick install tool",
		Version:  "v3.10",
		Commands: cmds,
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
