package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllPosts() (*models.PostList, error) {
	list := &models.PostList{}
	rows, err := db.Conn.Query("SELECT *, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy') as NewDateFormat FROM posts ORDER BY post_id DESC;")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.Title, &post.Tags, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted, &post.DateOfJob)
		if err != nil {
			return list, err
		}
		list.Posts = append(list.Posts, post)
	}
	return list, nil
}

func (db Database) GetActivePosts() (*models.PostList, error) {
	list := &models.PostList{}
	rows, err := db.Conn.Query("SELECT *, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy') FROM posts ORDER BY post_id DESC;")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err5 := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.Title, &post.Tags, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted, &post.DateOfJob)
		if err5 != nil {
			return list, err5
		}
		// var us models.User
		// err2 := db.Conn.QueryRow(`SELECT * FROM users WHERE user_id = $1;`, post.UserID).Scan(&us.UserID, &us.FirstName, &us.LastName, &us.Email, &us.Password, &us.DateJoined, &us.Status, &us.UserCategory, &us.Experience, &us.Bio)
		// if err2 != nil {
		// 	return list, err2
		// }
		if post.Completed == false && (post.CaregiverID).String() == "00000000-0000-0000-0000-000000000000" {
			// post.User = us
			list.Posts = append(list.Posts, post)
		}
	}
	return list, nil
}

func (db Database) GetActivePostsWithCaregiver(userId uuid.UUID) (*models.PostWithCaregiverList, error) {
	list := &models.PostWithCaregiverList{}
	rows, err := db.Conn.Query(`SELECT *, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy') FROM posts WHERE user_id=$1 ORDER BY post_id DESC;`, userId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.PostWithCaregiver
		err5 := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.Title, &post.Tags, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted, &post.DateOfJob)
		if err5 != nil {
			return list, err5
		}
		if post.Completed == false && (post.CaregiverID).String() != "00000000-0000-0000-0000-000000000000" {
			careUser := models.User{}
			row5 := db.Conn.QueryRow(`SELECT * FROM users WHERE user_id=$1;`, post.CaregiverID)
			errNew := row5.Scan(&careUser.UserID, &careUser.FirstName, &careUser.LastName, &careUser.Email, &careUser.Password, &careUser.DateJoined, &careUser.Status, &careUser.UserCategory, &careUser.Experience, &careUser.Bio)
			if errNew != nil {
				return list, errNew
			}

			row6 := db.Conn.QueryRow(`SELECT application_id FROM applications WHERE user_id = $1 AND post_id = $2;`, post.CaregiverID, post.PostID)
			errFinal := row6.Scan(&post.ApplicationID)
			if errFinal != nil {
				return list, errFinal
			}
			post.Caregiver = careUser
			list.Posts = append(list.Posts, post)
		}
	}
	return list, nil
}

func (db Database) AddPost(post *models.Post) (models.Post, error) {
	sqlStatement := `INSERT INTO posts (user_id, title, tags, care_description, care_type, date_of_job, start_time, end_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING caregiver_id, post_id, completed, date_posted;`
	var caregiver_id uuid.UUID
	var post_id int
	var completed bool
	var date_posted string
	postOut := models.Post{}

	err := db.Conn.QueryRow(sqlStatement, &post.UserID, &post.Title, &post.Tags, &post.CareDescription, &post.CareType, &post.DateOfJob, &post.StartTime, &post.EndTime).Scan(&caregiver_id, &post_id, &completed, &date_posted)

	if err != nil {
		return postOut, err
	}
	err2 := db.Conn.QueryRow(`SELECT *, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy') FROM posts WHERE post_id = $1;`, post_id).Scan(&postOut.UserID, &postOut.CaregiverID, &postOut.PostID, &postOut.Title, &post.Tags, &postOut.CareDescription, &postOut.CareType, &postOut.Completed, &postOut.DateOfJob, &postOut.StartTime, &postOut.EndTime, &postOut.DatePosted, &postOut.DateOfJob)
	if err2 != nil {
		return postOut, err2
	}
	fmt.Println("New post record created with postID and timestamp: ", post_id, date_posted)
	return postOut, nil
}

func (db Database) GetPostById(postId int) (models.Post, error) {
	postOut := models.Post{}
	query := `SELECT *, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy') FROM posts WHERE post_id = $1;`
	row := db.Conn.QueryRow(query, postId)
	switch err := row.Scan(&postOut.UserID, &postOut.CaregiverID, &postOut.PostID, &postOut.Title, &postOut.Tags, &postOut.CareDescription, &postOut.CareType, &postOut.Completed, &postOut.DateOfJob, &postOut.StartTime, &postOut.EndTime, &postOut.DatePosted, &postOut.DateOfJob); err {
	case sql.ErrNoRows:
		return postOut, ErrNoMatch
	default:
		return postOut, err
	}
}

func (db Database) GetPostsByUserId(userId uuid.UUID) (*models.PostWithApplicationsList, error) {
	list := &models.PostWithApplicationsList{}
	rows, err := db.Conn.Query("SELECT *, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy') FROM posts WHERE user_id = $1 ORDER BY post_id DESC;", userId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.PostWithApplications
		err := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.Title, &post.Tags, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted, &post.DateOfJob)
		if err != nil {
			return list, err
		}
		rows2, err2 := db.Conn.Query("SELECT * FROM applications WHERE post_id = $1 ORDER BY application_id DESC;", &post.PostID)
		if err2 != nil {
			return list, err2
		}
		list2 := &models.ApplicationWithUserList{}
		for rows2.Next() {
			var app models.ApplicationWithUser
			err3 := rows2.Scan(&app.UserID, &app.PostID, &app.ApplicationID, &app.Message, &app.Accepted, &app.DateCreated)
			if err3 != nil {
				return list, err3
			}
			var newUser models.User
			query := `SELECT * FROM users WHERE user_id = $1;`
			err4 := db.Conn.QueryRow(query, app.UserID).Scan(&newUser.UserID, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.Password, &newUser.DateJoined, &newUser.Status, &newUser.UserCategory, &newUser.Experience, &newUser.Bio)
			if err4 != nil {
				return list, err4
			}
			app.User = newUser
			list2.Applications = append(list2.Applications, app)
		}
		post.Applications = list2.Applications
		list.Posts = append(list.Posts, post)
	}
	return list, nil
}

func (db Database) DeletePost(postId int) error {
	var id int
	query := `DELETE FROM posts WHERE post_id = $1 RETURNING post_id;`
	err := db.Conn.QueryRow(query, postId).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Post deleted with PostID: ", id)
	return nil
}

func (db Database) UpdatePost(postId int, postData models.Post) (models.Post, error) {
	post := models.Post{}
	query := `UPDATE posts SET caregiver_id=$1, title=$2, tags=$3, care_description=$4, care_type=$5, completed=$6, date_of_job=$7, start_time=$8, end_time=$9 WHERE post_id=$10 RETURNING user_id, caregiver_id, post_id, title, tags, care_description, care_type, completed, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy'), start_time, end_time, date_posted;`

	query2 := `SELECT * FROM posts WHERE post_id = $1;`
	post2 := models.Post{}
	errTwo := db.Conn.QueryRow(query2, postId).Scan(&post2.UserID, &post2.CaregiverID, &post2.PostID, &post2.Title, &post2.Tags, &post2.CareDescription, &post2.CareType, &post2.Completed, &post2.DateOfJob, &post2.StartTime, &post2.EndTime, &post2.DatePosted)
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	err := db.Conn.QueryRow(query, postData.CaregiverID, postData.Title, &postData.Tags, postData.CareDescription, postData.CareType, postData.Completed, &postData.DateOfJob, &postData.StartTime, &postData.EndTime, postId).Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, err
	}
	return post, nil
}

func (db Database) AddApplicationToPost(postId int, postData models.Post) (models.Post, error) {
	post := models.Post{}
	query := `UPDATE posts SET caregiver_id=$1 WHERE post_id=$2 RETURNING user_id, caregiver_id, post_id, title, tags, care_description, care_type, completed, TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy'), start_time, end_time, date_posted;`

	query2 := `SELECT * FROM posts WHERE post_id=$1;`
	post2 := models.Post{}
	errTwo := db.Conn.QueryRow(query2, postId).Scan(&post2.UserID, &post2.CaregiverID, &post2.PostID, &post2.Title, &post2.Tags, &post2.CareDescription, &post2.CareType, &post2.Completed, &post2.DateOfJob, &post2.StartTime, &post2.EndTime, &post2.DatePosted)
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	err := db.Conn.QueryRow(query, postData.CaregiverID, postId).Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post2.Title, &post2.Tags, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, err
	}
	return post, nil
}
