package repos

import (
	"fmt"
	"log"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
	"github.com/spf13/cobra"
)

// NewCmdListRepos creates a new command called repos which is a subcommand of list
func NewCmdListRepos(f *cmdutil.Factory) *cobra.Command {
	var org string
	cmd := &cobra.Command{
		Use:   "repos",
		Short: "List all the repos in the current org",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := api.DefaultGraphQLClient()
			if err != nil {
				log.Fatal(err)
			}

			var query struct {
				Organization struct {
					Repositories struct {
						Nodes []struct {
							NameWithOwner string
						}
						TotalCount int
						PageInfo   struct {
							StartCursor string
							EndCursor   string
							HasNextPage bool
						} `graphql:"pageInfo"`
					} `graphql:"repositories(first: $first)"`
					Name string
				} `graphql:"organization(login: $org)"`
			}
			variables := map[string]interface{}{
				"first": graphql.Int(100),
				"org":   graphql.String(org),
			}
			err = client.Query("Organization", &query, variables)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Found %d repositories in the organization %s\n",
				query.Organization.Repositories.TotalCount, query.Organization.Name)
			fmt.Printf("Page information: %+v", query.Organization.Repositories.PageInfo)
		},
	}
	cmd.Flags().StringVarP(&org, "org", "o", "github", "the organization to pull stats for")
	return cmd
}
