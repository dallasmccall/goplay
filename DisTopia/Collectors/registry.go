package Collectors

import (
	"github.com/dallasmccall/goplay/proto/dataset"
	"github.com/dallasmccall/goplay/proto/schema"
	"log"
)

type Collector interface {
	GetSchema() *schema.Schema
	GetData() (*dataset.DataSet, error)
}

var Collectors = map[string]Collector{}
var ids = map[uint32]string{}

func Register(collector Collector) {
	cs := collector.GetSchema()
	name, ok := ids[cs.Id]
	if ok && name != cs.Name {
		log.Fatalln("ID ", cs.Id, " is in use for both ", cs.Name, " and ", name)
	}

	ids[cs.Id] = cs.Name
	Collectors[cs.Name] = collector
}
