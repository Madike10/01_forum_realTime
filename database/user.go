package database

import (
	"database/sql"
	"fmt"
	"realtimeforum/models"
)

func AddUser(url string, u models.User) error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
	defer db.Close()
	req, err := db.Prepare("INSERT INTO users(name_user, mail_user, password_user, age, first_name, last_name, gender) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = req.Exec(u.Name, u.Email, u.Password, u.Age, u.FirstName, u.LastName, u.Gender)
	if err != nil {
		return err
	}
	return nil
}

// ADD POST IN THE DATABASE

func AddPost(url string, p models.Post) error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
	defer db.Close()
	tabCategory := p.Categories

	techno := tabCategory[0]
	sport := tabCategory[1]
	other := tabCategory[2]
	sante := tabCategory[3]

	reqCategory, err := db.Prepare("SELECT id FROM category WHERE name_category = ? OR name_category = ? OR name_category = ? OR name_category = ?")
	if err != nil {
		return err
	}
	defer reqCategory.Close()

	rows, err := reqCategory.Query(techno, sport, other, sante)
	if err != nil {
		return err
	}
	defer rows.Close()

	// mettre les ID categories dans un tab category
	tabID_category := []int{}

	for rows.Next() {
		var id_category int
		err = rows.Scan(&id_category)
		if err != nil {
			return err
		}
		tabID_category = append(tabID_category, id_category)
	}
	fmt.Println(tabID_category)
	if len(tabID_category) == 0 {
		return fmt.Errorf("no category found")
	}

	req, err := db.Prepare("INSERT INTO posts(title_post, content_post, date_post, id_user) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = req.Exec(p.Title, p.Content, p.Date, p.Id_user)
	if err != nil {
		return err
	}
	id_post, err := db.Prepare("SELECT id FROM posts WHERE title_post = ? AND date_post = ?")
	if err != nil {
		return err
	}
	id, err := id_post.Query(p.Title, p.Date)
	var id_Post int
	for id.Next() {
		if err := id.Scan(&id_Post); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	defer id.Close()

	for i := 0; i < len(tabID_category); i++ {
		req_b, err := db.Prepare("INSERT INTO belong(id_post, id_category) VALUES(?,?)")
		if err != nil {
			return err
		}
		_, err = req_b.Exec(id_Post, tabID_category[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// GET ALL POST FROM THE BD

func GetAllPost(url string) ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var posts []models.Post
	rows, err := db.Query("SELECT posts.id, posts.title_post, posts.content_post, posts.date_post, posts.id_user, users.name_user FROM posts JOIN users ON posts.id_user = users.id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.Id, &p.Title, &p.Content, &p.Date, &p.Id_user, &p.Name_user)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

//GET USER BY POST ID

func GetUserByPostID(db *sql.DB, postID int) (*models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM users WHERE id = (SELECT id_user FROM posts WHERE id = ?)", postID)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Age, &user.FirstName, &user.LastName, &user.Gender)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GET ID_CATEGORIE BY ID_POST

func GetIdCategoryByIDPost(db *sql.DB, id int) []int {
	// db, err := sql.Open("sqlite3", "./database.db")
	// if err!= nil {
	//     return nil
	// }
	// defer db.Close()
	var tabID_category []int
	req, err := db.Prepare("SELECT id_category FROM belong WHERE id_post =?")
	if err != nil {
		return nil
	}
	rows, err := req.Query(id)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var id_category int
		err = rows.Scan(&id_category)
		if err != nil {
			return nil
		}
		tabID_category = append(tabID_category, id_category)
	}
	return tabID_category
}
