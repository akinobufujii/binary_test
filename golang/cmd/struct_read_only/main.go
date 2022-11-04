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

func readTestData(filename string, testData *TestData) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed os.Open: %s %w", filename, err)
	}
	defer file.Close()

	if err := binary.Read(file, binary.LittleEndian, testData); err != nil {
		return fmt.Errorf("failed binary.Read: %s %w", filename, err)
	}

	return nil
}

func main() {
	const dataName = "testdata.data"

	readData := TestData{}
	if err := readTestData(dataName, &readData); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", errors.WithStack(err))
		os.Exit(1)
	}

	fmt.Printf("readData: %+v\n", readData)
}
