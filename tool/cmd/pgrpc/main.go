package main

import (
	"log"

	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/markdown"
	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/project"
	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/proto"
	"github.com/XM-GO/panda-interface/tool/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "panda-tool",
	Short:   "An elegant toolkit for Go microservices.",
	Long:    `An elegant toolkit for Go microservices.`,
	Version: version.Version,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(markdown.CmdMarkdown)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
