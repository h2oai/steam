package db

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
)

var bucketMap map[string]bool

func init() {
	bucketMap = make(map[string]bool)
	for _, b := range Buckets {
		bucketMap[b] = true
	}
}

func isValidBucket(b string) bool {
	_, ok := bucketMap[b]
	return ok
}

type DS struct {
	db       *bolt.DB
	location string
}

func Open(location string) (*DS, error) {
	db, err := bolt.Open(location, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("Datastore open failed: %v", err)
	}
	return &DS{db, location}, nil
}

func (ds *DS) Close() error { //TODO call from exit
	return ds.db.Close()
}

func (ds *DS) HasBucket(name string) (bool, error) {
	has := false
	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(name))
		has = b != nil
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("Bucket %s check failed: %v", name, err)
	}
	return has, nil
}

func (ds *DS) CreateBuckets(names []string) error {
	err := ds.db.Update(func(tx *bolt.Tx) error {
		for _, name := range names {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return fmt.Errorf("%s: %v", name, err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Bucket creation failed: %v", err)
	}
	return nil
}

func (ds *DS) Dump() error {
	return ds.db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Printf("\n\n===== Bucket: %s =====\n", name)
			return b.ForEach(func(k, v []byte) error {
				fmt.Printf("----- %v -----", k)
				printer, ok := Printers[string(name)]
				if !ok {
					return fmt.Errorf("Printer not found for %s", name)
				}
				s, err := printer(v)
				if err != nil {
					return err
				}
				fmt.Println(s)
				return nil
			})
		})
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
