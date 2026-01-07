package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	// 获取 api key
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
	apiKey := os.Getenv("OPENAI_API_KEY")

	fileName := "toutest.md"
	// 读取文件内容
	data, err := os.ReadFile(fmt.Sprintf("./original/%s", fileName))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
		panic(err)
	}

	// 读取规则
	rule, err := os.ReadFile("./config/rule.md")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
		panic(err)
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-5-nano",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(string(rule) + "\n\n" + string(data))},
	})

	if err != nil {
		log.Fatalf("Response New error: %v", err)
		panic(err.Error())
	}

	result := resp.OutputText()
	fmt.Println(result)

	if err = os.WriteFile(fmt.Sprintf("./result/%s", fileName), []byte(result), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
		panic(err)
	}
}
