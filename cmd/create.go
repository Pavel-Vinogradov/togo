package cmd

import (
	"fmt"
	"togo/internal/task"

	"github.com/spf13/cobra"
)

var addTitle string
var addDesc string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить новую задачу",
	Run: func(cmd *cobra.Command, args []string) {
		t := task.Create(addTitle, addDesc)
		fmt.Printf("Задача создана: #%d %s\n", t.ID, t.Title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addTitle, "title", "t", "", "Название задачи")
	addCmd.Flags().StringVarP(&addDesc, "desc", "d", "", "Описание задачи")
	addCmd.MarkFlagRequired("title")
}
