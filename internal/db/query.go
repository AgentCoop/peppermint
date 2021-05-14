package db

import (
	"sync"
)

type query struct {
	writeQuery sync.Mutex
}

func NewQuery() *query {

}