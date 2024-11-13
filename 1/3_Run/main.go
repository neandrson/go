package main

import (
	"fmt"
	"os"
	"strconv"
)

func run() error {
	var par1, par2, par3 int
	if len(os.Args) < 3 {
		return fmt.Errorf("необходимо указать размер сетки и процент заполнения")
	}

	par1, err1 := strconv.Atoi(os.Args[1])
	par2, err2 := strconv.Atoi(os.Args[2])
	par3, err3 := strconv.Atoi(os.Args[3])

	if par1 == 0 {
		return err1
	} else if par2 == 0 {
		return err2
	} else if par3 == 0 {
		return err3
	}
	fmt.Println(par1, par2, par3)

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
