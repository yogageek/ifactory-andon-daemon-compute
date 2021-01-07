package logger

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-isatty"

	"github.com/sirupsen/logrus"
)

const (
	colorNoColor = "\033[0m"
	colorRed     = "\033[91m"
	colorGreen   = "\033[92m"
	colorYellow  = "\033[93m"
	colorBlue    = "\033[94m"
	colorMagenta = "\033[95m"
	colorCyan    = "\033[96m"
	Default      = "\033[39m"
	Black        = "\033[30m"
	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Blue         = "\033[34m"
	Magenta      = "\033[35m"
	Cyan         = "\033[36m"
	LightGray    = "\033[37m"
	DarkGray     = "\033[90m"
	LightRed     = "\033[91m"
	LightGreen   = "\033[92m"
	LightYellow  = "\033[93m"
	LightBlue    = "\033[94m"
	LightMagenta = "\033[95m"
	LightCyan    = "\033[96m"
	White        = "\033[97m"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var (
	isTerminal bool
)

func init() {
	isTerminal = isatty.IsTerminal(os.Stdout.Fd())
}

type textFormat struct {
	forceColors bool
}

func NewTextFormat(forceColor ...bool) *textFormat {
	return &textFormat{
		forceColors: len(forceColor) == 1 && forceColor[0],
	}
}

func (f *textFormat) Format(entry *logrus.Entry) ([]byte, error) {
	levelText := strings.ToUpper(entry.Level.String())[0:4]
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	//if (f.forceColors || isTerminal) && runtime.GOOS != "windows" {
	color := colorNoColor
	switch entry.Level {
	case logrus.DebugLevel:
		color = colorBlue
	case logrus.InfoLevel:
		color = White
	case logrus.WarnLevel:
		color = colorYellow
	case logrus.ErrorLevel:
		color = colorMagenta
	case logrus.PanicLevel, logrus.FatalLevel:
		color = colorRed
	}
	buf.WriteString(color)
	//}
	buf.WriteString(fmt.Sprintf("[%s] ", entry.Time.Format(timeFormat)))
	buf.WriteString(fmt.Sprintf("[%s] ", levelText))
	for k, v := range entry.Data {
		buf.WriteString(fmt.Sprintf("[%s=%v] ", k, v))
	}
	buf.WriteString(entry.Message)
	if (f.forceColors || isTerminal) && runtime.GOOS != "windows" {
		buf.WriteString(colorNoColor)
	}
	buf.WriteString("\n")
	buf.WriteString(colorNoColor)
	return buf.Bytes(), nil
}
