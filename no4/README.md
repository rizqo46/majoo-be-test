## Nomor 4
Pseudocode untuk megurutkan angka:
```go
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
```
