package cmd

import (
	"fmt"
	"togo/internal/task"

	"github.com/spf13/cobra"
)

var updID int
var updTitle string
var updDesc string
var updStatus string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Обновить задачу",
	Run: func(cmd *cobra.Command, args []string) {
		if task.Update(updID, updTitle, updDesc, updStatus) {
			fmt.Println("Задача обновлена")
		} else {
			fmt.Println("Задача не найдена")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updID, "id", "i", 0, "ID задачи")
	updateCmd.Flags().StringVarP(&updTitle, "title", "t", "", "Название задачи")
	updateCmd.Flags().StringVarP(&updDesc, "desc", "d", "", "Описание задачи")
	updateCmd.Flags().StringVarP(&updStatus, "status", "s", "", "Статус")
	updateCmd.MarkFlagRequired("id")
}
