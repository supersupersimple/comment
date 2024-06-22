package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/supersupersimple/comment/app/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run: func(cmd *cobra.Command, args []string) {
		https, err := cmd.Flags().GetBool("https")
		if err != nil {
			log.Fatal(err)
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}

		server.StartWebServer(https, host, port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.
	serveCmd.Flags().BoolP("https", "", true, "enable https and auto tls")
	serveCmd.Flags().StringP("host", "", "", "host")
	serveCmd.Flags().IntP("port", "", 8080, "port when use http mode")
}
