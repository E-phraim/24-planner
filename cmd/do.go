/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "done",
	Short: "Yay! Plan Executed.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args{
			id, err := strconv.Atoi(arg)
			if err != nil{
				fmt.Println("Failed to parse the argument:", arg)
			}else{
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
