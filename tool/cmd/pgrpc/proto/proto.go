package proto

import (
	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/proto/add"
	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/proto/client"
	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/proto/server"
	"github.com/XM-GO/panda-interface/tool/cmd/pgrpc/proto/service"

	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
	Run:   run,
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(service.CmdService)
	CmdProto.AddCommand(server.CmdServer)
}

func run(cmd *cobra.Command, args []string) {
	// Prompt help information If there is no parameter
	if len(args) == 0 {
		cmd.Help()
		return
	}
}
