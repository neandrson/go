package main

import (
	"fmt"
)

/*func createFile(filename string, data []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}*/

/*func createFile(filename string, data []byte) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    return nil
}*/
/*func main() {
	err := createFile("outfile.txt", []byte("Test"))
	fmt.Println("Ошибка записи:", err)
}*/

/*func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}*/

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Hello")
}

// Вывод:
// Hello
// World
