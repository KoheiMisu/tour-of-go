package util

import (
	url2 "net/url"
)

func GetRootUrl(url string) (result string, err error) {
	u, err := url2.Parse(url)
	if err != nil {
		return result, err
	}

	result = u.Scheme + "::/" + u.Host
	return result, nil
}
