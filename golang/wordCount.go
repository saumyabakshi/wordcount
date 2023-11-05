package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func countBytes(filename []byte) (int) {
	
	size := len(filename)
	return int(size)

}

func countLines(filename []byte)(int) {

	count := 0
	for _, char := range filename {
		if char == '\n' {
			count++
		}
	
	}
	return count

}

func numWords(filename []byte) (int){
	text := string(filename)
	words := len(strings.Fields(text))
	return words
}

func countChars(filename []byte) (int,) {
	
	return len(bytes.Runes(filename))

}

func main() {
		Chars, Words, Lines, Bytes := false, false, false, false
		flag.BoolVar(&Chars, "m", false, "print the character counts")
		flag.BoolVar(&Words, "w", false, "print the word counts")
		flag.BoolVar(&Lines, "l", false, "print the newline counts")
		flag.BoolVar(&Bytes, "c", false, "print the byte counts")
		flag.Parse()
		file := ""
		var filename []byte

		if !Chars && !Words && !Lines && !Bytes && len(os.Args) == 2 {
			var err error
			filename,err = os.ReadFile(os.Args[1])
			file = os.Args[1]
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else if len(os.Args) > 2 {
			var err error
			filename,err = os.ReadFile(os.Args[2])
			file = os.Args[2]
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			pipedInput, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filename = pipedInput
		}

		
		if Chars {
			count := countChars(filename)
			fmt.Printf("%d %s\n",count,file)
			
		} else if Words {
			count := numWords(filename)
			fmt.Printf("%d %s\n",count,file)
		} else if Lines {
			count := countLines(filename)
			fmt.Printf("%d %s\n",count,file)
		} else if Bytes {
			count := countBytes(filename)
			fmt.Printf("%d %s\n",count,file)
		} else {
			numBytes := countBytes(filename)
			numLines  := countLines(filename)
			numWords := numWords(filename)
			fmt.Printf("%d %d %d %s\n",numLines,numWords,numBytes,file)
		}
}