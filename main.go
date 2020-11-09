package main

import (
  "os"
  "fmt"
  "github.com/spf13/cobra"
)

func main() {
  var cmdGenWithName = &cobra.Command{
    Use:   "gen [Contest Name] [Extension]",
    Short: "Generate Files",
    Long: `Generate folders and files with the name you entered.`,
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
      info, err := os.Stat(args[0])
      if err != nil {
        panic(err)
      }

      if info == nil{
        fmt.Println(args[0] + " is already exits.")
        return
      }
      err = os.Mkdir(args[0], 0777);
      if err != nil {
        panic(err)
      }
    },
  }

  var rootCmd = &cobra.Command{Use: "acgen"}
  rootCmd.AddCommand(cmdGenWithName)
  rootCmd.Execute()
}