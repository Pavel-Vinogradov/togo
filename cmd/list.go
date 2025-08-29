package cmd

import (
	"fmt"
	"togo/internal/task"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать все задачи",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.List()
		if len(tasks) == 0 {
			fmt.Println("Нет задач")
			return
		}
		for _, t := range tasks {
			fmt.Printf("#%d: %s [%s]\n", t.ID, t.Title, t.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
