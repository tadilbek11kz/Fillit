package main

import (
	"fmt"
	"testing"
	"time"

	optimzer ".."
)

func TestSample(t *testing.T) {
	filename := "sample.txt"
	start := time.Now()
	data := optimzer.New(filename)
	start = time.Now()
	result := data.Solve()

	total := time.Since(start).Minutes()
	fmt.Println("File: ", filename)
	optimzer.Print(result)
	fmt.Println("Solve: ", total)
	if 5 < total {
		t.Error(
			"Error: Test2")
		t.Error("Took:", total)
		t.Error("Exp:", 5)
	}
}

func TestGoodexample00(t *testing.T) {
	filename := "goodexample00.txt"
	start := time.Now()
	data := optimzer.New(filename)
	start = time.Now()
	result := data.Solve()

	total := time.Since(start).Minutes()
	fmt.Println("File: ", filename)
	optimzer.Print(result)
	fmt.Println("Solve: ", total)
	if 5 < total {
		t.Error(
			"Error: Test2")
		t.Error("Took:", total)
		t.Error("Exp:", 5)
	}
}

func TestGoodexample01(t *testing.T) {
	filename := "goodexample01.txt"
	start := time.Now()
	data := optimzer.New(filename)
	start = time.Now()
	result := data.Solve()

	total := time.Since(start).Minutes()
	fmt.Println("File: ", filename)
	optimzer.Print(result)
	fmt.Println("Solve: ", total)
	if 5 < total {
		t.Error(
			"Error: Test2")
		t.Error("Took:", total)
		t.Error("Exp:", 5)
	}
}

func TestGoodexample02(t *testing.T) {
	filename := "goodexample02.txt"
	start := time.Now()
	data := optimzer.New(filename)
	start = time.Now()
	result := data.Solve()

	total := time.Since(start).Minutes()
	fmt.Println("File: ", filename)
	optimzer.Print(result)
	fmt.Println("Solve: ", total)
	if 5 < total {
		t.Error(
			"Error: Test2")
		t.Error("Took:", total)
		t.Error("Exp:", 5)
	}
}

func TestGoodexample03(t *testing.T) {
	filename := "goodexample03.txt"
	start := time.Now()
	data := optimzer.New(filename)
	start = time.Now()
	result := data.Solve()

	total := time.Since(start).Minutes()
	fmt.Println("File: ", filename)
	optimzer.Print(result)
	fmt.Println("Solve: ", total)
	if 5 < total {
		t.Error(
			"Error: Test2")
		t.Error("Took:", total)
		t.Error("Exp:", 5)
	}
}

func TestBadexample(t *testing.T) {
	catch := func() {
		if a := recover(); a != "ERROR" {
			t.Error("Error: TestBadexample")
			t.Error("Act:", a)
			t.Error("Exp:", "ERROR")
		}
	}
	filenames := []string{
		"badexample00.txt",
		"badexample01.txt",
		"badexample02.txt",
		"badexample03.txt",
		"badformat.txt"}
	for _, val := range filenames {
		defer catch()
		data := optimzer.New(val)
		data.Solve()
	}

}
