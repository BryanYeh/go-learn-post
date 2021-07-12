package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connect()
	post, err := getPostByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Post found: %v\n", post)

	posts, err := getPosts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Posts found: %v\n", posts)

	postsAuthor, err := getPostsByAuthor("Betty Carter")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Posts by author found: %v\n", postsAuthor)
}
