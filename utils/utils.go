package utils

import (
	"bytes"
	"strings"
)

func StringJoin(str ...string) string {

	var buffer bytes.Buffer
	for _, s := range str {
		buffer.WriteString(s)
	}

	return buffer.String()
}

func ParaseUrlParam(url string, params map[string]string) string {
	// easy_http.Mapstring{"namespace": "default"}
	// /api/v1/namespaces/:namespace/pods
	for k, v := range params {
		if !StartWith(k, ":") {
			k = StringJoin(":",k)
		}
		url = strings.Replace(url, k, v, -1)
	}

	return url
}

func StartWith(s string, sep string) bool {

	if i := strings.Index(s, sep); i == 0 {
		return true
	}
	return false
}