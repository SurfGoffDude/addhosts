package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const hostsFile = "/etc/hosts"

type HostsManager struct {
	filePath string
}

func NewHostsManager() *HostsManager {
	return &HostsManager{
		filePath: hostsFile,
	}
}

func (hm *HostsManager) readHosts() ([]string, error) {
	file, err := os.Open(hm.filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", hm.filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %w", hm.filePath, err)
	}

	return lines, nil
}

func (hm *HostsManager) writeHosts(lines []string) error {
	file, err := os.Create(hm.filePath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", hm.filePath, err)
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("ошибка записи в файл %s: %w", hm.filePath, err)
		}
	}

	return nil
}

func (hm *HostsManager) hostExists(host string) (bool, error) {
	lines, err := hm.readHosts()
	if err != nil {
		return false, err
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			for i := 1; i < len(fields); i++ {
				if fields[i] == host {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func (hm *HostsManager) addHosts(hosts []string) error {
	lines, err := hm.readHosts()
	if err != nil {
		return err
	}

	var newEntries []string
	for _, host := range hosts {
		exists, err := hm.hostExists(host)
		if err != nil {
			return err
		}
		
		if !exists {
			newEntries = append(newEntries, fmt.Sprintf("127.0.0.1\t%s", host))
		}
	}

	if len(newEntries) == 0 {
		fmt.Println("Все указанные хосты уже существуют")
		return nil
	}

	lines = append(lines, newEntries...)
	
	if err := hm.writeHosts(lines); err != nil {
		return err
	}

	fmt.Printf("Добавлено хостов: %d\n", len(newEntries))
	for _, entry := range newEntries {
		fmt.Printf("  %s\n", entry)
	}

	return nil
}

func (hm *HostsManager) listHosts() error {
	lines, err := hm.readHosts()
	if err != nil {
		return err
	}

	type hostEntry struct {
		ip    string
		hosts []string
	}

	var entries []hostEntry
	maxIPLen := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) < 2 {
			continue
		}

		ip := fields[0]
		if len(ip) > maxIPLen {
			maxIPLen = len(ip)
		}
		entries = append(entries, hostEntry{ip: ip, hosts: fields[1:]})
	}

	if len(entries) == 0 {
		fmt.Println("  Файл /etc/hosts пуст")
		return nil
	}

	headerIP := "IP"
	headerHost := "HOSTNAME"
	if maxIPLen < len(headerIP) {
		maxIPLen = len(headerIP)
	}

	topBorder := "┌" + strings.Repeat("─", maxIPLen+2) + "┬" + strings.Repeat("─", 40) + "┐"
	separator := "├" + strings.Repeat("─", maxIPLen+2) + "┼" + strings.Repeat("─", 40) + "┤"
	bottomBorder := "└" + strings.Repeat("─", maxIPLen+2) + "┴" + strings.Repeat("─", 40) + "┘"

	fmt.Println()
	fmt.Printf("  %s\n", topBorder)
	fmt.Printf("  │ %-*s │ %-38s│\n", maxIPLen, headerIP, headerHost)
	fmt.Printf("  %s\n", separator)

	for i, e := range entries {
		fmt.Printf("  │ %-*s │ %-38s│\n", maxIPLen, e.ip, strings.Join(e.hosts, ", "))
		if i < len(entries)-1 {
			fmt.Printf("  %s\n", separator)
		}
	}

	fmt.Printf("  %s\n", bottomBorder)
	fmt.Println()
	fmt.Printf("  Всего записей: %d\n\n", len(entries))

	return nil
}
