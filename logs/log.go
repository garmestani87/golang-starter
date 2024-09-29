package main

import (
	"fmt"
	"log"
	"os"
)

var(
	infoLogger *log.Logger
	errorLogger *log.Logger
)

func init() {
	fmt.Println("init !!!")
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.SetOutput(file)
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}else{
		log.Fatalln("unexpected error !")
	}
	flags := log.Ldate | log.Ltime | log.Llongfile
	infoLogger = log.New(file, "info - ", flags)
	errorLogger = log.New(file, "error - ", flags)
}

func main() {
	fmt.Println("main !!!")
	logPrint(5)
}

func logPrint(count int) {
	for i := 0; i < count; i++ {
		errorLogger.Println(i)
	}
}
