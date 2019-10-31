package mysql

import (
	"database/sql"

	"github.com/emwp/go-studies/snippetbox/pkg/models"
)

type SnippedModel struct {
	DB *sql.DB
}

func (m *SnippedModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippedModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippedModel) Latest(id int) ([]*models.Snippet, error) {
	return nil, nil
}
