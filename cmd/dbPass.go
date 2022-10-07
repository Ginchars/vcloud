/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/Ginchars/vcloud/utils"
)

// dbPassCmd represents the dbPass command
var dbPassCmd = &cobra.Command{
	Use:   "dbPass [project-name]",
	Short: "Fetch DB password",
	Long: `Fetch password for DB, for specified environment`,
	Run: func(cmd *cobra.Command, args []string) {
		var projectName = ""

		if len(args) >= 1 && args[0] != "" {
			projectName = args[0]
			fmt.Print(utils.GetDBPass(projectName))
		} else {
			fmt.Println("Please specify project")
		}
	},
}

func init() {
	rootCmd.AddCommand(dbPassCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbPassCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbPassCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
