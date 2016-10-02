package diskusage

import (
	"github.com/dallasmccall/goplay/DisTopia/Collectors"
	"github.com/dallasmccall/goplay/proto/dataset"
	"github.com/dallasmccall/goplay/proto/schema"
	"github.com/shirou/gopsutil/disk"
	"time"
)

var instance Collectors.Collector

func init() {
	instance = DiskCollector{
		DiskSchema: &schema.Schema{
			Name:        "Disk Usage",
			Id:          2,
			Version:     1,
			Topic:       "/System/Disk/Usage",
			Description: "Disk Usage",

			Value: []*schema.Schema_Value{
				&schema.Schema_Value{Name: "Path", Description: "Filesystem Path", DataType: schema.Schema_STRING, IsKey: true},
				&schema.Schema_Value{Name: "Used", Description: "Used Disk Space", DataType: schema.Schema_LONG, Type: schema.Schema_GAGUE},
				&schema.Schema_Value{Name: "Free", Description: "Free Disk Space", DataType: schema.Schema_LONG, Type: schema.Schema_GAGUE},
				&schema.Schema_Value{Name: "Total", Description: "Total Disk Space", DataType: schema.Schema_LONG, Type: schema.Schema_GAGUE},
				&schema.Schema_Value{Name: "Inodes Used", Description: "Used inodes", DataType: schema.Schema_LONG, Type: schema.Schema_GAGUE},
				&schema.Schema_Value{Name: "Inodes Free", Description: "Free inodes", DataType: schema.Schema_LONG, Type: schema.Schema_GAGUE},
				&schema.Schema_Value{Name: "Inodes Total", Description: "Total inodes", DataType: schema.Schema_LONG, Type: schema.Schema_GAGUE},
			},
		},
	}
}

type DiskCollector struct {
	DiskSchema *schema.Schema
}

func (collector DiskCollector) GetSchema() *schema.Schema {
	return collector.DiskSchema
}

func (collector DiskCollector) GetData() (*dataset.DataSet, error) {
	partitions, err := disk.Partitions(false)
	if nil != err {
		return nil, err
	}

	data := &dataset.DataSet{
		SchemaId:  collector.DiskSchema.Id,
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
		Row:       make([]*dataset.DataSet_Row, len(partitions)),
	}

	for i, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if nil != err {
			return nil, err
		}
		data.Row[i] = &dataset.DataSet_Row{
			Column: []*dataset.DataSet_Column{
				&dataset.DataSet_Column{StringValue: usage.Path},
				&dataset.DataSet_Column{LongValue: int64(usage.Used)},
				&dataset.DataSet_Column{LongValue: int64(usage.Free)},
				&dataset.DataSet_Column{LongValue: int64(usage.Total)},
				&dataset.DataSet_Column{LongValue: int64(usage.InodesUsed)},
				&dataset.DataSet_Column{LongValue: int64(usage.InodesFree)},
				&dataset.DataSet_Column{LongValue: int64(usage.InodesTotal)},
			},
		}
	}

	return data, nil
}
