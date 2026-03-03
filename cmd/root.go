package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "labcheck",
    Short: "A CLI for checking your homelab health",
    Long:  "labcheck connects to your homelab over SSH and reports on Docker, disk, and SnapRAID status.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}