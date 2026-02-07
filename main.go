package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "addhosts",
		Short: "Утилита для добавления хостов в /etc/hosts",
		Long: `addhosts - это простая утилита для добавления локальных хостов
в файл /etc/hosts с IP-адресом 127.0.0.1

Требуются права администратора для редактирования /etc/hosts`,
	}

	var addCmd = &cobra.Command{
		Use:   "add [хост1 хост2 ...]",
		Short: "Добавить хосты в /etc/hosts",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := checkPermissions(); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				fmt.Println("Попробуйте запустить с sudo: sudo addhosts add", strings.Join(args, " "))
				os.Exit(1)
			}

			hm := NewHostsManager()
			if err := hm.addHosts(args); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				os.Exit(1)
			}
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Показать текущие хосты для 127.0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			hm := NewHostsManager()
			if err := hm.listHosts(); err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkPermissions() error {
	file, err := os.OpenFile("/etc/hosts", os.O_WRONLY, 0644)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("недостаточно прав для записи в /etc/hosts")
		}
		return fmt.Errorf("не удалось проверить доступ к /etc/hosts: %w", err)
	}
	file.Close()
	return nil
}
