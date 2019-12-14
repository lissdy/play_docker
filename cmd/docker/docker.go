// https://github.com/docker/cli/blob/v19.03.5/cmd/docker/docker.go
package main

import (
	"github.com/docker/cli/cli/command"
	"fmt"
	"os"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/docker/cli/cli"
	"github.com/spf13/pflag"
)

func newDockerCommand(dockerCli *command.DockerCli) *cli.TopLevelCommand{
	var (
		opts    *cliflags.ClientOptions
		flags   *pflag.FlagSet
	)

	cmd := &cobra.Command{
		Use:              "docker [OPTIONS] COMMAND [ARG...]",
		Short:            "A self-sufficient runtime for containers",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return command.ShowHelp(dockerCli.Err())(cmd, args)
			}
			return fmt.Errorf("docker: '%s' is not a docker command.\nSee 'docker --help'", args[0])

		},
		DisableFlagsInUseLine: true,
	}

	opts, flags, _ = cli.SetupRootCommand(cmd)
	flags.BoolP("version", "v", false, "Print version information and quit")
	return cli.NewTopLevelCommand(cmd, dockerCli, opts, flags)
}


func runDocker(dockerCli *command.DockerCli) error {
	tcmd := newDockerCommand(dockerCli)

	cmd, _, err := tcmd.HandleGlobalFlags()
	if err != nil {
		return err
	}

	if err := tcmd.Initialize(); err != nil {
		return err
	}

	return cmd.Execute()
}

func main() {
	dockerCli, err := command.NewDockerCli()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	logrus.SetOutput(dockerCli.Err())
	runDocker(dockerCli)
}
