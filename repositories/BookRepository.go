package repositories

import (
	"challenge-3/database"
	"challenge-3/models"
	"database/sql"
)

var (
	db  *sql.DB
	err error
)

func AddBook(book models.Book) error {
	db, _ := database.Connect(db, err)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `
	INSERT INTO books (title, description, author)
	VALUES ($1, $2, $3)
	Returning *
	`

	_, err := db.Exec(sqlStatement, book.Title, book.Description, book.Author)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(book models.Book, id int) error {
	db, _ := database.Connect(db, err)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `
	UPDATE books 
	SET title = $2, author = $3, description = $4
	WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id, book.Title, book.Description, book.Author)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id int) error {
	db, _ := database.Connect(db, err)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `
	DELETE FROM books 
	WHERE id = $1;
	`

	_, err := db.Exec(sqlStatement, id)

	if err != nil {
		return err
	}

	return nil
}

func GetBook(id int) (models.Book, error) {
	db, _ := database.Connect(db, err)
	if err != nil {
		return models.Book{}, err
	}
	defer db.Close()

	sqlStatement := `
	SELECT * FROM books 
	WHERE id = $1;
	`

	var book models.Book
	err = db.QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.Description, &book.Author)

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func GetAllBooks() ([]models.Book, error) {
	db, _ := database.Connect(db, err)
	if err != nil {
		return []models.Book{}, err
	}
	defer db.Close()

	sqlStatement := `
	SELECT * FROM books;
	`

	var books []models.Book
	rows, err := db.Query(sqlStatement)

	if err != nil {
		return []models.Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author)

		if err != nil {
			return []models.Book{}, err
		}

		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}
