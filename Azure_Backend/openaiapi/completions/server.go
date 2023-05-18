package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/completions", helloHandler)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestData struct {
		Prompt string `json:"prompt"  binding:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, `{"error": "`+fmt.Sprintf("failed to parse request body: %v", err)+`"}`, http.StatusInternalServerError)
		return
	}

	textCompletionRequest := TextCompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      requestData.Prompt,
		MaxTokens:   50,
		Temperature: 0.6,
	}

	jsonRequestData, err := json.Marshal(textCompletionRequest)
	if err != nil {
		http.Error(w, `{"error": "`+fmt.Sprintf("failed to marshal request data: %v\n", err)+`"}`, http.StatusInternalServerError)
		return
	}

	url := "https://api.openai.com/v1/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestData))
	if err != nil {
		http.Error(w, `{"error": "`+fmt.Sprintf("failed to marshal request data: %v\n", err)+`"}`, http.StatusInternalServerError)
		return
	}

	bearer := getOpenAIBearerToken()
	if bearer == "" {
		http.Error(w, `{"error": "Bearer token not found"}`, http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)

	// Send the request using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, `{"error": "failed to send request"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	completions := TextCompletionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&completions)
	if err != nil {
		http.Error(w, `{"error": "failed to read response body"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(completions)
}

func getOpenAIBearerToken() string {
	return os.Getenv("OPENAI_BEARER_TOKEN")
}

type TextCompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type TextCompletionResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finishReason"`
}

type Usage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
}
