package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// ✔️
func main() {
	fmt.Println("\nWillkommen! Es ist aktuell", time.Now().Format("15:04"), "Uhr\n\nBitte geben Sie an was Sie tun wollen:\n1. To-Do-Liste bearbeiten\n2. Einträge abhaken\n3. Liste neu Erstellen")

	//Erstellen eines Scanners und lesen des Terminal inputs
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "1" {
		file, err := os.Create("ToDo.txt")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		file.WriteString("Hey mein Name ist Leon\nHey")
	}
	readFile()
}

func readFile() {
	file, err := os.Open("ToDo.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
