package chatgpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Chat() {

	apikey := "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"
	config := openai.DefaultConfig(apikey)
	//proxyUrl, err := url.Parse("http://localhost:3213")
	//if err != nil {
	//	panic(err)
	//}
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyUrl),
	//}
	//config.HTTPClient = &http.Client{
	//	Transport: transport,
	//}

	client := openai.NewClientWithConfig(config)
	//prompt := "需要你进行五轮问题回答。我将提供一个问题，然后您需要与ChatGPT进行对话，进行连续五个回合的提问和回答。您的目标是创建一种流畅的对话，并尽可能多地获取有关问题的相关信息。问答过程中，请确保问题和回答之间的连贯性，以及适当的上下文转换。结果应该是一个具有知识有深度的对话，为目标受众提供有用的信息。请考虑受众的需求和背景，并尽可能涵盖各个角度。"
	content := "罗密欧与朱丽叶是一对互相深爱的情侣, 但是他们的父母不允许他们在一起, 于是他们产生了一段感人肺腑的情感故事, 以下是他们的故事和对话, 请生成一段一千字的脚本"
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "你现在是莎士比亚, 你总能输出非常高质量的剧本, 你需要根据用户的需要来生成一个情感丰富的剧本",
			},
			{
				Role:    "assistant",
				Content: "（灯光昏暗，舞台中央是一个老旧的阳台。朱丽叶站在阳台上，仰望满天繁星，表现出内心的忧虑和苦闷。）\n\n朱丽叶：（望着星空）星星掩盖你的光芒，如果它们知道",
			},
			{
				Role:    "user",
				Content: content,
			},
		},
		Temperature: 0,
		MaxTokens:   100,
		N:           1,
		//Functions: []openai.FunctionDefinition{
		//	{
		//		Name:        "chat",
		//		Description: "",
		//		Parameters:  "",
		//	},
		//},
	})
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message)
}

func Chat1() {

	apikey := "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"
	config := openai.DefaultConfig(apikey)
	//proxyUrl, err := url.Parse("http://localhost:3213")
	//if err != nil {
	//	panic(err)
	//}
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyUrl),
	//}
	//config.HTTPClient = &http.Client{
	//	Transport: transport,
	//}

	client := openai.NewClientWithConfig(config)
	userA := "小明"
	userB := "小芳"
	dialog := "以下是" + userA + "和" + userB + "今天的对话"
	dialog += "\n" + userA + "：你好"
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			dialog += "\n" + userA + "："
		} else {
			dialog += "\n" + userB + "："
		}
		resp, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
			Model:       openai.GPT3TextDavinci003,
			Prompt:      dialog,
			Temperature: 0,
			MaxTokens:   100,
			N:           1,
		})
		if err != nil {
			fmt.Printf("Completion error: %v\n", err)
			return
		}
		dialog += resp.Choices[0].Text
	}
	fmt.Println(dialog)
}
