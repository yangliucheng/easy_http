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

	// /v2/apps/:app_id map[app_id:/super/admin]
	for k, v := range params {
		//判断v是不是以/开始
		if !StartWith(v, `/`) {
			k = StringJoin(`/:`, k)
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