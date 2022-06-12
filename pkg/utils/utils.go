package utils

import "net/http"

func CopyHeaders(dst http.Header, src http.Header) {
	for k, vv := range src {
		dst[k] = append(dst[k], vv...)
	}
}

// RemoveHeaders removes the header with the given names from the headers map.
func RemoveHeaders(headers http.Header, names ...string) {
	for _, h := range names {
		headers.Del(h)
	}
}
