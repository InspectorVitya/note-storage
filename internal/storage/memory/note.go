package memory

import (
	"context"
	"errors"
	"github.com/inspectorvitya/note-storage/internal/model"
	"github.com/inspectorvitya/note-storage/internal/storage"
	"sync"
	"time"
)

type itemList struct {
	value model.Note
	next  *itemList
	prev  *itemList
}

type list struct {
	head      *itemList
	tail      *itemList
	size      int
	lastIndex model.IDNote
	items     map[model.IDNote]*itemList
	mu        sync.RWMutex
}

func New() storage.NoteStorage {
	return &list{
		items: make(map[model.IDNote]*itemList),
	}
}

func (l *list) AddNote(_ context.Context, note model.Note) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	note.IDNote = l.lastIndex
	note.CreateTime = time.Now()
	item := itemList{value: note}
	if l.head == nil {
		l.head = &item
		l.tail = &item
		l.items[note.IDNote] = &item
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = &item
		item.prev = last
		l.tail = &item
		l.items[note.IDNote] = &item
	}
	l.size++
	l.lastIndex++
	return nil
}
func (l *list) DeleteNote(_ context.Context, id model.IDNote) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	note, ok := l.items[id]
	if !ok {
		return errors.New("not exist note")
	}
	if l.size == 1 {
		l.tail = nil
		l.head = nil
		l.size--
		delete(l.items, id)
		return nil
	}
	if note.prev != nil {
		note.prev.next = note.next
	} else {
		note.next.prev = nil
		l.head = note.next
	}
	if next := note.next; next != nil {
		next.prev = note.prev
	} else {
		note.prev.next = nil
		l.tail = note.prev
	}
	delete(l.items, id)
	l.size--

	return nil
}
func (l *list) GetAll(_ context.Context) ([]model.Note, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.size == 0 {
		return nil, errors.New("empty list")
	}
	result := make([]model.Note, l.size)
	note := l.tail
	for i := 0; ; i++ {
		if note == nil {
			break
		}
		result[i] = note.value
		note = note.prev
	}
	return result, nil
}
func (l *list) GetFirst(_ context.Context) (model.Note, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.size == 0 {
		return model.Note{}, errors.New("empty list")
	}
	return l.head.value, nil
}
func (l *list) GetLast(_ context.Context) (model.Note, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.size == 0 {
		return model.Note{}, errors.New("empty list")
	}
	return l.tail.value, nil
}
