package models

import "time"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       string `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}

type Credential struct {
	Username_email string `json:"username_email"`
	Password       string `json:"password"`
}

type Post struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Date       string   `json:"date"`
	Categories []string `json:"categories"`
	Name_user  string   `json:"name_user_post"`
	Id_user    int      `json:"id_user"`
}

type SessionData struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	IsLogged    bool      `json:"is_logged"`
	CreatedTime time.Time `json:"created_time"`
}

type ResponsePost struct {
	Message  string      `json:"message"`
	Post     []Post      `json:"post"`
	UserData SessionData `json:"user_data"`
}

type ResponseUser struct {
	Id_user_id     int    `json:"id_user"`
	NameUser       string `json:"name_user"`
	Username_email string `json:"username_email"`
	Password       string `json:"password"`
	Message        string `json:"message"`
	OK             bool   `json:"ok"`
}
type LikeRequest struct {
	PostID        int  `json:"PostID"`
	UserID        int  `json:"UserID"`
	Like          bool `json:"Like"`
	IsLikeComment bool `json:"IsLikeComment"`
}

type LikeResponse struct {
	Status        bool   `json:"status"`
	Message       string `json:"message"`
	IsLike        bool   `json:"IsLike"`
	IsDislike     bool   `json:"IsDislike"`
	NbrLike       int    `json:"NbrLike"`
	NbrDislike    int    `json:"NbrDislike"`
	TypeLike      string `json:"TypeLike"`
	IsLikeComment bool   `json:"IsLikeComment"`
}
