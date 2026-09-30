package tax_id_generator

import (
	"crypto/rand"
	"encoding/hex"

	AiChat "github.com/starkinfra/sdk-go/starkinfra/aichat"
)

func ExampleAiChat(agentId string) AiChat.AiChat {
	content := make([]byte, 6)
	rand.Read(content)
	return AiChat.AiChat{
		AgentId: agentId,
		Title:   "sdk-go-chat-" + hex.EncodeToString(content),
		Tags:    []string{"sdk-go", "test"},
		Context: map[string]interface{}{"name": "Ana", "balance": 1520.33},
	}
}
