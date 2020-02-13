package bbolt

import (
	"errors"
	"time"

	bolt "go.etcd.io/bbolt"
)

// DB holds the database connection instance
type DB struct {
	db *bolt.DB
}

// New instantiate the storage DB and open the database file
func New(database string) (*DB, error) {
	s := &DB{}
	err := s.Open(database)
	return s, err
}

// Open the database
func (s *DB) Open(database string) error {
	db, err := bolt.Open(database, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	s.db = db
	err = s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("jobs"))
		return err
	})
	return err
}

// Close database
func (s *DB) Close() error {
	return s.db.Close()
}

// Update the value of the key in the bucket
func (s *DB) Update(bucket, key string, payload []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put([]byte(key), payload)
	})
}

// View get bucket and key value
func (s *DB) View(bucket, key string) ([]byte, error) {
	var v []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket not found")
		}
		v = b.Get([]byte(key))
		return nil
	})
	return v, err
}
