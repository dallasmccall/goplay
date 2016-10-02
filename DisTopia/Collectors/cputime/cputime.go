package cputime

import (
	"github.com/dallasmccall/goplay/DisTopia/Collectors"
	"github.com/dallasmccall/goplay/proto/dataset"
	"github.com/dallasmccall/goplay/proto/schema"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

var instance Collectors.Collector

func init() {
	instance = CPUCollector{
		CPUSchema: &schema.Schema{
			Name:        "CPU Time",
			Id:          1,
			Version:     1,
			Topic:       "/System/CPU/Time",
			Description: "CPU Time",

			Value: []*schema.Schema_Value{
				&schema.Schema_Value{Name: "CPU", Description: "CPU Name", DataType: schema.Schema_STRING, IsKey: true},
				&schema.Schema_Value{Name: "Total", Description: "Total CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "User", Description: "User CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "System", Description: "System CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Idle", Description: "Idle CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Nice", Description: "Nice CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "I/O Wait", Description: "I/O Wait CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "IRQ", Description: "IRQ CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Soft IRQ", Description: "Soft IRQ CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Steal", Description: "Steal CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Stolen", Description: "Stolen CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Guest", Description: "Guest CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
				&schema.Schema_Value{Name: "Guest Nice", Description: "Guest Nice CPU Time", DataType: schema.Schema_DOUBLE, Type: schema.Schema_COUNTER},
			},
		},
	}
}

type CPUCollector struct {
	CPUSchema *schema.Schema
}

func (collector CPUCollector) GetSchema() *schema.Schema {
	return collector.CPUSchema
}

func (collector CPUCollector) GetData() (*dataset.DataSet, error) {
	times, err := cpu.Times(true)
	if nil != err {
		return nil, err
	}

	data := &dataset.DataSet{
		SchemaId:  collector.CPUSchema.Id,
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
		Row:       make([]*dataset.DataSet_Row, len(times)),
	}

	for i, time := range times {
		data.Row[i] = &dataset.DataSet_Row{
			Column: []*dataset.DataSet_Column{
				&dataset.DataSet_Column{StringValue: time.CPU},
				&dataset.DataSet_Column{DoubleValue: time.Total() - time.Idle},
				&dataset.DataSet_Column{DoubleValue: time.User},
				&dataset.DataSet_Column{DoubleValue: time.System},
				&dataset.DataSet_Column{DoubleValue: time.Idle},
				&dataset.DataSet_Column{DoubleValue: time.Nice},
				&dataset.DataSet_Column{DoubleValue: time.Iowait},
				&dataset.DataSet_Column{DoubleValue: time.Irq},
				&dataset.DataSet_Column{DoubleValue: time.Softirq},
				&dataset.DataSet_Column{DoubleValue: time.Steal},
				&dataset.DataSet_Column{DoubleValue: time.Stolen},
				&dataset.DataSet_Column{DoubleValue: time.Guest},
				&dataset.DataSet_Column{DoubleValue: time.GuestNice},
			},
		}
	}

	return data, nil
}
