package main

import (
	"fmt"

	"github.com/ultraderek/sttserver/llmmain"
	"github.com/ultraderek/sttserver/voice"
)

func main() {
	response := llmmain.Example("how much wood would a wood chuck chuck if a wood chuck could chuck wood?")
	fmt.Println(response)
	voice.Example(response)
	//voice.Example("I am a lumber jack and i'm ok")
	//mic.Example1()

}
