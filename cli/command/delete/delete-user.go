package delete

import (
	"github.com/spf13/cobra"
)

var delete_userCmd = &cobra.Command{
	Use:   "user USERID",
	Short: "Deletes the speciefied User",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteID(args, "User")
	},
}

func init() {
	deleteCmd.AddCommand(delete_userCmd)
}
