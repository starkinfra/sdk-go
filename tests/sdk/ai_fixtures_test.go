package sdk

import (
	"fmt"
	"os"
	"sync"
	"testing"

	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/sdk-go/starkinfra"
	AiAgent "github.com/starkinfra/sdk-go/starkinfra/aiagent"
	AiChat "github.com/starkinfra/sdk-go/starkinfra/aichat"
	AiKnowledgeBase "github.com/starkinfra/sdk-go/starkinfra/aiknowledgebase"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
)

type aiFixtures struct {
	once          sync.Once
	knowledgeBase AiKnowledgeBase.AiKnowledgeBase
	agent         AiAgent.AiAgent
	chat          AiChat.AiChat
	failure       string
}

var fixtures = &aiFixtures{}

func TestMain(m *testing.M) {
	code := m.Run()
	if !fixtures.cleanup() && code == 0 {
		code = 1
	}
	os.Exit(code)
}

func (f *aiFixtures) load(t *testing.T) {
	starkinfra.User = utils.ExampleProject
	f.once.Do(func() {
		knowledgeBase, errs := AiKnowledgeBase.Create(Example.ExampleAiKnowledgeBase(), nil)
		if errs.Errors != nil {
			f.failure = fmt.Sprintf("knowledge base: code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
			return
		}
		f.knowledgeBase = knowledgeBase
		agent, errs := AiAgent.Create(Example.ExampleAiAgent([]string{knowledgeBase.Id}), nil)
		if errs.Errors != nil {
			f.failure = fmt.Sprintf("agent: code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
			return
		}
		f.agent = agent
		chat, errs := AiChat.Create(Example.ExampleAiChat(agent.Id), nil)
		if errs.Errors != nil {
			f.failure = fmt.Sprintf("chat: code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
			return
		}
		f.chat = chat
	})
	if f.failure != "" {
		t.Fatalf("could not create the shared AI fixtures: %s", f.failure)
	}
}

func (f *aiFixtures) cleanup() bool {
	clean := true
	if f.chat.Id != "" {
		_, err := AiChat.Delete([]string{f.chat.Id}, nil)
		clean = reportDelete("AiChat", f.chat.Id, err.Errors) && clean
	}
	if f.agent.Id != "" {
		_, err := AiAgent.Delete([]string{f.agent.Id}, nil)
		clean = reportDelete("AiAgent", f.agent.Id, err.Errors) && clean
	}
	if f.knowledgeBase.Id != "" {
		_, err := AiKnowledgeBase.Delete([]string{f.knowledgeBase.Id}, nil)
		clean = reportDelete("AiKnowledgeBase", f.knowledgeBase.Id, err.Errors) && clean
	}
	return clean
}

func reportDelete(name string, id string, errs []Error.StarkError) bool {
	if errs == nil {
		return true
	}
	fmt.Fprintf(os.Stderr, "%s %s was not deleted: code: %s, message: %s\n", name, id, errs[0].Code, errs[0].Message)
	return false
}
