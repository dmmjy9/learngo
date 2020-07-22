package parser

import (
	"fmt"
	"learngo/arawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch(
		"http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}


	result := ParseCityList(contents)
	fmt.Println(result)

	const resultSize = 438
	expectedUrls := []string{
		"", "", "",
	}
	expectedCities := []string{
		"", "", "",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}
