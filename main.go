package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shayanfarzi/ollama-rag/pkg/document"
	"github.com/shayanfarzi/ollama-rag/utils"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	ollamaEmbedderModel, err := ollama.New(ollama.WithModel("nomic-embed-text:latest"))
	if err != nil {
		log.Fatal(err)
	}
	ollamaEmbedder, err := embeddings.NewEmbedder(ollamaEmbedderModel)
	if err != nil {
		log.Fatal(err)
	}

	docs := document.TextToChunks("./test.txt")

	prompt, err := utils.GetUserInput("Enter a question:")
	if err != nil {
		log.Fatal(err)
	}
	store := useStorage(docs, ollamaEmbedder)

	docRetrieved := useRetriaver(store, prompt)

	history := memory.NewChatMessageHistory()

	for _, doc := range docRetrieved {
		history.AddAIMessage(ctx, doc.PageContent)
	}

	conversation := memory.NewConversationBuffer(memory.WithChatHistory(history))

	executor, err := agents.Initialize(
		llm,
		nil,
		agents.ConversationalReactDescription,
		agents.WithMemory(conversation),
	)

	if err != nil {
		fmt.Println("Error initializing agents", err)
		return
	}

	options := []chains.ChainCallOption{
		chains.WithTemperature(0.8),
	}
	res, err := chains.Run(ctx, executor, prompt, options...)

	if err != nil {
		fmt.Println("Error running chains", err)
		return
	}

	fmt.Println(res)
}
