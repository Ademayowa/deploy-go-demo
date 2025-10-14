package models

import (
	"github.com/google/uuid"

	"github.com/Ademayowa/deploy-go-demo/db"
)

type Property struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Amount      float64 `json:"amount"`
}

// Save property into the database
func (property *Property) Save() error {
	property.ID = uuid.New().String()

	query := `
		INSERT INTO properties(id, title, description, location, amount)
		VALUES(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		property.ID,
		property.Title,
		property.Description,
		property.Location,
		property.Amount,
	)
	return err
}

// Get all properties
func GetAllProperties() ([]Property, error) {
	query := "SELECT * FROM properties"

	// Fetch properties
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []Property

	for rows.Next() {
		var property Property
		err := rows.Scan(
			&property.ID,
			&property.Title,
			&property.Description,
			&property.Location,
			&property.Amount,
		)
		if err != nil {
			return nil, err
		}

		properties = append(properties, property)
	}

	return properties, nil
}

// Get property by ID
func GetPropertyByID(id string) (Property, error) {
	var property Property

	query := "SELECT * FROM properties WHERE id =?"
	row := db.DB.QueryRow(query, id)
	err := row.Scan(
		&property.ID,
		&property.Title,
		&property.Description,
		&property.Location,
		&property.Amount,
	)
	if err != nil {
		return property, err
	}

	return property, nil
}

// Delete property
func (property Property) Delete() error {
	query := "DELETE FROM properties WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(property.ID)
	return err
}

// Update property
func UpdateProperty(id string, updatedProperty Property) error {
	query := `
		UPDATE properties
		SET title = ?, description = ?, location = ?, amount = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query,
		updatedProperty.Title,
		updatedProperty.Description,
		updatedProperty.Location,
		updatedProperty.Amount,
		id,
	)
	return err
}
