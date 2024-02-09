package config

import (
	"slices"
	"strings"

	"github.com/git-town/git-town/v12/src/cli/flags"
	"github.com/git-town/git-town/v12/src/cmd/cmdhelpers"
	"github.com/git-town/git-town/v12/src/execute"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

const removeConfigDesc = "Removes the Git Town configuration"

func removeConfigCommand() *cobra.Command {
	addVerboseFlag, readVerboseFlag := flags.Verbose()
	cmd := cobra.Command{
		Use:   "remove",
		Args:  cobra.NoArgs,
		Short: removeConfigDesc,
		Long:  cmdhelpers.Long(removeConfigDesc),
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeRemoveConfig(readVerboseFlag(cmd))
		},
	}
	addVerboseFlag(&cmd)
	return &cmd
}

func executeRemoveConfig(verbose bool) error {
	repo, err := execute.OpenRepo(execute.OpenRepoArgs{
		Verbose:          verbose,
		DryRun:           false,
		OmitBranchNames:  true,
		PrintCommands:    true,
		ValidateIsOnline: false,
		ValidateGitRepo:  true,
	})
	if err != nil {
		return err
	}
	err = repo.Runner.GitConfig.RemoveLocalGitConfiguration(repo.Runner.Lineage)
	if err != nil {
		return err
	}
	aliasNames := maps.Keys(repo.Runner.Aliases)
	slices.Sort(aliasNames)
	for _, aliasName := range aliasNames {
		if strings.HasPrefix(repo.Runner.Aliases[aliasName], "town ") {
			err = repo.Runner.Frontend.RemoveGitAlias(aliasName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
