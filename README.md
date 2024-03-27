## Create Collection for qdrant

```bash
curl -X PUT http://localhost:6333/collections/storage_qdrant \
  -H 'Content-Type: application/json' \
  --data-raw '{
    "vectors": {
      "size": 300,
      "distance": "Cosine"
    }
  }'
```
