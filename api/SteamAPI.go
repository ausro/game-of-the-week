package api

import (
	"fmt"
	"io"
	"net/http"
)

type ApiPath int

const (
	Host = "https://store.steampowered.com"

	GetApps ApiPath = iota
	GetDetails
)

var apiPath = map[ApiPath]string{
	GetApps:    "/api/getappsincategory/",
	GetDetails: "/api/appdetails",
}

func (path ApiPath) String() string {
	return apiPath[path]
}

func GET(apiType ApiPath, args string) ([]byte, error) {
	var result []byte
	var url = fmt.Sprintf("%s%s?%s", Host, apiType.String(), args)

	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}

	if resp != nil {
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return result, err
		}

		result = body
	}

	return result, err
}
