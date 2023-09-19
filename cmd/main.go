package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	return &cobra.Command{
		Use:   "jtos",
		Short: "jtos is a json to struct generator",
		Long:  "jtos is a json to struct generator",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello world")
		},

		// Uncomment the following line if your bare application
	}
}

func main() {
	root := Root()
	root.Execute()
}
