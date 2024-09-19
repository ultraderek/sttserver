package llmmain

import (
	"context"
	"log"

	"github.com/teilomillet/gollm"
)

func Example(question string) string {
	// Create a new LLM instance with Ollama provider
	llm, err := gollm.NewLLM(
		gollm.SetProvider("ollama"),
		gollm.SetModel("llama3.1"),

		gollm.SetDebugLevel(gollm.LogLevelWarn),
	)
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}

	// Create a prompt using NewPrompt function
	prompt := gollm.NewPrompt(question)

	// Generate a response
	ctx := context.Background()
	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate response: %v", err)
	}

	//fmt.Printf("Response: %s\n", response)
	return response
}
