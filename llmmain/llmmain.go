package llmmain

import (
	"context"
	"log"

	"github.com/teilomillet/gollm"
)

type LLM struct {
	llm      gollm.LLM
	behavior gollm.PromptOption
}

func Inittest() (LLM, error) {
	llm, err := gollm.NewLLM(
		gollm.SetProvider("ollama"),
		gollm.SetModel("llama3.1"),
		//gollm.SetMemory(12000),
		gollm.SetMaxTokens(250),
		gollm.SetDebugLevel(gollm.LogLevelWarn),
	)
	x := LLM{
		llm: llm,
		behavior: gollm.WithDirectives(
			"don't be chatty",
			"be a bit bratty",
			"you show subtle signs of being a witch",
			"your name is Hexia Hysteria",
			"use a couple of sentences",
			"don't use actions",
		),
	}
	if err != nil {
		return x, err
	}
	return x, nil
}
func Example(mybot *LLM, question string) string {

	// Create a prompt using NewPrompt function
	//prompt := gollm.NewPrompt(question, mybot.behavior)
	prompt := gollm.NewPrompt(question)

	// Generate a response
	ctx := context.Background()
	response, err := mybot.llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate response: %v", err)
	}

	//fmt.Printf("Response: %s\n", response)
	return response
}
