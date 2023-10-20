package cobra

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// this assumes the binary created is the same name as the base command
// So `./cobra command` would be the same as running `go run main.go command`
func DoExample() {
	baseCommand := baseCommand()
	baseCommand.Execute()
}

func baseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cobra",
		Short: "Base Cobra command",
		Run:   runCobraBaseCommand,
	}
	cmd.AddCommand(newCobraCommand())
	return cmd
}

func runCobraBaseCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("You ran the cobra base command.  Try running some of the underlying commands to see more functionality\n")
}

func newCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command",
		Short: "Sample command",
		Run:   runCobraCommand,
	}
	cmd.Flags().Uint("version", 0, "optional version")
	return cmd
}

func runCobraCommand(cmd *cobra.Command, args []string) {
	version, err := cmd.Flags().GetUint("version")
	if err != nil {
		log.Fatalf("invalid version specified")
	}

	fmt.Printf("ran cobra command, insert logic here.  Version: %d\n", version)
}
