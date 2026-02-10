package main

import (
	"log"
	"os"
	blogposts "github.com/hrutvikyadav/blogposts"

)

func main() {
	posts, err := blogposts.PostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
