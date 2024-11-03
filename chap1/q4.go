package chap1

import "fmt"

func q4() {
	numbers := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	fmt.Println(CalAvg(numbers))
}
func CalAvg(slice []float64) (float64, error) {

	if len(slice) == 0 {
		return 0, fmt.Errorf("slice can't be empty")
	}
	sum := 0.0

	for _, num := range slice {
		sum = sum + num
	}

	average := sum / float64(len(slice))
	return average, nil
}
