package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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

func main() {
	fmt.Println("\t \t \t __ENCODING__")
	message := make([]byte, 5) // Example of message
	for i := 0; i < cap(message); i++ {
		if i%2 == 1 {
			message[i] = 0x00
		} else {
			message[i] = 0x01
		}

	}
	fmt.Println("Message to send: ", message)

	f, err := os.Open("test.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()

	file, err := os.Create("rez.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

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
	fmt.Println("Possible message to send: ", message[len(message)-numOfBits:])
	defer file.Close()

	fmt.Println("\t \t \t __DECODING__")
	file, err = os.Open("rez.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	reader = bufio.NewReader(file)

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
