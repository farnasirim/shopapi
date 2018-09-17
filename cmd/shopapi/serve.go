package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	serveAddress string
)

// represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the api on the given address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if serveAddress == "" {
			log.Fatalln("you need to supply the serve address")
		}
		log.Println("TODO: serve on", serveAddress)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&serveAddress, "address", "", "ip:port to listen on")
}
