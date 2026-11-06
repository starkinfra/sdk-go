package tax_id_generator

import (
	"crypto/rand"
	"encoding/hex"

	AiAgent "github.com/starkinfra/sdk-go/starkinfra/aiagent"
)

func ExampleAiAgent(knowledgeBaseIds []string) AiAgent.AiAgent {
	content := make([]byte, 6)
	rand.Read(content)
	return AiAgent.AiAgent{
		Name:             "sdk-go-agent-" + hex.EncodeToString(content),
		Model:            "bender-1.0",
		SystemPrompt:     "Answer in one short sentence.",
		KnowledgeBaseIds: knowledgeBaseIds,
		MetadataSchema: map[string]interface{}{
			"order_id": map[string]interface{}{"type": "string", "description": "Order the customer mentions"},
		},
	}
}
