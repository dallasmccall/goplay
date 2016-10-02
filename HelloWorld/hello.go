package main

import (
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"math"
)

type Mapper func(cursor iterator.Iterator)

func main() {
	log.Println("Starting DB")
	db, err := leveldb.OpenFile("/share/0", nil)
	defer db.Close()
	if nil != err {
		log.Fatalln("Error Opening Database: ", err)
	}

	slice := &util.Range{
		Start: ToKey(0, 0),
		Limit: ToKey(math.MaxUint64, math.MaxUint64),
	}
	Map(db, slice, func(cursor iterator.Iterator) {
		time, id := FromKey(cursor.Key())
		value := string(cursor.Value())
		log.Println("From DB: Time: ", time, " ID: ", id, "\tValue: ", value)
	})

	key := ToKey(1234, 4321)
	value := []byte("Hello World")
	db.Put(key, value, nil)

	time, id := FromKey(key)

	log.Println("Time: ", time, " ID: ", id)
}

func Map(db *leveldb.DB, slice *util.Range, mapper Mapper) {
	cursor := db.NewIterator(slice, nil)
	defer cursor.Release()
	for cursor.Next() {
		mapper(cursor)
	}
}

func ToKey(time, id uint64) []byte {
	key := make([]byte, 16)
	binary.BigEndian.PutUint64(key, time)
	binary.BigEndian.PutUint64(key[8:], id)
	return key
}

func FromKey(key []byte) (time, id uint64) {
	time = binary.BigEndian.Uint64(key)
	id = binary.BigEndian.Uint64(key[8:])
	return time, id
}
