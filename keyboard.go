// Пакет keyboard читает пользовательский ввод с клавиатуры
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat читает числа с плавающей точкой с клавиатуры.
// Она возвращает считанное число и ошибку, если она произошла.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
