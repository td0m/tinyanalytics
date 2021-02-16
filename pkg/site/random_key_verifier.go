package site

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RandomKeyVerifier struct {
	m     map[string]string
	mutex sync.RWMutex
}

func NewRandomKeyService() *RandomKeyVerifier {
	return &RandomKeyVerifier{
		m: map[string]string{},
	}
}

// errors
var (
	ErrKeyNotYetGenerated = errors.New("key not yet generated")
	ErrNotMatching        = errors.New("keys do not match")
)

func (v *RandomKeyVerifier) Generate(domain string) (key string) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	_, ok := v.m[domain]
	if ok {
		return
	}
	key = randStringBytes(32)
	v.m[domain] = key
	fmt.Println(v.m[domain])
	return
}

func (v *RandomKeyVerifier) Check(domain string, key string) (err error) {
	v.mutex.RLock()
	defer v.mutex.RUnlock()
	actualKey, ok := v.m[domain]
	if !ok {
		return ErrKeyNotYetGenerated
	}
	if actualKey != key {
		return ErrNotMatching
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
