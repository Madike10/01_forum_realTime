package database

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateDatabase(database *sql.DB) {

	// Activer les contraintes FOREIGN KEY avec ON DELETE CASCADE et ON UPDATE CASCADE
	_, err := database.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}
	// Creation de la table user
	_, err = database.Exec(`
	CREATE TABLE IF NOT EXISTS users ( 
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			name_user TEXT,
			mail_user TEXT,
			password_user TEXT,
			age INTEGER,
			first_name TEXT,
			last_name TEXT,
			gender TEXT
			 ) 
			 `)
	if err != nil {
		fmt.Println("Users")
		log.Fatal(err)
	}

	// Creation de la table catégorie
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS category (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT, 
			name_category TEXT
		)
		`)

	if err != nil {
		fmt.Println("Categorie")
		log.Fatal()
	}

	var id_category string
	err = database.QueryRow("SELECT id FROM category WHERE name_category = 'other'").Scan(&id_category)
	if err != nil {
		_, err = database.Exec(`
			INSERT INTO category (name_category) VALUES ('technologie');
			INSERT INTO category (name_category) VALUES ('sport');
			INSERT INTO category (name_category) VALUES ('other');
			INSERT INTO category (name_category) VALUES ('sante');
			`)
	}

	// Inserer donnee de la table catégorie

	if err != nil {
		fmt.Println("Categorie")
		log.Fatal()
	}

	// Creation de la table users
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT, 
			title_post TEXT,
			content_post TEXT,
			date_post TEXT, 
			id_user INTEGER, 
			FOREIGN KEY("id_user") 
			REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
		 )
		`)
	if err != nil {
		fmt.Println("Post")
		log.Fatal(err)
	}

	// Création tavle belong
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS belong ( 
			id_post INTEGER, 
			id_category INTEGER, 
			PRIMARY KEY(id_post, id_category), 
			FOREIGN KEY(id_post) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE, 
			FOREIGN KEY(id_category) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE 
		)
		`)
	if err != nil {
		fmt.Println("Belong")
		log.Fatal(err)
	}

	// Création de la table likes-post

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS likes_post ( 
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT, 
			id_post INTEGER, 
			id_user INTEGER, 
			isLike INTEGER, 
			FOREIGN KEY(id_post) REFERENCES "posts"("id") ON DELETE CASCADE, FOREIGN KEY("id_user") REFERENCES "users"("id") ON DELETE CASCADE 
		)
		`)
	if err != nil {
		fmt.Println("Like_post")
		log.Fatal(err)
	}

	// Créate de la table commente
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS comment ( 
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT, 
			content_comment TEXT, 
			id_post INTEGER, 
			id_user INTEGER, 
			FOREIGN KEY(id_post) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE, 
			FOREIGN KEY(id_user) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE 
		)
		`)
	if err != nil {
		fmt.Println("Comment")
		log.Fatal(err)
	}

	// Creation de le table likes_comment
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS likes_comment (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			id_comment INTEGER,
			id_user INTEGER,
			isLike INTEGER, 
			FOREIGN KEY("id_comment") REFERENCES "comment"("id") ON DELETE CASCADE, FOREIGN KEY("id_user") REFERENCES "users"("id") ON DELETE CASCADE )
		`)
	if err != nil {
		fmt.Println("Like_comment")
		log.Fatal(err)
	}
	// Creation de la table session
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY,
    data TEXT,
    expiration DATETIME,
    id_user INTEGER,
    name_user TEXT
);
		`)
	if err != nil {
		fmt.Println("Session")
		log.Fatal(err)
	}
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS messages(
			id_message INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT, 
            id_sender INTEGER, 
			id_recever INTEGER,
            content_message TEXT,
			data_message TEXT,
			FOREIGN KEY(id_sender) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(id_recever) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		);
	`)
	if err != nil {
		fmt.Println("Message")
		log.Fatal(err)
	}
	fmt.Println("Data base create successfull")
}
