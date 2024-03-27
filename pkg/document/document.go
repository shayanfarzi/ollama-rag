package document

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func TextToChunks(dirFile string) []schema.Document {
	file, err := os.Open(dirFile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return nil
	}

	docLoaded := documentloaders.NewText(file)

	split := textsplitter.NewRecursiveCharacter()
	split.ChunkSize = 768
	split.ChunkOverlap = 64
	docs, err := docLoaded.LoadAndSplit(context.Background(), split)
	if err != nil {
		log.Fatalf("failed splitting text: %s", err)
	}

	log.Println("Document loaded:", len(docs))

	return docs
}

func HtmlToChunks(dirFile string) []schema.Document {
	file, err := os.Open(dirFile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return nil
	}

	docLoaded := documentloaders.NewHTML(file)

	split := textsplitter.NewRecursiveCharacter()
	split.ChunkSize = 768
	split.ChunkOverlap = 64
	docs, err := docLoaded.LoadAndSplit(context.Background(), split)
	if err != nil {
		log.Fatalf("failed splitting text: %s", err)
	}

	log.Println("Document loaded:", len(docs))

	return docs
}

func PdfToChunks(dirFile string) []schema.Document {
	file, err := os.Open(dirFile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return nil
	}
	fInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return nil
	}

	docLoaded := documentloaders.NewPDF(file, fInfo.Size())

	split := textsplitter.NewRecursiveCharacter()
	split.ChunkSize = 768
	split.ChunkOverlap = 64
	docs, err := docLoaded.LoadAndSplit(context.Background(), split)
	if err != nil {
		log.Fatalf("failed splitting text: %s", err)
	}

	log.Println("Document loaded:", len(docs))

	return docs
}
