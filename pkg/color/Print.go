package color

import "fmt"

// I love ANSI Escape Codes

func PrintRed(printee string) {
	fmt.Print("\033[31m" + printee + "\033[0m")
}

func PrintBlue(printee string) {
	fmt.Print("\033[34m" + printee + "\033[0m")
}

func PrintGreen(printee string) {
	fmt.Print("\033[32m" + printee + "\033[0m")
}

func PrintYellow(printee string) {
	fmt.Print("\033[33m" + printee + "\033[0m")
}

func PrintMagenta(printee string) {
	fmt.Print("\033[35m" + printee + "\033[0m")
}

func PrintRainbow(printee string) {
	magicNumber := 0
	for _, char := range(printee) {
		switch magicNumber {
		default:
			magicNumber = 0
			PrintBlue(string(char))

		case 1:
			PrintGreen(string(char))

		case 2:
			PrintMagenta(string(char))

		case 3:
			PrintYellow(string(char))

		case 4:
			PrintRed(string(char))
		}

		magicNumber++
	}
}
