package application

import (
	"context"
	"github.com/inspectorvitya/note-storage/internal/model"
	"github.com/inspectorvitya/note-storage/internal/storage/memory"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	exp1 := model.Note{
		IDNote: 0,
		Title:  "test",
		Text:   "test",
	}
	exp2 := model.Note{
		IDNote: 1,
		Title:  "test",
		Text:   "test",
	}
	storage := memory.New()
	app := New(storage)
	ctx := context.Background()
	t.Run("test create notes", func(t *testing.T) {
		note := model.Note{
			Title: "test",
			Text:  "test",
		}
		err := app.CreateNote(ctx, note)
		require.NoError(t, err)
		err = app.CreateNote(ctx, note)
		require.NoError(t, err)
	})
	t.Run("test get all notes", func(t *testing.T) {
		notes, err := app.GetAllNotes(ctx)
		require.Equal(t, exp1, notes[1])
		require.Equal(t, exp2, notes[0])
		require.NoError(t, err)
	})
	t.Run("test get first note", func(t *testing.T) {
		note, err := app.GetFirstNote(ctx)
		require.NoError(t, err)
		require.Equal(t, exp1, note)
	})
	t.Run("test get last note", func(t *testing.T) {
		note, err := app.GetLastNote(ctx)
		require.NoError(t, err)
		require.Equal(t, exp2, note)
	})
	t.Run("test delete note", func(t *testing.T) {
		err := app.DeleteNote(ctx, 0)
		require.NoError(t, err)
		notes, err := app.GetAllNotes(ctx)
		require.NoError(t, err)
		require.Equal(t, 1, len(notes))
	})
	t.Run("test delete on expiration of time", func(t *testing.T) {
		err := app.CreateNote(ctx, model.Note{
			IDNote:     0,
			Title:      "time",
			Text:       "time",
			ExpireTime: "2s",
		})
		require.NoError(t, err)
		time.Sleep(time.Second * 3)
		notes, err := app.GetAllNotes(ctx)
		require.NoError(t, err)
		require.Equal(t, 1, len(notes))
	})

}
