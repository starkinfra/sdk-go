package tax_id_generator

import (
	"crypto/rand"
	"encoding/hex"

	AiKnowledgeBase "github.com/starkinfra/sdk-go/starkinfra/aiknowledgebase"
)

func ExampleAiKnowledgeBase() AiKnowledgeBase.AiKnowledgeBase {
	content := make([]byte, 6)
	rand.Read(content)
	isRecursive := false
	return AiKnowledgeBase.AiKnowledgeBase{
		Name:        "sdk-go-" + hex.EncodeToString(content),
		RootUrl:     "https://docs.starkinfra.com",
		IsRecursive: &isRecursive,
		Tags:        []string{"sdk-go", "test"},
	}
}
