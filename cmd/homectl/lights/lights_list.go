package lights

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/bwilczynski/homeapi/lights"
	"github.com/spf13/cobra"
)

func newListCmd(client *lightServiceClient) *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List lights",
		RunE: func(cmd *cobra.Command, args []string) error {
			ll, err := client.List(context.Background())
			if err != nil {
				return err
			}
			print(ll)

			return nil
		},
	}

	return listCmd
}

func print(ll []*lights.Light) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "%s\t%s\n", "ID", "NAME")
	for _, l := range ll {
		fmt.Fprintf(w, "%s\t%s\n", l.Id, l.Name)
	}
}
