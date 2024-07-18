package util

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// zombie code
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func IsUserIitDataValid(initDataRaw string) bool {
	if initDataRaw == "" {
		return false
	}
	token := strings.ToLower(os.Getenv("TOKEN_BOT"))
	authDate := time.Now()
	dataHash := initdata.Sign(map[string]string{
		"query_id": initDataRaw,
	}, token, authDate)

	return dataHash != ""

}

func GetDataInInitDataRaw(query string) (map[string]string, error) {
	type User struct {
		ID              int    `json:"id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Username        string `json:"username"`
		LanguageCode    string `json:"language_code"`
		AllowsWriteToPM bool   `json:"allows_write_to_pm"`
	}

	v, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	userJSON := v.Get("user")
	var user User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"username":      user.Username,
		"user_id":       fmt.Sprintf("%d", user.ID),
		"first_name":    user.FirstName,
		"last_name":     user.LastName,
		"chat_instance": v.Get("chat_instance"),
		"chat_type":     v.Get("chat_type"),
		"auth_date":     v.Get("auth_date"),
		"hash":          v.Get("hash"),
	}, nil
}
