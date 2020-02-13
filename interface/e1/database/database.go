package database

// DB is the interface to database functions
type DB interface {
	// Open the database
	Open(database string) error
	// Close database
	Close() error
	// Update the value of the key in the bucket
	Update(bucket, key string, payload []byte) error
	// View the value of the key and bucket
	View(bucket, key string) ([]byte, error)
}
