// Package commands contains the CLI commands
package commands

import (
	"fmt"
	"os"

	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/repository"
	"github.com/spf13/cobra"
)

const rootCommandName = "git-bug"

// package scoped var to hold the repo after the PreRun execution
var repo repository.ClockedRepo

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   rootCommandName,
	Short: "A bug tracker embedded in Git",
	Long: `git-bug is a bug tracker embedded in git.

git-bug use git objects to store the bug tracking separated from the files
history. As bugs are regular git objects, they can be pushed and pulled from/to
the same git remote your are already using to collaborate with other peoples.

`,

	// For the root command, force the execution of the PreRun
	// even if we just display the help. This is to make sure that we check
	// the repository and give the user early feedback.
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
	},

	DisableAutoGenTag: true,

	// Custom bash code to connect the git completion for "git bug" to the
	// git-bug completion for "git-bug"
	BashCompletionFunction: `
_git_bug() {
    __start_git-bug "$@"
}
`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func loadRepo(cmd *cobra.Command, args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Unable to get the current working directory: %q\n", err)
	}

	repo, err = repository.NewGitRepo(cwd, bug.Witnesser)
	if err == repository.ErrNotARepo {
		return fmt.Errorf("%s must be run from within a git repo.\n", rootCommandName)
	}

	if err != nil {
		return err
	}

	return nil
}
