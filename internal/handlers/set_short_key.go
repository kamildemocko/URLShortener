package handlers

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
	"time"
	"urlshortener/internal/validators"
)

func (u *URLHandler) HandleSetShortKey(w http.ResponseWriter, r *http.Request) {
	inputUrl := r.FormValue("input-url")
	desiredKey := r.FormValue("desired-key")

	if err := validators.ValidateUrl(inputUrl); err != nil {
		u.HandleErrorPage(w, r, err.Error())
		return
	}

	if desiredKey == "" {
		for desiredKey == "" {
			randKey := generateRandomKey(8)
			exists, err := u.Config.Repository.KeyExists(randKey)
			if err != nil {
				u.HandleErrorPage(w, r, err.Error())
				return
			}
			if !exists {
				desiredKey = randKey
			}
		}
	}

	if err := validators.ValidateKey(desiredKey); err != nil {
		u.HandleErrorPage(w, r, err.Error())
		return
	}

	ip := GetIP(r)
	if len(ip) > 32 {
		ip = ip[:32]
	}

	err := u.Config.Repository.SetKey(time.Now(), ip, inputUrl, desiredKey)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "unique_key"):
			u.HandleErrorPage(w, r, "The provided key already exists")
		default:
			u.HandleErrorPage(w, r, "Cannot save key at the moment, try again later")
			log.Println(err.Error())
		}

		return
	}

	newUrl := fmt.Sprintf("%s://%s/%s", os.Getenv("PROTOCOL"), os.Getenv("DOMAIN"), desiredKey)
	u.HandleSuccessPage(w, r, newUrl)
}

func generateRandomKey(l int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#^*")
	buffer := make([]rune, l)
	for i := 0; i < l; i++ {
		buffer[i] = runes[rand.IntN(len(runes))]
	}

	return string(buffer)
}
