package tax_id_generator

import (
	AiMessage "github.com/starkinfra/sdk-go/starkinfra/aimessage"
)

func ExampleAiMessage(chatId string) AiMessage.AiMessage {
	return AiMessage.AiMessage{
		ChatId: chatId,
		Text:   "Say hello and mention order 123.",
	}
}
