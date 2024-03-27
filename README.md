# Ollama RAG
Project for adding RAG to ollama models using :
- [langchaingo](https://github.com/tmc/langchaingo)
- [qdrant vector database](https://github.com/qdrant/qdrant)
- [nomic-embed-text](https://ollama.com/library/nomic-embed-text)

## Installation :
Download packages : 
```bash
go mod download
```
Install qdrant : 
```bash
docker pull qdrant/qdrant
docker run -p 6333:6333 qdrant/qdrant
```
Create Collection for qdrant
use any http client for make a PUT request like example blow for creating a Collection 
- [ ] adding http request to make collection with functions
```bash
curl -X PUT http://localhost:6333/collections/romeo \
  -H 'Content-Type: application/json' \
  --data-raw '{
    "vectors": {
      "size": 768,
      "distance": "Dot"
    }
  }'
```

### Using

put you texts in text.txt and run :
``` go run . ```
then ask anything from the text you provided
