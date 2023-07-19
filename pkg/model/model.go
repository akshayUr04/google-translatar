package model

type Translate struct {
	SourceLanguage string `json:"source"`
	TragetLanguage string `json:"target"`
	Text           string `json:"text"`
}
