package cli

import (
	ishell "github.com/pikpakidm/pikpakidm/internal/shell"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start an interactive PikPak shell",
	Run: func(cmd *cobra.Command, args []string) {
		ishell.Start(rootCmd)
	},
}
