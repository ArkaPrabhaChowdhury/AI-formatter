package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const geminiAPIURL = "https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateText"

type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiContent struct {
	Role  string `json:"role"`
	Parts []struct {
		Text string `json:"text"`
	} `json:"parts"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func FormatDataWithGemini(inputText, formatType string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("Gemini API key is missing")
	}

	prompt := fmt.Sprintf("Format this data as %s: %s", formatType, inputText)

	requestBody := GeminiRequest{
		Contents: []GeminiContent{
			{
				Role: "user",
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s?key=%s", geminiAPIURL, apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var geminiResponse GeminiResponse
	err = json.Unmarshal(body, &geminiResponse)
	if err != nil {
		return "", err
	}

	if len(geminiResponse.Candidates) > 0 &&
		len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		return geminiResponse.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no valid response from Gemini API")
}
