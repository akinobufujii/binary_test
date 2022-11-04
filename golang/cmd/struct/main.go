package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type TestData struct {
	Hoge int32
	Fuga int8
	Fizz float32
	Buzz int16
}

func readTestData(filename string) (*TestData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed os.Open: %s %w", filename, err)
	}
	defer file.Close()

	testData := new(TestData)
	if err := binary.Read(file, binary.LittleEndian, testData); err != nil {
		return nil, fmt.Errorf("failed binary.Read: %s %w", filename, err)
	}

	return testData, nil
}

func writeTestData(filename string, testData *TestData) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed os.Create: %s %w", filename, err)
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, testData); err != nil {
		return fmt.Errorf("failed binary.Write: %s %w", filename, err)
	}

	return nil
}

func main() {
	writeData := TestData{
		Hoge: 10,
		Fuga: 20,
		Fizz: 30,
		Buzz: 40,
	}

	const dataName = "testdata.data"

	if err := writeTestData(dataName, &writeData); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", errors.WithStack(err))
		os.Exit(1)
	}

	readData, err := readTestData(dataName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", errors.WithStack(err))
		os.Exit(1)
	}

	fmt.Printf("writeData: %+v\n", writeData)
	fmt.Printf("readData: %+v\n", readData)
}
