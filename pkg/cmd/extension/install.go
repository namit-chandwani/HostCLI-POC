package extension

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install extensions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
}
