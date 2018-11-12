package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

var version = "undefined"

func main() {
	app := cli.NewApp()
	app.Name = "n0cli"
	app.Version = version
	// app.Usage = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-endpoint",
			Value:  "localhost:20180",
			EnvVar: "N0CLI_API_ENDPOINT",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "get",
			Usage:     "Get resource if set resource name, List resources if not set",
			ArgsUsage: "[resource type] [resource name]",
			Action:    Get,
		},
		{
			Name:      "delete",
			Usage:     "Delete resource",
			ArgsUsage: "[resource type] [resource name]",
			Action:    Delete,
		},
		{
			Name:  "do",
			Usage: "Do DAG tasks (Detail n0stack/pkg/dag)",
			Description: `
	## File format

	---
	task_name:
		type: Network
		action: GetNetwork
		args:
			name: test-network
		depend_on:
			- dependency_task_name
		ignore_error: true
	dependency_task_name:
		type: ...
	---

	- task_name
			- 任意の名前をつけ、ひとつのリクエストに対してユニークなものにする
	- type
			- gRPC メッセージを指定する
			- VirtualMachine や virtual_machine という形で指定できる
	- action
			- gRPC の RPC を指定する
			- GetNetwork など定義のとおりに書く
	- args
			- gRPC の RPCのリクエストを書く
	- depend_on
			- DAG スケジューリングに用いられる
			- task_name を指定する
	- ignore_error
	    - タスクでエラーが発生しても継続する
			`,
			ArgsUsage: "[file name]",
			Action:    Do,
		},
	}

	log.SetFlags(log.Lshortfile)
	log.SetOutput(ioutil.Discard)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to command: %v", err.Error())
		os.Exit(1)
	}
}
