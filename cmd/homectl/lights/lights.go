package lights

import (
	"github.com/spf13/cobra"
)

func NewLightsCmd() *cobra.Command {
	var target string
	client := &lightServiceClient{target: &target}

	var lightsCmd = &cobra.Command{
		Use:   "lights",
		Short: "Control lights",
	}

	lightsCmd.PersistentFlags().StringVar(&target, "target", "", "GRPC address")

	lightsCmd.AddCommand(newToggleCmd(client))

	return lightsCmd
}
