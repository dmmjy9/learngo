package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/reals"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string {
			"name": "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
	r = &reals.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}


	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inpsecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *reals.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
