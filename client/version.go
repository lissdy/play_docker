package main

import (
	"github.com/docker/docker/client"
	"fmt"
	"context"
)

func main() {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	versions, _ := cli.ServerVersion(context.Background())

	fmt.Printf("server versions is %s \n",versions.GoVersion)
}
