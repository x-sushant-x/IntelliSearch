package utils

import "net/url"

func GetURLHostName(link string) (string, error) {
	URL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	return URL.Host, nil
}
