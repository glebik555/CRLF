package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("\t \t \t __ENCODING__")
	dbg := flag.Bool("dbg", false, "debug")

	flag.Parse()

	if *dbg {
		logrus.Info("Running in debug mode")
		logrus.SetLevel(logrus.DebugLevel)
	}

	message := makeMessage(100)

	f := openFile("texts\\test.txt")
	defer f.Close()

	file := createFile("texts\\rez.txt")

	logrus.Debug("Injecting message into container")
	injectProcessing(f, message, file)
	logrus.Debug("Message injection completed successfully")
	file.Close()

	fmt.Println("\t \t \t __DECODING__")
	logrus.Debug("Start extracting")

	file = openFile("texts\\rez.txt")
	logrus.Debug("Extracting message from container")
	extractProcessing(file)

	logrus.Debug("Message extracting completed successfully")

	defer file.Close()
}
