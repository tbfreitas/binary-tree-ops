package nodes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	error      = "Exiting..."
	size_array = 3
)

func GetData() []int {
	ascii := 0
	var values []int

	for ascii != 27 {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		input, _ := consoleReader.ReadByte()
		ascii := input
		asciiInt := string(ascii)

		i, err := strconv.Atoi(asciiInt)
		if err != nil {
			fmt.Println("Não é um inteiro , tente novamente...")
		} else {
			if len(values) < size_array {
				values = append(values, i)
			} else {
				fmt.Println("Not included! Tree values are fullfield!")
			}
		}
	}

	return values
}

func ThrowTree() {

}
