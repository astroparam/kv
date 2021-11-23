package store

import "errors"

var ErrorNoSuchKey = errors.New("no such key")
var store = make(map[string]string)

// Put adds record to the store
func Put(key, value string) error {
	store[key] = value

	return nil
}

// Get returns value associated with key or
// error if key does not exists.
func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

// Delete deletes record associated with key
func Delete(key string) error {
	delete(store, key)

	return nil
}
