package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("numbers.txt")
	st := strings.Split(string(file), "\n")
	var nums []float64
	for _, v := range st {
		n, _ := strconv.Atoi(v)
		nums = append(nums, float64(n))
	}

	length := len(nums)
	nums = nums[:length-1]

	mean := Mean(nums)
	median := Median(nums)
	div := Variance(nums, mean)
	tdeviance := SD(div)

	fmt.Println("Average: " + fmt.Sprint(math.Round(mean)))

	fmt.Println("Median: " + fmt.Sprint(math.Round(median)))

	fmt.Println("variance: " + fmt.Sprint(int(math.Round(div))))

	fmt.Println("Standard Deviation: " + fmt.Sprint(math.Round(tdeviance)))
}

func Mean(nums []float64) float64 {
	var sum float64

	for j := 0; j < len(nums); j++ {
		sum += nums[j]
	}

	mean := sum / float64(len(nums))

	return mean
}

func Median(nums []float64) float64 {
	sort.Float64s(nums)
	if len(nums)%2 != 0 {
		return nums[(len(nums)+1)/2]
	} else {
		return (nums[len(nums)/2] + nums[(len(nums)/2)-1]) / 2
	}
	// check to see if length of nums is odd or even using %
	// if odd .... else if even ....
}

func Variance(nums []float64, mean float64) float64 {
	var sdsum float64
	var div float64

	for sd := 0; sd < len(nums); sd++ {
		sdsum += (nums[sd] - mean) * (nums[sd] - mean)
		div = sdsum / float64(len(nums))
	}
	return div
}

func SD(div float64) float64 {
	tdeviance := math.Sqrt(div)
	return tdeviance
}
