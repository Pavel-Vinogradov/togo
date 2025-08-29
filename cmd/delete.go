package cmd

import (
	"fmt"
	"togo/internal/task/service"

	"github.com/spf13/cobra"
)

var id int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Удалить задачу",
	Run: func(cmd *cobra.Command, args []string) {
		if service.Delete(id) {
			fmt.Printf("Задача удалена: #%d\n", id)
		} else {
			fmt.Printf("Задача с ID %d не найдена\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "ID задачи")
	_ = deleteCmd.MarkFlagRequired("id")
}
