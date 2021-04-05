package phone

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"regexp"
)

type PhoneNumber struct {
	tableName struct{} `pg:"phone_numbers"`

	ID     int `pg:",pk"`
	Number string
}

func Main() {
	// Create the table
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		Database: "gophercises",
		User:     "postgres",
		Password: "qwe123",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		log.Fatal(err)
		return
	}

	phoneNumbers := &[]PhoneNumber{}
	err = db.Model(phoneNumbers).Select()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Normalize numbers
	for _, phoneNumber := range *phoneNumbers {
		reg, _ := regexp.Compile("[^0-9]+")
		normalized := string(reg.ReplaceAll([]byte(phoneNumber.Number), []byte("")))

		// Update numbers
		_, err = db.Model(&PhoneNumber{ID: phoneNumber.ID, Number: normalized}).WherePK().Update()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*PhoneNumber)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
