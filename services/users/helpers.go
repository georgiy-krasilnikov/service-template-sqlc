package users

import (
	"grpc-service-template-sqlc/db"
	"grpc-service-template-sqlc/pb"
)

func convertUserToPb(u *db.User) *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{
		Id:       u.ID,
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
		Age:      u.Age,
	}
}

func convertUsersToPb(u []db.User) []*pb.User {
	users := make([]*pb.User, len(u))
	for i, user := range u {
		users[i] = &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
			Age:      user.Age,
		}
	}

	return users
}

func convertPostsToPb(p []db.Post) []*pb.Post {
	posts := make([]*pb.Post, len(p))
	for i, post := range p {
		posts[i] = &pb.Post{
			Id:     post.ID,
			UserId: post.UserID.Int32,
			Title:  post.Title,
			Text:   post.Text,
		}
	}

	return posts
}

func convertPostTableToPb(p *db.JoinPostsAndUsersTablesRow) *pb.PostTable {
	if p == nil {
		return nil
	}

	return &pb.PostTable{
		Id:           p.ID,
		UserName:     p.Name,
		UserLastName: p.LastName,
		Title:        p.Title,
		Text:         p.Text,
	}
}

func convertCommentPostTableToPb(p *db.JoinPostsAndUsersTablesRow, c *db.JoinCommentsAndUsersTablesRow) *pb.PostTableWithComment {
	if p == nil {
		return nil
	}

	return &pb.PostTableWithComment{
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
}

func convertCommentsToPb(c []db.Comment) []*pb.Comment {
	comments := make([]*pb.Comment, len(c))
	for i, comment := range c {
		comments[i] = &pb.Comment{
			Id:     comment.ID,
			PostId: comment.PostID.Int32,
			UserId: comment.UserID.Int32,
			Text:   comment.Text,
		}
	}

	return comments
}

func convertPostToPb(p *db.Post, c []db.Comment) *pb.Post {
	if p == nil {
		return nil
	}

	return &pb.Post{
		Id:       p.ID,
		UserId:   p.UserID.Int32,
		Title:    p.Title,
		Text:     p.Text,
		Comments: convertCommentsToPb(c),
	}
}
