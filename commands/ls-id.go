package commands

import (
	"fmt"
	"strings"

	"github.com/MichaelMure/git-bug/cache"
	"github.com/spf13/cobra"
)

func runLsID(cmd *cobra.Command, args []string) error {

	var backend *cache.RepoCache

	prefix := args[0]

	for _, id := range backend.AllBugsIds() {
		if prefix == "" || strings.HasPrefix(id, prefix) {
			fmt.Println(id)
		}
	}

	return nil
}

var listBugIDCmd = &cobra.Command{
	Use:     "ls-id [<prefix>]",
	Short:   "List Bug Id",
	PreRunE: loadRepo,
	RunE:    runLsID,
}

func init() {
	RootCmd.AddCommand(listBugIDCmd)
}
