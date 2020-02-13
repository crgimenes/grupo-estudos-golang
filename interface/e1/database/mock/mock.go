package mock

// DB holds the database connection instance
type DB struct {
}

// New instantiate the storage DB and open the database file
func New(database string) (*DB, error) {
	s := &DB{}
	return s, nil
}

// Open the database
func (s *DB) Open(database string) error {
	return nil
}

// Close database
func (s *DB) Close() error {
	return nil
}

// Update the value of the key in the bucket
func (s *DB) Update(bucket, key string, payload []byte) error {
	return nil
}

// View get bucket and key value
func (s *DB) View(bucket, key string) ([]byte, error) {
	return []byte("test"), nil
}
