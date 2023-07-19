package helper

import (
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

func Translate(obj model.Translate) ([]byte, error) {
	str2 := ""
	str2 = str2 + "q=" + obj.Text
	str2 = str2 + "&target=" + obj.TragetLanguage
	str2 = str2 + "&source=" + obj.SourceLanguage

	fmt.Println("str2-------", str2)

	payload := strings.NewReader(str2)

	req, err := http.NewRequest(http.MethodPost, TranslateURL, payload)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("GOOGLE_API_KEY")) //api key
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
