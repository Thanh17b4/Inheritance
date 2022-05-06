package Inheri

import "fmt"

type Author struct {
	Firstname string
	Lastname  string
	Bio       string
}
type BlogPost struct {
	Title   string
	Content string
	Author
}

func (a Author) Fullname() string {
	return fmt.Sprintf(" %s %s", a.Firstname, a.Lastname)
}
func (b BlogPost) Details() {
	fmt.Println("Title: ", b.Title)
	fmt.Println("Content: ", b.Content)
	fmt.Println("Author: ", b.Fullname())
	fmt.Println("Bio: ", b.Bio)
}

type Website struct {
	BlogPosts []BlogPost
}

func (w Website) Contents() {
	fmt.Println(" Contents of Website\n ")
	for _, v := range w.BlogPosts {
		v.Details()
		fmt.Println()
	}
}
