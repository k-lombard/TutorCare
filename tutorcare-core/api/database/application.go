package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllApplications() (*models.ApplicationList, error) {
	list := &models.ApplicationList{}
	rows, err := db.Conn.Query("SELECT * FROM applications ORDER BY application_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var app models.Application
		err := rows.Scan(&app.UserID, &app.PostID, &app.ApplicationID, &app.Message, &app.Accepted, &app.DateCreated)
		if err != nil {
			return list, err
		}
		list.Applications = append(list.Applications, app)
	}
	return list, nil
}

// func (db Database) GetActivePosts() (*models.PostList, error) {
// 	list := &models.PostList{}
// 	rows, err := db.Conn.Query("SELECT * FROM posts ORDER BY post_id DESC")
// 	if err != nil {
// 		return list, err
// 	}
// 	for rows.Next() {
// 		var post models.Post
// 		err5 := rows.Scan(&post.UserID, &post.CaregiverID, &post.PostID, &post.CareDescription, &post.CareType, &post.Completed, &post.DateOfJob, &post.StartTime, &post.EndTime, &post.DatePosted)
// 		if err5 != nil {
// 			return list, err5
// 		}
// 		// var us models.User
// 		// err2 := db.Conn.QueryRow(`SELECT * FROM users WHERE user_id = $1;`, post.UserID).Scan(&us.UserID, &us.FirstName, &us.LastName, &us.Email, &us.Password, &us.DateJoined, &us.Status, &us.UserCategory, &us.Experience, &us.Bio)
// 		// if err2 != nil {
// 		// 	return list, err2
// 		// }
// 		if post.Completed == false && (post.CaregiverID).String() == "00000000-0000-0000-0000-000000000000" {
// 			// post.User = us
// 			list.Posts = append(list.Posts, post)
// 		}
// 	}
// 	return list, nil
// }

func (db Database) AddApplication(app *models.Application) (models.Application, error) {
	var poster uuid.UUID
	appOut := models.Application{}
	errLatest := db.Conn.QueryRow(`SELECT user_id FROM posts where post_id=$1`, &app.PostID).Scan(&poster)
	if errLatest != nil {
		return appOut, errLatest
	}

	var userType string
	errFinal := db.Conn.QueryRow(`SELECT user_category from users where user_id=$1`, app.UserID).Scan(&userType)
	if errFinal != nil {
		return appOut, errFinal
	}

	if poster != app.UserID && userType != "careseeker" {
		sqlStatement := `INSERT INTO applications (user_id, post_id, message) VALUES ($1, $2, $3) RETURNING application_id, accepted, date_created;`
		var application_id int
		var accepted bool
		var date_created string

		err := db.Conn.QueryRow(sqlStatement, &app.UserID, &app.PostID, &app.Message).Scan(&application_id, &accepted, &date_created)

		if err != nil {
			return appOut, err
		}
		err2 := db.Conn.QueryRow(`SELECT * FROM applications WHERE application_id = $1;`, application_id).Scan(&appOut.UserID, &appOut.PostID, &appOut.ApplicationID, &appOut.Message, &appOut.Accepted, &appOut.DateCreated)
		if err2 != nil {
			return appOut, err2
		}
		fmt.Println("New post record created with applicationID and timestamp: ", application_id, date_created)
		return appOut, nil
	} else {
		return appOut, ErrSameUser
	}
}

func (db Database) GetApplicationById(appId int) (models.ApplicationWithUser, error) {
	appOut := models.ApplicationWithUser{}
	query := `SELECT * FROM applications WHERE application_id = $1;`
	row := db.Conn.QueryRow(query, appId)
	switch err := row.Scan(&appOut.UserID, &appOut.PostID, &appOut.ApplicationID, &appOut.Message, &appOut.Accepted, &appOut.DateCreated); err {
	case sql.ErrNoRows:
		return appOut, ErrNoMatch
	default:
		userOut := models.User{}
		query2 := `SELECT * FROM users WHERE user_id = $1;`
		row2 := db.Conn.QueryRow(query2, appOut.UserID)
		err3 := row2.Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
		if err3 != nil {
			return appOut, err3
		}
		appOut.User = userOut
		return appOut, nil
	}
}

func (db Database) GetApplicationsByPostId(postId int) (*models.ApplicationList, error) {
	list := &models.ApplicationList{}
	rows, err := db.Conn.Query("SELECT * FROM applications WHERE post_id = $1 ORDER BY application_id DESC", postId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var appOut models.Application
		err := rows.Scan(&appOut.UserID, &appOut.PostID, &appOut.ApplicationID, &appOut.Message, &appOut.Accepted, &appOut.DateCreated)
		if err != nil {
			return list, err
		}
		list.Applications = append(list.Applications, appOut)
	}
	return list, nil
}

func (db Database) GetApplicationsByUserId(userId uuid.UUID) (*models.ApplicationList, error) {
	list := &models.ApplicationList{}
	rows, err := db.Conn.Query("SELECT * FROM applications WHERE user_id = $1 ORDER BY application_id DESC", userId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var appOut models.Application
		err := rows.Scan(&appOut.UserID, &appOut.PostID, &appOut.ApplicationID, &appOut.Message, &appOut.Accepted, &appOut.DateCreated)
		if err != nil {
			return list, err
		}
		list.Applications = append(list.Applications, appOut)
	}
	return list, nil
}

func (db Database) DeleteApplication(applicationId int) error {
	var id int
	query := `DELETE FROM applications WHERE application_id = $1 RETURNING application_id;`
	err := db.Conn.QueryRow(query, applicationId).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Application deleted with ApplicationID: ", id)
	return nil
}

func (db Database) UpdateApplication(applicationId int, appData models.Application) (models.Application, error) {
	app := models.Application{}
	query := `UPDATE applications SET message=$1, accepted=$2 WHERE application_id=$3 RETURNING user_id, post_id, application_id, message, accepted, date_created;`

	query2 := `SELECT * FROM applications WHERE application_id=$1;`
	app2 := models.Application{}
	errTwo := db.Conn.QueryRow(query2, applicationId).Scan(&app2.UserID, &app2.PostID, &app2.ApplicationID, &app2.Message, &app2.Accepted, &app2.DateCreated)
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return app, ErrNoMatch
		}
		return app, errTwo
	}
	err := db.Conn.QueryRow(query, appData.Message, appData.Accepted, applicationId).Scan(&app.UserID, &app.PostID, &app.ApplicationID, &app.Message, &app.Accepted, &app.DateCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			return app, ErrNoMatch
		}
		return app, err
	}
	return app, nil
}
