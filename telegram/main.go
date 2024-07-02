package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v3"
)

var (
	openaiAPIKey     string
	telegramBotToken string
	systemContext    string
	help             string
)

func getOpenAI(user, query string) (string, error) {
	c := openai.NewClient(openaiAPIKey)
	ctx := context.Background()

	systemQuery := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemContext,
	}

	message := []openai.ChatCompletionMessage{
		systemQuery,
		{
			Role:    openai.ChatMessageRoleUser,
			Content: query,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4o,
		Messages: message,
		User:     user,
	}

	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	content := resp.Choices[0].Message.Content

	return content, nil
}

func main() {

	openaiAPIKey = os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		log.Println("OPENAI_API_KEY environment variable is not set")
		return
	}

	telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken == "" {
		log.Println("TELEGRAM_BOT_TOKEN environment variable is not set")
		return
	}

	systemContextAux, err := os.ReadFile("context.txt")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Error reading context.txt: %v", err)
		}
		log.Println("context.txt not found")
	}
	systemContext = string(systemContextAux)

	helpAux, err := os.ReadFile("help.md")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Error reading help.md: %v", err)
		}
		log.Println("help.md not found")
	}
	help = string(helpAux)

	pref := tele.Settings{
		Token:     telegramBotToken,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/ask", func(c tele.Context) error {
		/*
			if c.Message().Private() {
				c.Reply("This command is only available in group chats")
				return nil
			}
		*/
		args := c.Message().Payload
		if args == "" {
			c.Reply("Usage: /ask <question>")
			return nil
		}

		question := c.Message().Payload
		answer, err := getOpenAI(c.Message().Chat.Username, question)
		if err != nil {
			c.Reply("Error: " + err.Error())
			return nil
		}

		return c.Reply(answer)
	})

	b.Handle("/help", func(c tele.Context) error {
		log.Println("Help command")
		return c.Reply(help)
	})

	log.Println("Bot started")
	b.Start()
}
