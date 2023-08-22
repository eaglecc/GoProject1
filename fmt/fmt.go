package fmt

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, error := os.OpenFile("trace.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if error != nil {
		panic(error)
	}
	Logger = log.New(io.MultiWriter(os.Stdout, file), "cyc Log:", log.LstdFlags)
}
