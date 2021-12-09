package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Posts []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comments []struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/practice5")
	if err != nil {
		log.Fatal(err)
	}

	var posts Posts
	var comments Comments

	post := getPosts(posts)
	for i := 0; i < 10; i++ {
		writePosts(db, post, i)
		comment := getComments(post, comments, post[i].ID)
		for j := 0; j < 5; j++ {
			writeComments(db, comment, j)
		}
	}

}

func getPosts(posts Posts) Posts {
	i := 3
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=" + strconv.Itoa(i))
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalln(err)
	}

	return posts
}

func getComments(post Posts, comments Comments, i int) Comments {

	resp2, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(i))
	if err != nil {
		log.Fatal(err)
	}

	body2, _ := ioutil.ReadAll(resp2.Body)
	err = json.Unmarshal(body2, &comments)

	return comments
}

func writePosts(db *sql.DB, posts Posts, i int) {

	insert, err := db.Query("INSERT INTO `sql.public.posts` (`user_id`, `post_id`, `post_tittle`, `post_body`) VALUES (?, ?, ?, ?)", posts[i].UserID, posts[i].ID, posts[i].Body, posts[i].Title)
	if err != nil {
		log.Fatalln(err)
	}
	defer insert.Close()
}

func writeComments(db *sql.DB, comments Comments, i int) {

	insert, err := db.Query("INSERT INTO  `sql.public.comments`(`post_id`, `comments_id`, `comments_name`, `email`, `comments_body`) VALUES (?, ?, ?, ?, ?)", comments[i].PostID, comments[i].ID, comments[i].Name, comments[i].Email, comments[i].Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer insert.Close()
}
