package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Manage blog posts",
}

// CRUD commands for posts
var createPostCmd = &cobra.Command{
	Use:   "create [title]",
	Short: "Create a new post",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		// Add code to create post in the database
		fmt.Printf("Post '%s' created.\n", title)
	},
}

var readPostCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Read a post by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// Add code to read post from the database
		fmt.Printf("Details of post with ID '%s'.\n", id)
	},
}

var updatePostCmd = &cobra.Command{
	Use:   "update [id] [new title]",
	Short: "Update a post by its ID",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		newTitle := args[1]
		// Add code to update post in the database
		fmt.Printf("Post with ID '%s' updated to '%s'.\n", id, newTitle)
	},
}

var deletePostCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a post by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// Add code to delete post from the database
		fmt.Printf("Post with ID '%s' deleted.\n", id)
	},
}

// init initializes post commands
func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.AddCommand(createPostCmd, readPostCmd, updatePostCmd, deletePostCmd)
}
