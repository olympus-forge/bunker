package bunker

import (
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func IsEmptyString(str string) bool {
	return strings.TrimSpace(str) == ""
}

func PrintErr(msg string) {
	fmt.Println(color.Ize(color.Red, fmt.Sprintf("[DEBUG-ERROR] %s", msg)))
}

func FatalPrintErr(msg string) {
	fmt.Println(color.Ize(color.Red, fmt.Sprintf("[DEBUG-FATAL] %s", msg)))
	os.Exit(1)
}

func LogInfo(msg string) {
	fmt.Println(color.Ize(color.Gray, fmt.Sprintf("[DEBUG-INFO] %s", msg)))
}

func LogWarn(msg string) {
	fmt.Println(color.Ize(color.Yellow, fmt.Sprintf("[DEBUG-WARN] %s", msg)))
}
