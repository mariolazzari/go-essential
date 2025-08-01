package main

import (
	"fmt"
	"os"
)

func main() {
	err := killServer("server.pid")
	if err != nil {
		fmt.Println("Error killing server:", err)
	} else {
		fmt.Println("Server killed successfully")
	}
	fmt.Println("Exiting main function")
}

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return fmt.Errorf("failed to read PID file: %w", err)
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("failed to parse PID from file: %w", err)
	}

	fmt.Println("Killing server with PID:", pid)

	if err := os.Remove(pidFile); err != nil {
		return fmt.Errorf("failed to remove PID file: %w", err)
	}

	return nil
}
