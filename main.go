package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/ultraderek/sttserver/llmmain"
	"github.com/ultraderek/sttserver/voice"
)

const testflag = true

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	llm, err := llmmain.Inittest()
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}

	for scanner.Scan() {
		mytext := scanner.Text()
		if strings.Contains(mytext, "/bye") {
			return
		}
		//mytext = "Classic tongue twister!\n\nThe answer, of course, is \"a woodchuck would chuck as much wood as a woodchuck could chuck if a woodchuck could chuck wood.\""
		//voice.Example2(mytext, testflag)

		response := llmmain.Example(&llm, mytext)
		//	fmt.Println("\n---------------\n" + response + "\n---------------\n")
		voice.Example(response, testflag)

	}

	//voice.Example("I am a lumber jack and i'm ok")
	//mic.Example1()

}
