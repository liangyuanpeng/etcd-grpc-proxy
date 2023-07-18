package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.etcd.io/etcd/proxy/pkg/server"
)

var (
	rootCmd = &cobra.Command{
		Use:        "etcd",
		Short:      "etcd grpc-proxy",
		SuggestFor: []string{"etcd"},
	}
)

func main() {
	rootCmd.AddCommand(server.NewGRPCProxyCommand())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
