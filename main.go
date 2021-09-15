package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Coder(inital string, message []byte, counter int) (lineProcessing string) {
	if counter <= len(message) {
		lineProcessing = inital[:len(inital)-2]

		if message[counter] == 1 {
			lineProcessing = lineProcessing + "\n\r "
			return lineProcessing
		} else {
			lineProcessing = lineProcessing + "\r\n "
			return lineProcessing
		}
	} else {
		return inital
	}

}

func DeCoder(inital string) (PartOfMessage byte) {
	if strings.Contains(inital, "\r\n") {
		return 0x00
	} else {
		if strings.HasPrefix(inital, "\r") && len(inital) == 1 {
			return 0x00
		}
		return 0x01
	}
}

func main() {
	fmt.Println("\t \t \t __ENCODING__")
	slice := make([]byte, 5)
	for i := 0; i < cap(slice); i++ {
		if i%2 == 1 {
			slice[i] = 0x00
		} else {
			slice[i] = 0x01
		}

	}
	fmt.Println("Message:", slice)


	f, err := os.Open("Path") //\r\n - 0, \n\r - 1
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
	counter := 0

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
		result := (Coder(line, slice, counter))
		file.WriteString(result)
		counter++

	}
	file.Close()

	fmt.Println("\t \t \t __DECODING__")
	file, err = os.Open("Path")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	reader = bufio.NewReader(file)
	var DecodeMessage []byte
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println(err)
					return
				}

			}
		}
		DecodeMessage = append(DecodeMessage, DeCoder(line))
	}
	fmt.Println(DecodeMessage)
}
