package contacts

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

//Contact class
type Contact struct {
	ID        uuid.UUID `json:"id" validate:"min=8"`
	FirstName string    `json:"first_name" validate:"nonzero" creating:"nonzero"`
	LastName  string    `json:"last_name" validate:"nonzero" creating:"nonzero"`
	Email     string    `json:"email" validate:"min=8" creating:"min=8"`
	Phone     string    `json:"phone" validate:"min=8" creating:"len=12,regexp=^([0-9]{3}-){2}[0-9]{4}$"`
}

func (c *Contact) getContact(db *sql.DB) error {
	return db.QueryRow(
		"SELECT first_name, last_name, email, phone FROM contacts WHERE id=$1",
		c.ID,
	).Scan(&c.FirstName, &c.LastName, &c.Email, &c.Phone)
}

func (c *Contact) updateContact(db *sql.DB) error {

	_, err := db.Exec(
		"UPDATE contacts SET first_name=$1, last_name=$2, email=$3, phone=$4 WHERE id=$5",
		c.FirstName, c.LastName, c.Email, c.Phone, c.ID)

	return err
}
func (c *Contact) updateCSVContact(db *sql.DB) error {
	// "INSERT INTO contact () VAlUES () on DUPLICATE KEY UPDATE name, age"
	_, err := db.Exec(
		"UPDATE contacts SET first_name=$1, last_name=$2 email=$3 phone=$4 WHERE id=$5; INSERT INTO contacts (first_name, last_name, email, phone) SELECT $1, $2, $3, $4 WHERE NOT EXISTS (SELECT 1 FROM table WHERE id=$5);")
	return err
}

func (c *Contact) deleteContact(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM contacts WHERE id=$1", c.ID)

	return err
}

func (c *Contact) getContacts(db *sql.DB, start, count int) ([]Contact, error) {
	rows, err := db.Query(
		"SELECT id, first_name, last_name, email, phone FROM contacts LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Contacts := []Contact{}

	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone); err != nil {
			return nil, err
		}
		Contacts = append(Contacts, c)
	}

	return Contacts, nil
}

func (c *Contact) createContact(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO contacts(first_name, last_name, email, phone, id) VALUES($1, $2, $3, $4, $5) RETURNING id",
		c.FirstName, c.LastName, c.Email, c.Phone, c.ID).Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}
