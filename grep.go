package main

import "os"
import "fmt"
import "io/ioutil"

func main() {
	validate()
  //grep -UPc "\r$" file
	content, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	count := 0
	for i := 0; i < len(content) - 1; i++ {
		if content[i] == 0xD && content[i + 1] == 0xA {
				//count dos line ending
				count++
		}
	}
	fmt.Println(count)
}

func validate() {
	if len(os.Args) != 4 || os.Args[1] != "-UPc" || os.Args[2] != "\\r$" {
		fmt.Println("Supports only \"grep -UPc \"\\r$\" file\" operation")
		os.Exit(1)
	}
}
