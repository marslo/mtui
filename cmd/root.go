package cmd

import (
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "mtui",
  Short: "A beautiful terminal UI input tool",
  Long:  "mtui is a small CLI utility to collect user input, confirmations, and messages in the terminal with style.",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    os.Exit(1)
  }
}

// vim:tabstop=2:softtabstop=2:shiftwidth=2:expandtab:filetype=go:
