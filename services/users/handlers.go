package users

import (
	"context"
	"database/sql"

	"grpc-service-template-sqlc/db"
	"grpc-service-template-sqlc/pb"

	log "github.com/sirupsen/logrus"
)

type Handler struct {
	pb.UsersServer
	q *db.Queries
}

func New(q *db.Queries) *Handler {
	return &Handler{q: q}
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.WithFields(log.Fields{
		"name":      req.Name,
		"last_name": req.LastName,
		"email":     req.Age,
		"age":       req.Age,
	}).Debug("create new user req")

	userID, err := h.q.CreateUser(ctx, db.CreateUserParams{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Age:      req.Age,
	})
	if err != nil {
		log.WithError(err).Error("failed to create new user")
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: userID,
	}, nil
}

func (h *Handler) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	log.WithFields(log.Fields{
		"id": req.Id,
	}).Debug("get user by id req")

	user, err := h.q.GetUserByID(ctx, req.Id)
	if err != nil {
		log.WithError(err).Error("failed to get user by id")
		return nil, err
	}

	return &pb.GetUserByIdResponse{
		User: convertUserToPb(&user),
	}, nil
}

func (h *Handler) GetListOfUsersByIds(ctx context.Context, req *pb.GetListOfUsersByIdsRequest) (*pb.GetListOfUsersByIdsResponse, error) {
	log.WithFields(log.Fields{
		"id": req.Id,
	}).Debug("get users by ids req")

	users, err := h.q.GetListOfUsersByIDs(ctx, req.Id)
	if err != nil {
		log.WithError(err).Error("failed to get users by ids")
		return nil, err
	}

	return &pb.GetListOfUsersByIdsResponse{
		Users: convertUsersToPb(users),
	}, nil
}

func (h *Handler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	log.WithFields(log.Fields{
		"user_id": req.UserId,
		"title":   req.Title,
		"text":    req.Text,
	}).Debug("create new post req")

	postID, err := h.q.CreatePost(ctx, db.CreatePostParams{
		UserID: sql.NullInt32{Int32: int32(req.UserId), Valid: true},
		Title:  req.Title,
		Text:   req.Text,
	})
	if err != nil {
		log.WithError(err).Error("failed to create new post")
		return nil, err
	}

	table, err := h.q.JoinPostsAndUsersTables(ctx, postID)
	if err != nil {
		log.WithError(err).Error("failed to join posts and users tables")
		return nil, err
	}

	return &pb.CreatePostResponse{
		PostTable: convertPostTableToPb(&table),
	}, nil
}

func (h *Handler) GetPostsOfUser(ctx context.Context, req *pb.GetPostsOfUserRequest) (*pb.GetPostsOfUserResponse, error) {
	log.WithFields(log.Fields{
		"id": req.UserId,
	}).Debug("get posts of user req")

	posts, err := h.q.GetPostsOfUser(ctx, sql.NullInt32{Int32: int32(req.UserId), Valid: true})
	if err != nil {
		log.WithError(err).Error("failed to get posts of user")
		return nil, err
	}

	return &pb.GetPostsOfUserResponse{
		Posts: convertPostsToPb(posts),
	}, nil
}

func (h *Handler) CreateCommentForPost(ctx context.Context, req *pb.CreateCommentForPostRequest) (*pb.CreateCommentForPostResponse, error) {
	log.WithFields(log.Fields{
		"post_id": req.PostId,
		"user_id": req.UserId,
		"text":    req.Text,
	}).Debug("create comment for post req")

	postID, err := h.q.CreateCommentForPost(ctx, db.CreateCommentForPostParams{
		PostID: sql.NullInt32{Int32: int32(req.PostId), Valid: true},
		UserID: sql.NullInt32{Int32: int32(req.UserId), Valid: true},
		Text:   req.Text,
	})
	if err != nil {
		log.WithError(err).Error("failed create comment for post")
		return nil, err
	}

	post, err := h.q.JoinPostsAndUsersTables(ctx, postID.Int32)
	if err != nil {
		log.WithError(err).Error("failed join posts and users tables")
		return nil, err
	}

	comment, err := h.q.JoinCommentsAndUsersTables(ctx, postID)
	if err != nil {
		log.WithError(err).Error("failed join comments and users tables")
		return nil, err
	}

	return &pb.CreateCommentForPostResponse{
		PostTableWithComment: convertCommentPostTableToPb(&post, &comment),
	}, nil
}

func (h *Handler) DeleteCommentFromPost(ctx context.Context, req *pb.DeleteCommentFromPostRequest) (*pb.DeleteCommentFromPostResponse, error) {
	log.WithFields(log.Fields{
		"comment_id": req.CommentId,
	}).Debug("delete comment from post req")

	err := h.q.DeleteCommentFromPost(ctx, req.CommentId)
	if err != nil {
		log.WithError(err).Error("failed to delete comment from post")
		return nil, err
	}

	return &pb.DeleteCommentFromPostResponse{
		Message: "",
	}, nil
}

func (h *Handler) GetPostWithComments(ctx context.Context, req *pb.GetPostWithCommentsRequest) (*pb.GetPostWithCommentsResponse, error) {
	log.WithFields(log.Fields{
		"post_id": req.PostId,
	}).Debug("get post with comments req")

	post, err := h.q.GetPostByID(ctx, req.PostId)
	if err != nil {
		log.WithError(err).Error("failed get post by id")
		return nil, err
	}

	comments, err := h.q.GetCommentsOfPostByID(ctx, sql.NullInt32{Int32: int32(req.PostId), Valid: true})
	if err != nil {
		log.WithError(err).Error("failed get comments of post by id")
		return nil, err
	}

	return &pb.GetPostWithCommentsResponse{
		Post: convertPostToPb(&post, comments),
	}, nil
}
