package coqui

type Request struct {
	VoiceID string  `json:"voice_id"`
	Name    string  `json:"name,omitempty"`
	Text    string  `json:"text"`
	Speed   float64 `json:"speed,omitempty"`
	Prompt  string  `json:"prompt,omitempty"`
}

type Response struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
	AudioURL  string `json:"audio_url"`
	Emotion   string `json:"emotion"`
	Language  string `json:"language"`
}
