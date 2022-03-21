package main

import (
	"bufio"   //scanner
	"fmt"     //print
	"io"      //reader
	"log"     //fatalln for error
	"os"      //open
	"sort"    //sort
	"strings" //tolower
	"unicode" //first char
)

var totalWords int

func sorted(conc map[string]int) []string { //this sorts the strings from map count from function Concordance from File which is shakespear.txt
	keys := make([]string, 0, len(conc)) //makes the keys

	for k := range conc { //keys for all of conc
		keys = append(keys, k)
	}
	sort.Strings(keys) //sort all keys from conc

	return keys //returns all keys sorted
}
func Concordance(File io.Reader) map[string]int {
	conc := map[string]int{}          //sets map
	scanner := bufio.NewScanner(File) //sets scanner

	scanner.Split(bufio.ScanWords) // only reads words
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text()) //turns all words to lowercase
		if unicode.IsLetter(rune(word[0])) {    //if the first char in go is rune is a letter it will be added to the map and then the int will be increased by one
			conc[word]++
			totalWords++ //test the total amount of words
		}

	}
	return conc
}

func main() {
	totalWords = 0
	File, err := os.Open("shakespeare.txt") //this opens shakespear.txt
	if err != nil {
		log.Fatalln(err) //if unable to open file log the error
	}

	defer File.Close() //this Closes shakespear.txt

	count := Concordance(File)
	keys := sorted(count)
	/**													this sorts the strings then uses the keys to call concordance map with that key
		for _, k := range keys {
			fmt.Println(k, count[k])
		}
	**/

	fmt.Println(count)             //map prints out sorted in newer versions of go
	fmt.Println("first 10 words:") //first 10 words
	for i := 0; i < 10; i++ {
		fmt.Println(keys[i], count[keys[i]])
	}
	fmt.Println("last 10 words:") //last 10 words
	for i := 0; i < 10; i++ {
		fmt.Println(keys[len(count)+i-10], count[keys[len(count)+i-10]])
	}
}
