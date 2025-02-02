package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func readPrompt() string {
	read := bufio.NewReader(os.Stdin)
	reader, _ := read.ReadString('\n')
	reader = strings.ReplaceAll(reader, "\n", "")

	return reader
}

func main() {
	// get entry from user
	fmt.Print("enter user name of github User : ")
	name := readPrompt()

	// fetch data from github api and show them
	url := fmt.Sprintf("https://api.github.com/users/%v/events", name)
	res, _ := http.Get(url)
	response, _ := io.ReadAll(res.Body)



	/*
	*
	* decode json to slice ( why slice ? )
	* because when we have something which doesnt have
	* multi type we should use map of interface but
	* if that was something like below ,must be a slice
	*
	**/
	var jsonDecode []interface{}
	err := json.Unmarshal(response, &jsonDecode)




	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	response, _ = json.MarshalIndent(jsonDecode, "", " ")
	fmt.Println(string(response))
}
