package contacts

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

//Contact class
type Contact struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

func (c *Contact) getContact(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM contacts WHERE id=$1",
		c.ID).Scan(&c.FirstName, &c.LastName, &c.Email, &c.Phone)
}

func (c *Contact) updateContact(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE contacts SET first_name=$1, last_name=$2, email=$3, price=$4 WHERE id=$5",
			c.FirstName, c.LastName, c.Email, c.Phone, c.ID)

	return err
}

func (c *Contact) deleteContact(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM contacts WHERE id=$1", c.ID)

	return err
}

func (c *Contact) getContacts(db *sql.DB, start, count int) ([]Contact, error) {
	rows, err := db.Query(
		"SELECT id, first_name, last_name, email, phone FROM Contacts LIMIT $1 OFFSET $2",
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
