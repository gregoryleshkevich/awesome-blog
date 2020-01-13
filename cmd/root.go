package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "awesome-blog",
	Short: "This should locate a short description",
	Long: "This should locate a long description",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

// Execute asd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
