package main

import (
	"fmt"
	"net/http"

	"github.com/pbreedt/simwebapp/post"
)

var posts *post.Posts

func main() {
	posts = post.New()
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/new", newHandler)

	http.ListenAndServe(":8080", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html>"))
	defer func() {
		w.Write([]byte("<a href='/new'>New post</a>"))
		w.Write([]byte("</html>"))
	}()

	w.Write([]byte("<h2>Posts</h2>"))

	if posts.IsEmpty() {
		w.Write([]byte("<p>No posts</p>"))
		return
	}
	w.Write([]byte("<ul>"))
	for _, post := range *posts {
		w.Write([]byte(fmt.Sprintf("<li>%s: %s</li>", post.User, post.Content)))
	}
	w.Write([]byte("</ul>"))
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><form action=\"/add\" method=\"POST\">"))
	defer func() {
		w.Write([]byte("</form></html>"))
	}()

	w.Write([]byte("<div><label>Name:</label><input type='text' name='user' /></div>"))
	w.Write([]byte("<div><label>Comment:</label><input type='text' name='comment' /></div>"))
	w.Write([]byte("<div><input type='submit'>Submit</input></div>"))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	usr := r.FormValue("user")
	fmt.Println("Got user from POST:", usr)
	cmt := r.FormValue("comment")
	fmt.Println("Got content from POST:", cmt)
	pst, err := post.Make(usr, cmt)
	if err != nil {
		w.Write([]byte("<div>Error occurred saving post</div>"))
		w.Write([]byte("<div>" + err.Error() + "</div>"))
	} else {
		posts.Add(pst)
		listHandler(w, r)
	}
}
