package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hinkolas/clai/pkg/application"
	"github.com/spf13/cobra"
)

var app *application.App = nil

var rootCmd = &cobra.Command{
	Version: fmt.Sprintf("%s, %s/%s", "0.0.1", runtime.GOOS, runtime.GOARCH),
	Use:     "clai",
	Short:   "clai (short for \"command line ai\") is a simple cli for chatting with your LLMs (and maybe more in the future) written in Go.",
	Run:     Run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	var err error

	app, err = application.NewApp()
	if err != nil {
		fmt.Printf("Failed to initialize app: %v", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}

func Run(cmd *cobra.Command, args []string) {

	fmt.Println("Hello World")

	fmt.Println(app.Path)

}
