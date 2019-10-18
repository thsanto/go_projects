package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "gifm",
		Short: "Hello go",
	}

	cmd.AddCommand(printTimeCmd())

	cmd.AddCommand(listDir())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey Gophers1 Time is", prettyTime)
			return nil
		},
	}
}

func listDir() *cobra.Command {
	return &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			files, err := ioutil.ReadDir("./")
			if err != nil {
				fmt.Println("erro")
			}

			for _, f := range files {
				fmt.Println(f.Name())
			}

			return nil
		},
	}
}
