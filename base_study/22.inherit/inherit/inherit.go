package inherit

import (
	"fmt"
)

// 组合->也就嵌套结构体来代替继承
// 1.作者,父类,定义了子类需要嵌套的属性
type author struct {
	firstName string
	lastName  string
	bio       string
}

// Newauthor :创建作者对象
func Newauthor(firstName string, lastName string, bio string) author {
	author := author{firstName, lastName, bio}
	return author
}

func (a author) FullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// 2.博客,子类,除了父类的作者信息, 还包括内容和正文
type post struct {
	title   string
	content string
	author
}

// Newpost :创建博客对象
func Newpost(title string, content string, author author) post {
	post := post{title, content, author}
	return post
}

// Newposts :返回博客切片
func Newposts(xxx ...post) []post {
	posts := xxx
	return posts
}
func (p post) Details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.FullName())
	fmt.Println("Bio: ", p.author.bio)
}

// 3.网站, 博客切片体的嵌套
type website struct {
	posts []post
}

// Newwebsite :创建网站对象
func Newwebsite(posts []post) website {
	website := website{posts}
	return website
}
func (w website) Contents() {
	fmt.Println("Contents of Website: ")
	for _, v := range w.posts {
		fmt.Println("====================")
		v.Details()
	}
}
