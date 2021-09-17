package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func encodingMessage(initial string, message []byte, counter int) (lineProcessing string) {
	lineProcessing = initial[:len(initial)-2]
	if counter > 0 {

		if message[counter] == 1 {
			lineProcessing = lineProcessing + "\n\r "
			return lineProcessing
		} else {
			lineProcessing = lineProcessing + "\r\n "
			return lineProcessing
		}
	} else {
		lineProcessing = lineProcessing + "\r\n "
		return lineProcessing
	}

}

func decodingMessage(initial string) (partOfMessage byte) {
	if strings.Contains(initial, "\r\n") {
		return 0x00
	} else {
		if strings.HasPrefix(initial, "\r") && len(initial) == 1 {
			return 0x00
		}
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
	fmt.Println("Initial message:", message)

	f, err := os.Open("test.txt") //\r\n - 0, \n\r - 1
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
	counter := len(message) - 1
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
		result := (encodingMessage(line, message, counter))
		file.WriteString(result)
		if 0 <= counter {
			counter--
			numOfBits++
		}

	}
	fmt.Println("Possible message to send: ", message[len(message)-numOfBits:])
	defer file.Close()

	fmt.Println("\t \t \t __DECODING__")
	file, err = os.Open("C:\\Users\\User\\GolandProjects\\CRLF\\rez.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	reader = bufio.NewReader(file)

	var decodeMessage []byte
	var reverseMessage []string
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
		reverseMessage = append(reverseMessage, line)
	}

	for i := len(reverseMessage) - 1; i >= 0; i-- {
		decodeMessage = append(decodeMessage, decodingMessage(reverseMessage[i]))

	}
	fmt.Println(decodeMessage)
}
