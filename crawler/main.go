package main

import (
	"github.com/wangning0/crawler/engine"
	"github.com/wangning0/crawler/zhenai/parser"
	"github.com/wangning0/crawler/scheduler"
	"github.com/wangning0/crawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
