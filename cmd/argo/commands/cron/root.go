package cron

import (
	"github.com/spf13/cobra"
)

func NewCronWorkflowCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "cron",
		Short: "manage cron workflows",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewListCommand())
	command.AddCommand(NewCreateCommand())
	command.AddCommand(NewDeleteCommand())
	command.AddCommand(NewLintCommand())
	command.AddCommand(NewSuspendCommand())
	command.AddCommand(NewResumeCommand())

	return command
}
