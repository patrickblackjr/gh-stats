package root

import (
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/patrickblackjr/gh-stats/cmd/gh-stats/list"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// NewRootCmd creates the root command for the CLI
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
