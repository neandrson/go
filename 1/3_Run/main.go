package main

import (
	"fmt"
	"os"
)

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("необходимо указать размер сетки и процент заполнения")
	}
	a := os.Args[1]
	b := os.Args[2]
	c := os.Args[3]
	if a < "0" || b < "0" || c <= "0" {
		return fmt.Errorf("Должно быть больше нуля")
	}

	gridSize := os.Args[1] + "x" + os.Args[2]
	fillPercentage := os.Args[3] + "%"

	file, err := os.Create("config.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(gridSize + " " + fillPercentage)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
