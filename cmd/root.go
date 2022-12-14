package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "planner",
	Short: "Planner is a CLI that helps you specify and organise tasks.",
}