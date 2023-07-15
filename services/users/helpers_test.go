package users

import (
	"database/sql"
	"testing"
	"time"

	"grpc-service-template-sqlc/db"
	"grpc-service-template-sqlc/pb"

	"github.com/stretchr/testify/assert"
)

func Test_convertUserToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		user := db.User{
			ID:        100,
			Name:      "Ivan",
			LastName:  "Ivanov",
			Email:     "ivan@fletn.com",
			Age:       100,
			CreatedAt: time.Now(),
		}

		want := &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
			Age:      user.Age,
		}

		got := convertUserToPb(&user)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertUsersToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		users := make([]db.User, 2)
		for i, _ := range users {
			users[i] = db.User{
				ID:        int32(i + 1),
				Name:      "Ivan",
				LastName:  "Ivanov",
				Email:     "ivan@fletn.com",
				Age:       int32((i + 1)) * 10,
				CreatedAt: time.Now(),
			}
		}

		pbUsers := make([]*pb.User, 2)
		for i, u := range users {
			pbUsers[i] = &pb.User{
				Id:       u.ID,
				Name:     u.Name,
				LastName: u.LastName,
				Email:    u.Email,
				Age:      u.Age,
			}
		}

		want := pbUsers
		got := convertUsersToPb(users)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostsToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		posts := make([]db.Post, 2)
		for i, _ := range posts {
			posts[i] = db.Post{
				ID:     int32(i + 1),
				UserID: sql.NullInt32{Int32: int32(2), Valid: true},
				Title:  "Holidays",
				Text:   "Meet friends",
			}
		}

		pbPosts := make([]*pb.Post, 2)
		for i, p := range posts {
			pbPosts[i] = &pb.Post{
				Id:     p.ID,
				UserId: p.UserID.Int32,
				Title:  p.Title,
				Text:   p.Text,
			}
		}

		want := pbPosts
		got := convertPostsToPb(posts)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostTableToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p := &db.JoinPostsAndUsersTablesRow{
			ID:       int32(1),
			Title:    "Weekends",
			Text:     "watch film",
			Name:     "Gosha",
			LastName: "Krasilnikov",
		}

		c := &db.JoinCommentsAndUsersTablesRow{
			ID:       int32(2),
			Text:     "how are you",
			Name:     "Ilya",
			LastName: "Krasnov",
		}

		pbPostTable := &pb.PostTableWithComment{
			Id:                  p.ID,
			UserName:            p.Name,
			UserLastName:        p.LastName,
			Title:               p.Title,
			Text:                p.Text,
			CommentId:           c.ID,
			Comment:             c.Text,
			CommentUserName:     c.Name,
			CommentUserLastName: c.LastName,
		}

		want := pbPostTable
		got := convertCommentPostTableToPb(p, c)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostTable(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p := &db.JoinPostsAndUsersTablesRow{
			ID:       int32(1),
			Title:    "Weekends",
			Text:     "watch film",
			Name:     "Gosha",
			LastName: "Krasilnikov",
		}

		pbPostTable := &pb.PostTable{
			Id:           p.ID,
			UserName:     p.Name,
			UserLastName: p.LastName,
			Title:        p.Title,
			Text:         p.Text,
		}

		want := pbPostTable
		got := convertPostTableToPb(p)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertCommentsToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comments := make([]db.Comment, 2)
		for i, _ := range comments {
			comments[i] = db.Comment{
				ID:     int32(i + 1),
				PostID: sql.NullInt32{Int32: int32(i * 2), Valid: true},
				UserID: sql.NullInt32{Int32: int32(2), Valid: true},
				Text:   "hello",
			}
		}

		pbComments := make([]*pb.Comment, 2)
		for i, c := range comments {
			pbComments[i] = &pb.Comment{
				Id:     c.ID,
				PostId: c.PostID.Int32,
				UserId: c.UserID.Int32,
				Text:   c.Text,
			}
		}

		want := pbComments
		got := convertCommentsToPb(comments)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := &db.Post{
			ID:     int32(1),
			UserID: sql.NullInt32{Int32: int32(2), Valid: true},
			Title:  "holidays",
			Text:   "meet with freinds",
		}

		comments := make([]db.Comment, 2)
		for i, _ := range comments {
			comments[i] = db.Comment{
				ID:     int32(i + 1),
				PostID: sql.NullInt32{Int32: int32(1), Valid: true},
				UserID: sql.NullInt32{Int32: int32(2), Valid: true},
				Text:   "hello",
			}
		}

		pbPost := &pb.Post{
			Id:       post.ID,
			UserId:   post.UserID.Int32,
			Title:    post.Title,
			Text:     post.Text,
			Comments: convertCommentsToPb(comments),
		}

		want := pbPost
		got := convertPostToPb(post, comments)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}
