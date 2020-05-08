/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"inabajunmr/dicc/api"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search word",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			// TODO
		}

		// TODO using cache

		dictionaryApi := api.GetApi(api.WEBSTER)
		condition := api.SearchCondition{args[0]}
		result, _ := dictionaryApi.SearchWords(condition)
		for _, def := range result.Definitions {
			fmt.Println("------")
			fmt.Printf("fl:%v\n", def.FunctionalLabel)
			for i, desc := range def.Descriptions {
				fmt.Printf("%v. %v\n", i, desc)
			}
		}

		// TODO save as history
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
