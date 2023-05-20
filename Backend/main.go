package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/ping", getPing)
	router.POST("/completions", getCompletions)
	router.Run(":8080")
}

// getAlbums responds with the list of all albums as JSON.
func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "Hello OK!!"})
}

// getAlbums responds with the list of all albums as JSON.
func getCompletions(c *gin.Context) {
	var requestData struct {
		Prompt string `json:"prompt"  binding:"required"`
	}

	// リクエストボディから Prompt を取得
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to parse request body: %v", err)})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to marshal request data: %v\n", err)})
		return
	}

	url := "https://api.openai.com/v1/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}

	bearer := getOpenAIBearerToken()
	if bearer == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bearer token not found"})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)

	// Send the request using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request"})
		return
	}
	defer resp.Body.Close()

	completions := TextCompletionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&completions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response body"})
		return
	}

	c.JSON(http.StatusOK, completions)
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
