package utils

import "os"

func ToNode(str string) {
	os.Stdout.WriteString(str)
}
