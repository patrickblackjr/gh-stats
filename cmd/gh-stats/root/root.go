package root

import (
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/patrickblackjr/gh-stats/cmd/gh-stats/list"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-stats",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-stats.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

func NewRootCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh-stats <command>",
		Short: "Pull GitHub organization stats",
	}
	cmd.AddCommand(list.NewCmdList(f))
	err := doc.GenMarkdownTree(cmd, "./docs")
	if err != nil {
		panic(err)
	}
	return cmd
}
