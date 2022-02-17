package lights

import (
	"context"

	"github.com/spf13/cobra"
)

func newToggleCmd(client *lightServiceClient) *cobra.Command {
	var toggleCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggles light state",
		RunE: func(cmd *cobra.Command, args []string) error {
			gid, _ := cmd.Flags().GetString("groupId")

			return client.ToggleGroup(context.Background(), gid)
		},
	}

	toggleCmd.Flags().String("groupId", "", "Group ID to toggle")
	toggleCmd.MarkFlagRequired("groupId")

	return toggleCmd
}
