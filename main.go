package main

import (
	"demo/chatbotDemo"
	"flag"
	"fmt"
	"os"
)

func main() {
	jsonFlag := flag.Bool("json", false, "Start ChatBot which store data and return in JSON format.")
	streamFlag := flag.Bool("stream", false, "Start ChatBot with streaming response.")
	flag.Parse()

	if *jsonFlag && *streamFlag {
		fmt.Printf("Error: Cannot execute JSON and Stream at the same time.")
		os.Exit(1)
	}

	switch {
	case *jsonFlag:
		chatbotDemo.ChatJSON()
		chatbotDemo.ShowJSONFromDB()
	case *streamFlag:
		chatbotDemo.ChatStream()
	default:
		chatbotDemo.ChatStream()
	}
}
