package main

import (
  "fmt"
  "hash_example/hash_table"
  "os"
)

func main() {
	hashTable := hash_table.New()

	hashTable.Set("lang", "GoLang")
	hashTable.Set("programmer", "Ozeias")

	programmer := hashTable.Get("programmer")

	if programmer == nil {
		fmt.Println("Programmer is null")
	} else {
		fmt.Println("Programmer:", *programmer)
	}

	lang := hashTable.Get("lang")

	if lang == nil {
		fmt.Println("Lang is null")
	} else {
		fmt.Println("Lang:", *lang)
	}

  if len(os.Args) > 1 {
    serverString := os.Args[1]
    hashTable.Set("server", serverString)
  }

	server := hashTable.Get("server")

	if server == nil {
		fmt.Println("Server is null")
	} else {
		fmt.Println("Server:", *server)
	}
}
