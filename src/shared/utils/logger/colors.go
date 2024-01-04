package logger

import "fmt"

func magenta(text string) string {
	return fmt.Sprintf("\x1B[95m%s\x1B[39m", text)
}

func bold(text string) string {
	return fmt.Sprintf("\x1B[1m%s\x1B[0m", text)
}

func red(text string) string {
	return fmt.Sprintf("\x1B[31m%s\x1B[39m", text)
}

func green(text string) string {
	return fmt.Sprintf("\x1B[32m%s\x1B[39m", text)
}

func yellow(text string) string {
	return fmt.Sprintf("\x1B[33m%s\x1B[39m", text)
}

func cyan(text string) string {
	return fmt.Sprintf("\x1B[96m%s\x1B[39m", text)
}
