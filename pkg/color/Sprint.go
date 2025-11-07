package color

func SprintRed(printee string) string {
	return "\033[31m" + printee + "\033[0m"
}

func SprintBlue(printee string) string {
	return "\033[34m" + printee + "\033[0m"
}

func SprintGreen(printee string) string {
	return "\033[32m" + printee + "\033[0m"
}

func SprintYellow(printee string) string {
	return "\033[33m" + printee + "\033[0m"
}

func SprintMagenta(printee string) string {
	return "\033[35m" + printee + "\033[0m"
}