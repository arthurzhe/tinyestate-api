package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/pragmaticivan/tinyestate-api/domain"
	stateRepo "github.com/pragmaticivan/tinyestate-api/state/repository"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestFeatch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Info(err)
		}
	}()

	mockStates := []domain.State{
		domain.State{
			ID:        "1asdgd7agds7",
			Name:      "California",
			Acronym:   "CA",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		domain.State{
			ID:        "2asdgd7agds7",
			Name:      "Texas",
			Acronym:   "TX",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		domain.State{
			ID:        "3asdgd7agds7",
			Name:      "Washington",
			Acronym:   "WA",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "acronym", "updated_at", "created_at"}).
		AddRow(mockStates[0].ID, mockStates[0].Name, mockStates[0].Acronym,
			mockStates[0].UpdatedAt, mockStates[0].CreatedAt).
		AddRow(mockStates[1].ID, mockStates[1].Name, mockStates[1].Acronym,
			mockStates[1].UpdatedAt, mockStates[1].CreatedAt).
		AddRow(mockStates[2].ID, mockStates[2].Name, mockStates[2].Acronym,
			mockStates[2].UpdatedAt, mockStates[2].CreatedAt)

	query := "SELECT id,name,acronym, updated_at, created_at FROM states WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := stateRepo.NewPostgresStateRepository(db)
	list, err := a.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, list, 3)
}
