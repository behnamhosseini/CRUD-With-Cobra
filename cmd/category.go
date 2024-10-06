package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage blog categories",
}

// CRUD commands for categories
var createCategoryCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new category",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// Add code to create category in the database
		fmt.Printf("Category '%s' created.\n", name)
	},
}

var readCategoryCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Read a category by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// Add code to read category from the database
		fmt.Printf("Details of category with ID '%s'.\n", id)
	},
}

var updateCategoryCmd = &cobra.Command{
	Use:   "update [id] [new name]",
	Short: "Update a category by its ID",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		newName := args[1]
		// Add code to update category in the database
		fmt.Printf("Category with ID '%s' updated to '%s'.\n", id, newName)
	},
}

var deleteCategoryCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a category by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// Add code to delete category from the database
		fmt.Printf("Category with ID '%s' deleted.\n", id)
	},
}

// init initializes category commands
func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.AddCommand(createCategoryCmd, readCategoryCmd, updateCategoryCmd, deleteCategoryCmd)
}
