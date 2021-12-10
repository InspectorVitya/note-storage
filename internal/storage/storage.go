package storage

import (
	"context"
	"github.com/inspectorvitya/note-storage/internal/model"
)

type NoteStorage interface {
	AddNote(ctx context.Context, note model.Note) (model.IDNote, error)
	DeleteNote(ctx context.Context, id model.IDNote) error
	GetAll(ctx context.Context) ([]model.Note, error)
	GetFirst(ctx context.Context) (model.Note, error)
	GetLast(ctx context.Context) (model.Note, error)
}
