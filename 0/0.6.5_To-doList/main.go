package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func PrettyArrayOutput(array [9]string) {
	var j int

	for i := 0; i <= 8; i++ {
		//fmt.Println(array[i])
		if i <= 6 {
			j++
			fmt.Println(j, "я уже сделал:", array[i])
		} else {
			j++
			fmt.Println(j, "не успел сделать:", array[i])
		}
	}
}

func main() {
	input := [9]string{
		"проснуться",
		"позавтракать",
		"сходить в школу",
		"пообедать",
		"погулять с друзьями",
		"сделать домашнюю работу",
		"попрограммировать на Go",
		"поужинать",
		"лечь спать",
	}
	PrettyArrayOutput(input)
}
