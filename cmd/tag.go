package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage blog tags",
}

// CRUD commands for tags
var createTagCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new tag",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// Add code to create tag in the database
		fmt.Printf("Tag '%s' created.\n", name)
	},
}

var readTagCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Read a tag by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// Add code to read tag from the database
		fmt.Printf("Details of tag with ID '%s'.\n", id)
	},
}

var updateTagCmd = &cobra.Command{
	Use:   "update [id] [new name]",
	Short: "Update a tag by its ID",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		newName := args[1]
		// Add code to update tag in the database
		fmt.Printf("Tag with ID '%s' updated to '%s'.\n", id, newName)
	},
}

var deleteTagCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a tag by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// Add code to delete tag from the database
		fmt.Printf("Tag with ID '%s' deleted.\n", id)
	},
}

// init initializes tag commands
func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.AddCommand(createTagCmd, readTagCmd, updateTagCmd, deleteTagCmd)
}
