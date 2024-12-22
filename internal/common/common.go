package common

import (
	"fmt"
	"os"
)

func GetURI() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func mapLinkWithMethod(link, method string) map[string]string {
	return map[string]string{"href": link, "method": method}
}

func GenerateLinks(uri string, endpoint string, id string, currentPage int64) map[string]map[string]string {
	result := make(map[string]map[string]string)
	result["self"] = mapLinkWithMethod(uri+"/"+endpoint+"/"+id, "GET")
	result["collection"] = mapLinkWithMethod(uri+"/"+endpoint, "GET")
	if currentPage > 0 {
		result["next"] = mapLinkWithMethod(
			fmt.Sprintf("%s/%s?page=%d", uri, endpoint, currentPage+1),
			"GET",
		)
		result["prev"] = mapLinkWithMethod(
			fmt.Sprintf("%s/%s?page=%d", uri, endpoint, currentPage-1),
			"GET",
		)
	}

	return result
}
