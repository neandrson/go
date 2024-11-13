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
	fmt.Println(par1, par2, par3, err3)

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if err3 != nil {
		return err3
	}

	//fmt.Println(par1, par2, par3)

	gridSize := strconv.Itoa(par1) + "x" + strconv.Itoa(par2)
	fillPercentage := strconv.Itoa(par3) + "%"

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
