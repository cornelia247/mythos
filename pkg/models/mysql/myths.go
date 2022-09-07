package mysql

import (
	"database/sql"

	"github.com/cornelia247/mythos/pkg/models"
)

// Define a MythModel type which wraps a sql.DB connection pool
type MythModel struct {
	DB *sql.DB
}

// // This will insert a new myth into the database.
func (m *MythModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO myths (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the statement. The first parameter is the SQL statement, followed by the  title, content and expiry values for the placeholder parameters. This method returns a sql.Result object, which contains some basic information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result object to get the ID of our // newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// The ID returned has the type int64, so we convert it to an int type // before returning.
	return int(id), nil
}

// This will return a specific myth based on its id.
func (m *MythModel) Get(id int) (*models.Myth, error) {
	return nil, nil
}

// // This will return the 10 most recently created myths.
func (m *MythModel) Latest() ([]*models.Myth, error) {
	return nil, nil
}
