# Ollama RAG
Project for adding RAG to ollama models using :
- [langchaingo](https://github.com/tmc/langchaingo)
- [qdrant vector database](https://github.com/qdrant/qdrant)
- [nomic-embed-text](https://ollama.com/library/nomic-embed-text)
- [Mistral](https://ollama.com/library/mistral)

## Installation :
Download packages : 
```bash
go mod download
```
Pull Ollama models :
```
ollama pull nomic-embed-text
ollama pull mistral
```

Install qdrant : 
```bash
docker pull qdrant/qdrant
docker run -p 6333:6333 qdrant/qdrant
```
Create Collection for qdrant
use any http client for make a PUT request like example blow for creating a Collection 

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

## Using

put you texts in text.txt and run :
``` go run . ```
then ask anything from the text you provided

-------
### TODOS:
- [ ] adding http request to make collection with functions
- [ ] dynamic file selection
- [ ] check if some file is embedded and already exist in vector database
