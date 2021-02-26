package post

import (
	"fmt"
	"testing"
)

func TestMakeOK(t *testing.T) {
	post, _ := Make("name", "content")
	if post.User != "name" {
		t.Errorf("Post user not set properly: %s", post.User)
	}
	if post.Content != "content" {
		t.Errorf("Post content not set properly: %s", post.Content)
	}

}

func TestMakeNotOK(t *testing.T) {
	_, err := Make("", "content")
	if err == nil {
		t.Errorf("Post user not set.  Should have returned error but didn't")
	}
	_, err = Make("user", "")
	if err == nil {
		t.Errorf("Post content not set.  Should have returned error but didn't")
	}

}

func TestAddNoInit(t *testing.T) {
	posts := New()
	// defer posts.Flush()
	lenBefore := len(*posts)
	// fmt.Println("len before", lenBefore)
	if lenBefore != 0 {
		t.Errorf("Number of posts did not eq 0")
	}

	post, _ := Make("name", "content")
	posts.Add(post)
	lenAfter := len(*posts)
	// fmt.Println("len after", lenAfter)
	if lenAfter != 1 {
		t.Errorf("Number of posts did not eq 1")
	}
}

func TestAdd(t *testing.T) {
	// defer posts.Flush()
	posts := New()

	post, _ := Make("aa", "aaaa")
	posts.Add(post)
	lenBefore := len(*posts)

	post, _ = Make("bb", "bbbb")
	posts.Add(post)
	lenAfter := len(*posts)

	if lenBefore+1 != lenAfter {
		t.Errorf("Number of posts did not increase by 1")
	}
}

func TestIsEmpty(t *testing.T) {
	// defer posts.Flush()
	posts := New()

	empty := posts.IsEmpty()
	if !empty {
		t.Errorf("Expected posts to be empty, but it's not (%v)", posts)
	}
	post, _ := Make("aa", "aaaa")
	posts.Add(post)
	empty = posts.IsEmpty()
	if empty {
		t.Errorf("Expected posts NOT to be empty, but it is (%v)", posts)
	}
}

func TestFlush(t *testing.T) {
	// defer posts.Flush()
	posts := New()

	addXPosts(posts, 5)
	empty := posts.IsEmpty()
	if empty {
		t.Errorf("Expected posts NOT to be empty, but it is (%v)", posts)
	}
	posts.Flush()
	empty = posts.IsEmpty()
	if !empty {
		t.Errorf("Expected posts to be empty, but it's not (%v)", posts)
	}
}

// helper functions
func addXPosts(p *Posts, num int) {
	for i := 0; i < num; i++ {
		post, _ := Make("name"+fmt.Sprint(i), "content"+fmt.Sprint(i))
		p.Add(post)
	}

	// return posts
}
