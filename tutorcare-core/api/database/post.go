package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllPosts() (*models.PostList, error) {
	list := &models.PostList{}
	rows, err := db.Conn.Query("SELECT * FROM posts ORDER BY post_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
		if err != nil {
			return list, err
		}
		list.Posts = append(list.Posts, post)
	}
	return list, nil
}

func (db Database) GetActivePosts() (*models.PostList, error) {
	list := &models.PostList{}
	rows, err := db.Conn.Query("SELECT * FROM posts ORDER BY post_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err5 := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
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

func (db Database) AddPost(post *models.Post) (models.Post, error) {
	sqlStatement := `INSERT INTO posts (user_id, care_description, care_type, date_of_job, start_time, end_time) VALUES ($1, $2, $3, $4, $5, $6) RETURNING caregiver_id, post_id, completed, date_posted;`
	var caregiver_id uuid.UUID
	var post_id int
	var completed bool
	var date_posted string
	postOut := models.Post{}

	err := db.Conn.QueryRow(sqlStatement, &post.UserID, &post.CareDescription, &post.CareType, &post.DateOfJob, &post.StartTime, &post.EndTime).Scan(&caregiver_id, &post_id, &completed, &date_posted)

	if err != nil {
		return postOut, err
	}
	err2 := db.Conn.QueryRow(`SELECT * FROM posts WHERE post_id = $1;`, post_id).Scan(&postOut.UserID, &postOut.CaregiverID, &postOut.PostID, &postOut.CareDescription, &postOut.CareType, &postOut.Completed, &postOut.DateOfJob, &postOut.StartTime, &postOut.EndTime, &postOut.DatePosted)
	if err2 != nil {
		return postOut, err2
	}
	fmt.Println("New post record created with postID and timestamp: ", post_id, date_posted)
	return postOut, nil
}

func (db Database) GetPostById(postId int) (models.Post, error) {
	postOut := models.Post{}
	query := `SELECT * FROM posts WHERE post_id = $1;`
	row := db.Conn.QueryRow(query, postId)
	switch err := row.Scan(&postOut.UserID, &postOut.CaregiverID, &postOut.PostID, &postOut.CareDescription, &postOut.CareType, &postOut.Completed, &postOut.DateOfJob, &postOut.StartTime, &postOut.EndTime, &postOut.DatePosted); err {
	case sql.ErrNoRows:
		return postOut, ErrNoMatch
	default:
		return postOut, err
	}
}

func (db Database) GetPostsByUserId(userId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	rows, err := db.Conn.Query("SELECT * FROM posts WHERE user_id = $1 ORDER BY post_id DESC", userId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
		if err != nil {
			return list, err
		}
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
	query := `UPDATE posts SET caregiver_id=$1, care_description=$2, care_type=$3, completed=$4, date_of_job=$5, start_time=$6, end_time=$7 WHERE post_id=$8 RETURNING user_id, caregiver_id, post_id, care_description, care_type, completed, date_of_job, start_time, end_time, date_posted;`

	query2 := `SELECT * FROM geolocation WHERE user_id = $1;`
	post2 := models.Post{}
	errTwo := db.Conn.QueryRow(query2, postId).Scan(&post2.UserID, &post2.CaregiverID, &post2.PostID, &post2.CareDescription, &post2.CareType, &post2.Completed, &post2.DateOfJob, &post2.StartTime, &post2.EndTime, &post2.DatePosted)
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	err := db.Conn.QueryRow(query, postData.CaregiverID, postData.CareDescription, postData.CareType, postData.Completed, &postData.DateOfJob, &postData.StartTime, &postData.EndTime, postId).Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, err
	}
	return post, nil
}
