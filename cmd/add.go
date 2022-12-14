package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a plan to your activity",
	Run: func (cmd *cobra.Command, args []string)  {
		plan := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your plans for today.\n", plan)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}