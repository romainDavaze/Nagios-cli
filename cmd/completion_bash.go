package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(nagiosxi-cli completion bash)

To configure your bash shell to load completions for each session, add to your bashrc

# ~/.bashrc or ~/.profile
. <(nagiosxi-cli completion bash)
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(completionBashCmd)
}
