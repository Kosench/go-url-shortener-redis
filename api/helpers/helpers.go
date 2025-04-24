package helpers

import (
	"net/url"
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if len(url) < 4 || (!strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://")) {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(inputURL string) bool {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		// Если DOMAIN не установлен, считаем, что проверка не требуется
		return true
	}

	// Парсим URL для безопасной работы с его компонентами
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		// Если URL некорректный, считаем его недопустимым
		return false
	}

	// Удаляем префиксы "www." и сравниваем хост с DOMAIN
	host := strings.TrimPrefix(parsedURL.Hostname(), "www.")
	return host != domain
}

//func RemoveDomainError(url string) bool {
//	if url == os.Getenv("DOMAIN") {
//		return false
//	}
//
//	newURL := strings.Replace(url, "http://", "", 1)
//	newURL = strings.Replace(newURL, "http://", "", 1)
//	newURL = strings.Replace(newURL, "www.", "", 1)
//	newURL = strings.Split(newURL, "/")[0]
//
//	if newURL == os.Getenv("DOMAIN") {
//		return false
//	}
//
//	return true
//}
