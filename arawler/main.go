package main

import (
	"learngo/arawler/engine"
	"learngo/arawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
