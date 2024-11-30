package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func main() {
	m := map[string]string{
		"Яндекс":        "+74957397000",
		"Музей Яндекса": "+74991101133",
	}
	fmt.Println(SwapKeysAndValues(m))
}
