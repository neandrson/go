package main

import (
	"flag"
	"fmt"
	"os"
)

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("необходимо указать размер сетки и процент заполнения")
	}

	if flag.Int(os.Args[1]) < "0" || flag.Int(os.Args[2]) < "0" || flag.Int(os.Args[3]) < "0" {
		return fmt.Errorf("Не должно быть нулем")
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