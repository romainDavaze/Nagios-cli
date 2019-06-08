package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generates zsh completion scripts",
	Long: `To load completion run

. <(nagiosxi-cli completion zsh)

To configure your zsh shell to load completions for each session, add to your zshrc

# ~/.zshrc
. <(nagiosxi-cli completion zsh)
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(completionZshCmd)
}
