package models

import "sync"

type App struct {
	People []Person
	Mutex  sync.RWMutex
}
