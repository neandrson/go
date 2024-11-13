package main

import (
	"fmt"
	"os"
)

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("необходимо указать размер сетки и процент заполнения")
	}

	par1 := os.Args[1]
	par2 := os.Args[2]
	par3 := os.Args[3]

	if par1 < string("0") || par2 < string("0") || par3 < string("0") {
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
