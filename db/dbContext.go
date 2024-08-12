package db

import (
	"math/rand"
	"sync"
	"time"

	"quiz_me/db/entities"
)

type DBContext struct {
	questions map[int]entities.Question
	responses map[string]map[int]entities.Response
	results   map[string]entities.Result
	mu        sync.RWMutex
	rng       *rand.Rand
}

func NewDBContext() *DBContext {
	return &DBContext{
		questions: make(map[int]entities.Question),
		responses: make(map[string]map[int]entities.Response),
		results:   make(map[string]entities.Result),
		rng:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
