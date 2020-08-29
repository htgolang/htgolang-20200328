package cmds

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCommand = &cobra.Command{
	Use:   "cmdb",
	Short: "cmdb program",
	Long:  "cmdb programe",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb")
		return nil
	},
}

func init() {
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
