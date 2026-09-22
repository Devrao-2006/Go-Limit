package helper

import (
	"errors"
	"net/http"

	constants "github.com/Devrao-2006/Go-Limit/constants"
)

func GetIPfromheader(headers http.Header) (string, error) {
	var ip string

	for _, val := range constants.IP_HEADERS_PRIORITY_LIST {
		ip := headers.Get(val)

		if ip != "" {
			break
		}
	}

	if ip != "" {
		return ip, nil
	}

	return "", errors.New("ip address not found")
}