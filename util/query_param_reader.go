package util

import "net/url"

func QueryParams(queryValues url.Values) map[string]interface{} {
	queryParams := make(map[string]interface{})
	for k, v := range queryValues {
		queryParams[k] = v[0]
	}
	return queryParams
}
