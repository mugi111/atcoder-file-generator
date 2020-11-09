package main

import (
  "os"
  "fmt"
  "path/filepath"
  "github.com/spf13/cobra"
)

func main() {
  var probList = []string {"A", "B", "C", "D", "E", "F"}
  var cmdGenWithName = &cobra.Command{
    Use:   "gen [Contest Name] [Extension]",
    Short: "Generate Files",
    Long: `Generate folders and files with the name you entered.`,
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
      _, err := os.Stat(args[0])
      if err == nil {
          fmt.Println(args[0] + " is already exists.")
          return
      }

      err = os.Mkdir(args[0], 0777);
      if err != nil {
        panic(err)
      }

      for _, v := range probList {
        genFile, err := os.Create(filepath.Join("./", args[0], v + "." + args[1]))
        if err != nil {
          panic(err)
        }
        genFile.Close()
      }
    },
  }

  var rootCmd = &cobra.Command{Use: "acgen"}
  rootCmd.AddCommand(cmdGenWithName)
  rootCmd.Execute()
}