package application

import (
	"context"
	"github.com/inspectorvitya/note-storage/internal/model"
	"github.com/inspectorvitya/note-storage/internal/storage"
	"time"
)

type App struct {
	noteStorage storage.NoteStorage
}

func New(noteStorage storage.NoteStorage) *App {
	return &App{noteStorage}
}

func (app *App) GetAllNotes(ctx context.Context) ([]model.Note, error) {
	notes, err := app.noteStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return notes, err
}

func (app *App) GetLastNote(ctx context.Context) (model.Note, error) {
	note, err := app.noteStorage.GetLast(ctx)
	if err != nil {
		return model.Note{}, err
	}
	return note, err
}

func (app *App) GetFirstNote(ctx context.Context) (model.Note, error) {
	note, err := app.noteStorage.GetFirst(ctx)
	if err != nil {
		return model.Note{}, err
	}
	return note, err
}

func (app *App) DeleteNote(ctx context.Context, id model.IDNote) error {
	return app.noteStorage.DeleteNote(ctx, id)
}

func (app *App) CreateNote(ctx context.Context, note model.Note) error {
	id, err := app.noteStorage.AddNote(ctx, note)
	if err != nil {
		return err
	}

	if note.ExpireTime != "" {
		expire, err := time.ParseDuration(note.ExpireTime)
		if err != nil {
			return err
		}
		time.AfterFunc(expire, func() {
			_ = app.DeleteNote(ctx, id)
		})
	}
	return err
}
