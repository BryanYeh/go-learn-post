package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connect() {
	var err error
	db, err = sql.Open("mysql", "root@/go_blog")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func getPostByID(id int64) (Post, error) {
	var post Post

	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	if err := row.Scan(&post.ID, &post.Title, &post.Author, &post.Content, &post.Created_At, &post.Updated_At); err != nil {
		if err == sql.ErrNoRows {
			return post, fmt.Errorf("postByID %d: no such post", id)
		}
		return post, fmt.Errorf("postByID %d: %v", id, err)
	}
	return post, nil
}

func getPosts() ([]Post, error) {
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, fmt.Errorf("getPosts: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var pos Post
		if err := rows.Scan(&pos.ID, &pos.Title, &pos.Author, &pos.Content, &pos.Created_At, &pos.Updated_At); err != nil {
			return nil, fmt.Errorf("getPosts: %v", err)
		}
		posts = append(posts, pos)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getPosts: %v", err)
	}
	return posts, nil
}

func getPostsByAuthor(name string) ([]Post, error) {
	// Posts slice to hold data from returned rows.
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts WHERE author = ?", name)
	if err != nil {
		return nil, fmt.Errorf("postsByAuthor %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var pos Post
		if err := rows.Scan(&pos.ID, &pos.Title, &pos.Author, &pos.Content, &pos.Created_At, &pos.Updated_At); err != nil {
			return nil, fmt.Errorf("postsByAuthor %q: %v", name, err)
		}
		posts = append(posts, pos)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postsByAuthor %q: %v", name, err)
	}
	return posts, nil
}

func addPost(post Post) (int64, error) {
	result, err := db.Exec("INSERT INTO posts (title, author, content) VALUES (?, ?, ?)", post.Title, post.Author, post.Content)
	if err != nil {
		return 0, fmt.Errorf("addPost: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addPost: %v", err)
	}
	return id, nil
}
