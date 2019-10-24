package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(array []int, size_array int) {
	var numberOfSwaps = 0
	for i := 0; i < size_array; i++ {

		for j := 0; j < size_array-1; j++ {
			if array[j] > array[j+1] {
				array = Swap(j, array)
				numberOfSwaps++
			}
		}

		if numberOfSwaps == 0 {
			break
		}

	}

	s := fmt.Sprintf("Array is sorted in %d swaps.\nFirst Element: %d\nLast Element: %d", numberOfSwaps, array[0], array[len(array)-1])
	io.WriteString(os.Stdout, s)

}

func Swap(val int, a []int) []int {
	for i, el := range a {
		if i == val {
			aux := a[i+1]
			a[i+1] = el
			a[i] = aux
			break
		}
	}
	return a
}

func main() {

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	var buffer []int
	size_array, _ := strconv.Atoi(in.Text())
	in.Scan()
	input := in.Text()
	bufferStrings := strings.Split(input, " ")
	for _, element := range bufferStrings {
		data, _ := strconv.Atoi(element)
		buffer = append(buffer, data)
	}

	BubbleSort(buffer, size_array)

}
