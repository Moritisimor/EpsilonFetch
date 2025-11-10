package color

import "fmt"

func PrintRedln(printee string) {
	fmt.Println("\033[31m" + printee + "\033[0m")
}

func PrintBlueln(printee string) {
	fmt.Println("\033[34m" + printee + "\033[0m")
}

func PrintCyanln(printee string) {
	fmt.Println("\033[36m" + printee + "\033[0m")
}

func PrintBlackln(printee string) {
	fmt.Println("\033[30m" + printee + "\033[0m")
}

func PrintGreenln(printee string) {
	fmt.Println("\033[32m" + printee + "\033[0m")
}

func PrintYellowln(printee string) {
	fmt.Println("\033[33m" + printee + "\033[0m")
}

func PrintMagentaln(printee string) {
	fmt.Println("\033[35m" + printee + "\033[0m")
}

func PrintRainbowln(printee string) {
	PrintRainbow(fmt.Sprintf("%s\n", printee))
}
