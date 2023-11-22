package list

import (
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/patrickblackjr/gh-stats/cmd/gh-stats/list/repos"
	"github.com/spf13/cobra"
)

// NewCmdList creates a new command called list
func NewCmdList(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <repos|prs|commits>",
		Short:   "Pull GitHub organization stats",
		Aliases: []string{"ls"},
	}
	cmd.AddCommand(repos.NewCmdListRepos(f))
	return cmd
}
