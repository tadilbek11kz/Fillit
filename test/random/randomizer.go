package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func TetroRandomizer(num int) []string {
	var arr []string
	file, err := os.Open("tetro.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	var tetro []string
	for _, v := range strings.Split(string(b), "\n\n") {
		if v != "" {
			tetro = append(tetro, v)
		}
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		arr = append(arr, tetro[rand.Intn(len(tetro))])
	}
	return arr

}

func Write(fileName string, num int) {
	toWrite := TetroRandomizer(num)
	file, err := os.Create(fileName)
	if err != nil {
		panic("Cannot write to file")
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	formatedString := strings.Join(toWrite, "\n\n")
	w.WriteString(formatedString + "\n\n")
	w.Flush()
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	Write("test.txt", n)
}
