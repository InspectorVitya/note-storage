package memory

import (
	"context"
	"github.com/inspectorvitya/note-storage/internal/model"
	"github.com/inspectorvitya/note-storage/internal/storage"
	"sync"
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
	mu        sync.RWMutex
}

func New() storage.NoteStorage {
	return &list{}
}

func (l *list) AddNote(_ context.Context, note model.Note) (model.IDNote, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	note.IDNote = l.lastIndex
	item := itemList{value: note}
	if l.head == nil {
		l.head = &item
		l.tail = &item
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = &item
		item.prev = last
		l.tail = &item
	}
	l.size++
	l.lastIndex++
	return note.IDNote, nil
}
func (l *list) DeleteNote(_ context.Context, id model.IDNote) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	note, err := l.findById(id)
	if err != nil {
		return err
	}
	if l.size == 1 {
		l.tail = nil
		l.head = nil
		l.size--
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
	l.size--

	return nil
}
func (l *list) GetAll(_ context.Context) ([]model.Note, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.size == 0 {
		return nil, model.ErrEmptyList
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
		return model.Note{}, model.ErrEmptyList
	}
	return l.head.value, nil
}
func (l *list) GetLast(_ context.Context) (model.Note, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.size == 0 {
		return model.Note{}, model.ErrEmptyList
	}
	return l.tail.value, nil
}

func (l *list) findById(id model.IDNote) (*itemList, error) {
	item := l.head
	for {
		if item != nil && item.value.IDNote == id {
			return item, nil
		}
		if item.next == nil {
			return nil, model.ErrNotExistNote
		}
		item = item.next
	}
}
