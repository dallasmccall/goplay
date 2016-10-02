package diskusage

import (
	_ "github.com/dallasmccall/goplay/DisTopia/Collectors"
	"github.com/dallasmccall/goplay/proto/dataset"
	"github.com/dallasmccall/goplay/proto/schema"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestGetSchema(t *testing.T) {
	data := instance.GetSchema()
	bytes, err := proto.Marshal(data)
	if nil != err {
		t.Error("Could not marshall Disk Schema: ", err)
	}
	log.Println("CPU Schema Size: ", len(bytes))
	data = &schema.Schema{}
	err = proto.Unmarshal(bytes, data)
	if nil != err {
		t.Error("Could not unmarshall Disk Schema: ", err)
	}
	//	log.Println(Collectors.ToJSON(data))
}

func TestGetData(t *testing.T) {
	data, err := instance.GetData()
	if nil != err {
		t.Error("Could not get Disk Data: ", err)
	}
	bytes, err := proto.Marshal(data)
	if nil != err {
		t.Error("Could not marshall Disk Data: ", err)
	}
	log.Println("CPU Data Size: ", len(bytes))
	data = &dataset.DataSet{}
	err = proto.Unmarshal(bytes, data)
	if nil != err {
		t.Error("Could not unmarshall Disk Data: ", err)
	}
	//	log.Println(Collectors.ToJSON(data))
}

func BenchmarkGetSchema(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instance.GetSchema()
	}
}

func BenchmarkGetData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instance.GetData()
	}
}
