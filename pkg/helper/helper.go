package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"strings"

	"github.com/akshayUr04/google-translator/pkg/model"
)

const (
	TranslateURL = "https://google-translate1.p.rapidapi.com/language/translate/v2"
)

type TranslationResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func Translate(obj model.Translate) (string, error) {
	str := ""
	str = str + "q=" + obj.Text
	str = str + "&target=" + obj.TragetLanguage
	str = str + "&source=" + obj.SourceLanguage

	payload := strings.NewReader(str) //the string is converted into strings.Reader type which implements the io.Reader interface

	req, err := http.NewRequest(http.MethodPost, TranslateURL, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded") // headers are additional information sent with the
	req.Header.Add("Accept-Encoding", "application/gzip")               // request to provide metadata about the request or to
	req.Header.Add("X-RapidAPI-Key", os.Getenv("GOOGLE_API_KEY"))       //specify certain requirements for the server.
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var translationResponse TranslationResponse
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		return "", err
	}

	// Extract the translated text
	if len(translationResponse.Data.Translations) > 0 {
		translatedText := translationResponse.Data.Translations[0].TranslatedText
		return translatedText, nil
	}

	return "", fmt.Errorf("no translated text found in the response")

	// return body, nil
}
