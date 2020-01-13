package cmd

import (
	"fmt"
	"awesome-blog/routes"
	"github.com/spf13/cobra"
	"net/http"
	"github.com/gorilla/mux"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "server",
	Short: "Run http web server",
	Run: func(cmd *cobra.Command, args []string) {

		r := *mux.NewRouter()
		routes := routes.Routes{R: &r}
		routes.InitRoutes()

		http.Handle("/", &r)

		fmt.Println("Server is listening...")
		http.ListenAndServe(":8181", nil)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
