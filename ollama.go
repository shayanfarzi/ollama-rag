package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func useStorage(docs []schema.Document, embedder *embeddings.EmbedderImpl) *qdrant.Store {
	qdrantUrl, err := url.Parse("http://localhost:6333")
	if err != nil {
		log.Fatalf("failed parsing url: %s", err)
	}
	qdrant.New()
	store, err := qdrant.New(
		qdrant.WithURL(*qdrantUrl),
		qdrant.WithAPIKey(""),
		qdrant.WithCollectionName("romeo"),
		qdrant.WithEmbedder(embedder),
	)
	if err != nil {
		fmt.Println("Qdrant creation failed:", err)
		return nil
	}

	if len(docs) == 0 {
		_, err = store.AddDocuments(context.Background(), docs)
		if err != nil {
			fmt.Println("Error adding documents", err)
			return nil
		}
	}

	return &store
}

func useRetriaver(store *qdrant.Store, prompt string) []schema.Document {
	optionsVector := []vectorstores.Option{
		vectorstores.WithScoreThreshold(0.80),
	}

	retriever := vectorstores.ToRetriever(store, 5, optionsVector...)
	// search
	docRetrieved, err := retriever.GetRelevantDocuments(context.Background(), prompt)

	if err != nil {
		fmt.Println("Error getting relevant documents", err)
		return nil
	}

	return docRetrieved
}
