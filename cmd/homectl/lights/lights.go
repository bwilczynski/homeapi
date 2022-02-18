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

	lightsCmd.PersistentFlags().StringVar(&target, "target", "localhost:50051", "GRPC address")
	lightsCmd.AddCommand(newToggleCmd(client))
	lightsCmd.AddCommand(newListCmd(client))

	return lightsCmd
}
