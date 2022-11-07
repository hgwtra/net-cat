package TCPChat

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

//======= FUNCTIONS =========

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Exit() {
	quit := ""
	fmt.Println("[Type 'x' to quit]")
	fmt.Scan(&quit)
	if quit != "x" {
		Exit()
	} else {
		EmptyHistory()
		os.Exit(0)
	}
}

func ReadFile(filename string) string {
	joinMsg, err := ioutil.ReadFile(filename)
	CheckError(err, "ERROR: Unable to read the file")
	if err != nil {
		os.Exit(1)
	}
	if len(joinMsg) > 0 {
		joinMsg = joinMsg[:len(joinMsg)-1]
	}
	return string(joinMsg)
}

// Write a string to a txt file (append)
func AddToFile(s string) {
	file, err := os.OpenFile("./log/history.txt", os.O_APPEND|os.O_WRONLY, 0644)
	CheckError(err, "ERROR, could not open file")
	defer file.Close()
	_, err = file.WriteString(s + "\n")
	CheckError(err, "ERROR, could not write to file")
}

func EmptyHistory() {
	err := ioutil.WriteFile("./log/history.txt", []byte(""), 0666)
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckError(err error, message string) {
	if err != nil {
		fmt.Printf("%s\n", message)
		panic(err)
	}
}

func ValidPort(port string) bool {
	i, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if i < 1024 || i > 65535 {
		return false
	}
	return true
}

func CheckPort() (string, bool) {
	if len(os.Args) == 1 {
		return "8989", true
	}
	if len(os.Args) == 2 {
		if ValidPort(os.Args[1]) {
			return os.Args[1], true
		} else {
			return "", false
		}
	}
	return "", false
}
