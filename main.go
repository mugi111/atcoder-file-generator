package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "github.com/spf13/cobra"
)

func main() {
  var cmdGenWithName = &cobra.Command{
    Use:   "gen [Contest Name]",
    Short: "Generate Files",
    Long: `Generate folders and files with the name you entered.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      var path, _ = os.Getwd()
      files, err := ioutil.ReadDir(path)
      if err != nil {
          panic(err)
      }
      var paths []string
      for _, file := range files {
        paths = append(paths, file.Name())
      }
      fmt.Println(paths)
    },
  }

  var rootCmd = &cobra.Command{Use: "acgen"}
  rootCmd.AddCommand(cmdGenWithName)
  rootCmd.Execute()
}