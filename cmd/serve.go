package cmd

import (
	"github.com/spf13/cobra"
	"github.com/supersupersimple/comment/app/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartWebServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("enable_https", "https", true, "enable https and auto tls")
	// serveCmd.Flags().IntP("port", "p", 8080, "port when use http mode")
}
