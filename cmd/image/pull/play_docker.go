package main

import (
	"fmt"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client" // v1.13.1
	"github.com/spf13/cobra"
)


func RunPull(dockerCli *client.Client) {
	containers, err := dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("运行中的容器ID:")
	for _, container := range containers {
		fmt.Printf("%s \n", container.ID[:10])
	}

}


func main() {
	dockerCli, _ := client.NewEnvClient()

	// 定义ps命令
	var cmdPull = &cobra.Command{
		Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
		Short: "Pull an image or a repository from a registry",
		Run: func(cmd *cobra.Command, args []string)  {
			RunPull(dockerCli)
		},
	}
	// 定义根命令
	var rootCmd = &cobra.Command{Use: "play_docker"}
	// 加入ps命令
	rootCmd.AddCommand(cmdPull)
	// 初始化cobra
	rootCmd.Execute()
}
