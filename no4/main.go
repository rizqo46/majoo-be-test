package main

import "fmt"

func sort(input []float64, arg string) []float64 {
	n := len(input)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			switch arg {
			case "desc", "descending":
				if input[i] < input[j] {
					input[i], input[j] = input[j], input[i]
				}
			default:
				if input[i] > input[j] {
					input[i], input[j] = input[j], input[i]
				}
			}
		}
	}
	return input
}

func main() {
	t := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	res := sort(t, "desc")
	fmt.Println(res)

}
