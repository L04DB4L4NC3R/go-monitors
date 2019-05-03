package main

import (
	"fmt"
	"sync"
)

type Monitor interface {
	Wait()
	Signal()
	GetData() []string
	PutData(string)
}

type Words struct {
	mutex         *sync.Mutex
	wordsArray    []string
	isInitialized bool
}

func (m *Words) Init() {
	m.mutex = &sync.Mutex{}
	m.wordsArray = []string{}
	m.isInitialized = true
}

func (m *Words) Wait() {
	if m.isInitialized {
		m.mutex.Lock()
	}
}

func (m *Words) Signal() {
	if m.isInitialized {
		m.mutex.Unlock()
	}
}

func (m *Words) GetData() []string { return m.wordsArray }

func (m *Words) PutData(word string) {
	m.Wait()

	// critical section
	m.wordsArray = append(m.wordsArray, word)
	// critical section done

	m.Signal()
}

func main() {
	m := &Words{}
	m.Init()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		m.PutData("Angad")
	}()
	go func() {
		defer wg.Done()
		m.PutData("Sharma")
	}()
	wg.Wait()
	fmt.Println(m.GetData())
}
