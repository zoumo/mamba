package main

import (
	goflag "flag"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/zoumo/mamba/cmd/flag-gen/generators"

	"github.com/golang/glog"
	"k8s.io/gengo/args"
)

func main() {
	arguments := &args.GeneratorArgs{
		OutputBase:       args.DefaultSourceTree(),
		GoHeaderFilePath: filepath.Join(args.DefaultSourceTree(), "github.com/zoumo/mamba/boilerplate/boilerplate.go.txt"),
	}
	arguments.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	_ = pflag.Set("logtostderr", "true")
	pflag.Parse()
	_ = goflag.CommandLine.Parse(nil)

	// Run it
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
		os.Exit(1)
	}

	glog.V(2).Info("Completed successfully.")
}
