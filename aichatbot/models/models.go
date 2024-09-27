package models

import "time"

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}
type Prompt struct {
	Prompt string `json:"prompt"`
}

type PromptResponse struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Prompt    string    `json:"prompt"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
}
