package db

import (
	"sync"

	"quiz_me/db/entities"
)

type DBContext struct {
	questions map[int]entities.Question
	responses map[string]map[int]entities.Response
	results   map[string]entities.Result
	mu        sync.RWMutex
}

func NewDBContext() *DBContext {
	return &DBContext{
		questions: make(map[int]entities.Question),
		responses: make(map[string]map[int]entities.Response),
		results:   make(map[string]entities.Result),
	}
}
