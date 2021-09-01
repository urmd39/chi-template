package infrastructure

import (
	"github.com/gookit/color"
	//"os"
	"log"
	"os"
)

var InfoLog, ErrLog *log.Logger

func init() {
	color.New()
	InfoLog = log.New(os.Stdout, "\u001B[1;34m[INFO]\u001B[0m ", log.Ldate|log.Ltime|log.Llongfile)
	//color.Set(color.FgRed)
	ErrLog = log.New(os.Stderr, "\033[1;31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Llongfile)

	loadEnvParameters()
}
