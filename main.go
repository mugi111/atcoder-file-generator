package main

import (
  "github.com/spf13/cobra"
)

func main() {
  var cmdGenWithName = &cobra.Command{
    Use:   "Generate Files For [Contest Name]",
    Short: "Generate Files",
    Long: `Generate folders and files with the name you entered.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      
    },
  }

  var rootCmd = &cobra.Command{Use: "acgen"}
  rootCmd.AddCommand(cmdGenWithName)
  rootCmd.Execute()
}