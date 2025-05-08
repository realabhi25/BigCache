// Productionized BigCache Example
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/allegro/bigcache"
)

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

// Logging setup
func setupLogger() *log.Logger {
	file, err := os.OpenFile("bigcache.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open log file, using stderr")
		return log.New(os.Stderr, "LOG: ", log.LstdFlags)
	}
	return log.New(file, "LOG: ", log.LstdFlags)
}

var logger = setupLogger()

// User struct
type user struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

// BigCache struct
type bigCache struct {
	users *bigcache.BigCache
}

// Create new BigCache instance
func newBigCache() (*bigCache, error) {
	bCache, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         1 * time.Hour,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            false,
		HardMaxCacheSize:   256,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	})
	if err != nil {
		return nil, fmt.Errorf("cache initialization failed: %w", err)
	}
	logger.Println("BigCache initialized")
	return &bigCache{users: bCache}, nil
}

// Add or update user in cache
func (bc *bigCache) update(u user) error {
	bs, err := json.Marshal(&u)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	if err := bc.users.Set(userKey(u.Id), bs); err != nil {
		return fmt.Errorf("cache update failed: %w", err)
	}
	logger.Println("User updated in cache:", u.Id)
	return nil
}

// Generate user key
func userKey(id int64) string {
	return strconv.FormatInt(id, 10)
}

// Read user from cache
func (bc *bigCache) read(id int64) (user, error) {
	bs, err := bc.users.Get(userKey(id))
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return user{}, errUserNotInCache
		}
		return user{}, fmt.Errorf("cache read error: %w", err)
	}
	var u user
	if err := json.Unmarshal(bs, &u); err != nil {
		return user{}, fmt.Errorf("unmarshal error: %w", err)
	}
	logger.Println("User retrieved from cache:", u.Id)
	return u, nil
}

// Delete user from cache
func (bc *bigCache) delete(id int64) {
	bc.users.Delete(userKey(id))
	logger.Println("User deleted from cache:", id)
}

func main() {
	bc, err := newBigCache()
	if err != nil {
		logger.Fatalf("Failed to initialize cache: %v", err)
	}

	user1 := user{Id: 1, Email: "user1@example.com"}
	if err := bc.update(user1); err != nil {
		logger.Println("Error updating user:", err)
	}

	u, err := bc.read(1)
	if err != nil {
		logger.Println("Error reading user:", err)
	} else {
		logger.Println("User from cache:", u)
	}

	bc.delete(1)
}
