package post

import (
	"errors"
	"fmt"
	"sync"
)

type Post struct {
	User    string
	Content string
}
type Posts []Post

var (
	pmu sync.Mutex
)

func Make(user string, content string) (Post, error) {
	if len(user) <= 0 {
		return Post{}, errors.New("User can't be empty")
	} else if len(content) <= 0 {
		return Post{}, errors.New("Content can't be empty")
	}

	return Post{User: user, Content: content}, nil
}

func New() *Posts {
	p := Posts(make([]Post, 0))
	fmt.Println("Created new post", p)
	return &p
}

func (p *Posts) Add(post Post) {
	fmt.Println("Add new post", post)

	pmu.Lock()
	defer pmu.Unlock()
	tmp := append([]Post(*p), post)
	*p = Posts(tmp)
}

func (p *Posts) IsEmpty() bool {
	fmt.Println("Is posts empty?", (len(*p) <= 0))
	return len(*p) <= 0
}

func (p *Posts) Flush() {
	fmt.Println("Flushing posts")
	pmu.Lock()
	defer pmu.Unlock()
	*p = (*p)[:0]
}
