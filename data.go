package optimizer

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

type Data struct {
	filename    string
	tetrominoes *Sequence
}

// New creates new struct
func New(filename string) (d Data) {
	d.setFilename(filename)
	d.Read(d.filename)
	return d
}

//setFilename set filename
func (d *Data) setFilename(filename string) {
	if filename == "" {
		panic("Empty filename")
	}
	d.filename = filename
}

//seTetrominoes read tetrominoes from file
func (d *Data) setTetrominoes(scanner *bufio.Scanner) {
	d.tetrominoes = &Sequence{}
	obj := make([][]byte, 4)
	for i := 0; scanner.Scan(); {
		rowVal := scanner.Text()
		if len(rowVal) == 0 {
			if len(obj[0]) == 0 {
				continue
			}
			d.tetrominoes.PushBack(obj)
			obj = make([][]byte, 4)
			i = 0
			continue
		}
		if i == 4 {
			panic("ERROR")
		}
		obj[i] = []byte(rowVal)
		i++
	}
}

//Read start reading from file
func (d *Data) Read(path string) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Cannot read such file")
	}
	reader := bytes.NewReader(buff)
	scanner := bufio.NewScanner(reader)
	d.setTetrominoes(scanner)
}

//Print print solution
func Print(solution string) {
	length := int(math.Sqrt(float64(len(solution))))
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if solution[i*length+j] == '.' {
				fmt.Print(string('.'))
			} else {
				fmt.Print(string('A' + solution[i*length+j] - 1))
			}
		}
		fmt.Print(string('\n'))
	}
}
