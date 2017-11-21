package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zoumo/mamba"
)

func main() {
	cmd := mamba.NewCommand(&cobra.Command{
		Use:  "example",
		Long: "this is an cli example",
		Run: func(cmd *cobra.Command, args []string) {
			f := cmd.PersistentFlags().Lookup("log")
			fmt.Println("hello world")
			fmt.Printf("l: %v\n", f)
			fmt.Printf("viper.log: %v", viper.GetString("log"))
		},
	})

	fs := []mamba.Flag{
		mamba.BoolFlag{
			Name:      "log",
			Shorthand: "l",
			// the EnvKey will override
			EnvKey:              "LOG",
			Persistent:          true,
			Deprecated:          "move",
			ShorthandDeprecated: "move2",
			Hidden:              true,
		},
	}

	// set env
	os.Setenv("LOG", "test")

	mamba.AutomaticEnv()
	mamba.SetEnvPrefix("AAA")
	cmd.AddFlag(fs...)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
