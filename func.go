package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"strings"
)

func makeMessage(length int) []byte {
	logrus.Debug("Making message")
	message := make([]byte, length) // Example of message
	for i := 0; i < cap(message); i++ {
		if i%2 == 1 {
			message[i] = 0x00
		} else {
			message[i] = 0x01
		}

	}
	log.Println("Message to send: ", message)
	return message
}

func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err.Error())
	}
	return f
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		panic(1)
	}
	return file
}

func injectProcessing(f *os.File, message []byte, file *os.File) {
	reader := bufio.NewReader(f)
	indexOfMessage := len(message) - 1
	numOfBits := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}

		}

		fmt.Println(line)
		result := injectMessage(line, message, indexOfMessage)
		_, err = file.WriteString(result)
		if err != nil {
			fmt.Println(err.Error())
		}
		if 0 <= indexOfMessage {
			indexOfMessage--
			numOfBits++
		}

	}
	fmt.Println("Possible message to send: ", message[len(message)-numOfBits:], ".\nIt's", numOfBits, "bit.")
	if numOfBits%8 == 0 {
		fmt.Println("It's", numOfBits/8, "byte")
	}
}

func extractProcessing(file *os.File) {
	reader := bufio.NewReader(file)

	var decodeMessage []byte
	var containerText []string

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}

		}

		containerText = append(containerText, line)
	}

	for i := len(containerText) - 1; i >= 0; i-- {
		decodeMessage = append(decodeMessage, extractMessage(containerText[i]))
	}

	fmt.Println(decodeMessage)
}

func injectMessage(containerString string, message []byte, bitCounter int) (enterMessage string) { // bitCounter - controls the size of the transmitted message and container

	enterMessage = containerString[:len(containerString)-2]
	if bitCounter >= 0 {

		if message[bitCounter] == 1 {
			enterMessage = enterMessage + "\n\r "
			return enterMessage
		} else {
			enterMessage = enterMessage + "\r\n "
			return enterMessage
		}
	} else {
		enterMessage = enterMessage + "\r\n "
		return enterMessage
	}

}

func extractMessage(containerString string) (partOfMessage byte) {
	if strings.Contains(containerString, "\r\n") {
		return 0x00
	} else {
		return 0x01
	}
}
