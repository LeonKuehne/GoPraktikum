package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {

	//Erstellt eine Neue To-Do-Liste falls diese noch nicht existiert
	if !checkIfExists() {
		createFile()
	}

	//Unendliche Schleife um das Programm nicht nach einer anweisung zu beenden
	for {
		fmt.Println("\nWillkommen! Es ist aktuell", time.Now().Format("15:04"), "Uhr\n\nBitte geben Sie an was Sie tun wollen:\n1. Eintrag hinzufügen\n2. Liste anzeigen\n3. Einträge abhaken\n4. Liste neu Erstellen\n5. Beenden")

		//Erstellen eines Scanners und lesen des Terminal inputs
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		//Kontrollstruktur um die verschiedenen Befehle zu differenzieren
		if input == "1" {
			editFile(len(readFile(true)))
		} else if input == "2" {
			readFile(false)
			fmt.Println("\nDrücken Sie \"Enter\" um zurrück zu gehen:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
		} else if input == "3" {
			checkOff(readFile(true))
		} else if input == "4" {
			createFile()
		} else if input == "5" {
			break
		}
	}

}

// Funktion zum erstellen einer neuen Datei mit dem Namen "ToDo.txt"
func createFile() {
	file, err := os.Create("ToDo.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
}

// Funktion zum lesen der Datei "ToDo.txt". Sollte der Übergabeparameter auf true gesetzt werden, so wird lediglich das Slice zurrückgegeben ohne eine Ausgabe
func readFile(arrayOnly bool) []string {
	a := []string{}
	file, err := os.Open("ToDo.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	fmt.Println()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if !arrayOnly {
			fmt.Println(scanner.Text())
		} else {
			a = append(a, scanner.Text())
		}
	}
	return a
}

// Überprüfung um zu gucken ob es schon eine To-Do-Liste gibt
func checkIfExists() bool {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, f := range files {
		if f.Name() == "ToDo.txt" {
			return true
		}
	}
	return false
}

// Funktion um einen neuen Eintrag in die To-Do-Liste einzufügen. Der übergabeparameter wird genutzt um die Einträge zu numerieren
func editFile(count int) {
	readFile(false)
	fmt.Println("\nGeben Sie \"0\" zum abbrechen ein")
	fmt.Printf("-> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "0" || input == "" {
		return
	} else {
		file, err := os.OpenFile("ToDo.txt", os.O_APPEND, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		input = strconv.Itoa(count+1) + ") " + input + "\n"
		file.WriteString(input)
	}
}

// Funktion die dazu da ist einen Eintrag abzuhaken. Die To-Do-Liste wird als Slice übergeben um den N-ten Eintrag abzuhaken. Zuletzt wird die Datei überschrieben
func checkOff(arr []string) {
	readFile(false)
	fmt.Printf("\nWelchen eintrag wollen Sie abhaken?: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	number, _ := strconv.Atoi(input)
	arr[number-1] = arr[number-1] + "✔️"

	file, _ := os.Create("ToDo.txt")
	defer file.Close()
	for i := 0; i < len(arr); i++ {
		file.WriteString(arr[i] + "\n")
	}
}
