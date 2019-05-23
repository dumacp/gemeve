package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"log"
	"os"
)

var fileMessages string

func init() {
	flag.StringVar(&fileMessages, "fileMessages", "./example.json", "filename of file with messages builder to generate")
}

func main() {
	flag.Parse()

	fmt.Printf("Hola mundo: %s\n", gofakeit.Name())

	file, err := os.Open(fileMessages)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
