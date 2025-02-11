package guest

import (
	"github.com/spf13/cobra"
)

var guest_lxcCmd = &cobra.Command{
	Use:   "lxc GUESTID NODEID",
	Short: "Creates a new Guest System of the type Lxc on the specified Node",
	Long: `Creates a new Guest System of the type Lxc on the specified Node.
The config can be set with the --file flag or piped from stdin.
For config examples see "example guest lxc"`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return createGuest(args, "LxcGuest")
	},
}

func init() {
	guestCmd.AddCommand(guest_lxcCmd)
}
