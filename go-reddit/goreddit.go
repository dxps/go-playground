package goreddit

import "github.com/google/uuid"

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID            uuid.UUID `db:"id"`
	ThreadID      uuid.UUID `db:"thread_id"`
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	Votes         int       `db:"votes"`
	CommentsCount int       `db:"comments_count"`
	ThreadTitle   string    `db:"thread_title"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
}

type ThreadStore interface {
	GetThread(id uuid.UUID) (Thread, error)
	GetThreads() ([]Thread, error)
	SaveThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	GetPosts() ([]Post, error)
	GetPostsByThread(threadID uuid.UUID) ([]Post, error)
	GetPost(id uuid.UUID) (Post, error)
	SavePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	GetComment(id uuid.UUID) (Comment, error)
	GetCommentsByPost(postID uuid.UUID) ([]Comment, error)
	SaveComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id uuid.UUID) error
}

type UserStore interface {
	GetUser(id uuid.UUID) (User, error)
	GetUserByUsername(username string) (User, error)
	SaveUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(id uuid.UUID) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
	UserStore
}
