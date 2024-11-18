package main

import "io"

func ReadString(r io.Reader) (string, error) {
	var rs string
	buf, err := io.ReadAll(r)
	if err == nil {
		rs = string(buf)
	}
	return rs, err
}

/*func ReadString(r io.Reader) (string, error) {
    bytes, err := ioutil.ReadAll(r) // читаем все данные из r в срез байтов
    if err != nil {
        return "", err // возвращаем пустую строку и ошибку
    }
    return string(bytes), nil // преобразуем срез байтов в строку и возвращаем её и nil
}*/

func main() {

}
