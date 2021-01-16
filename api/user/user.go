package user

import (
	"OSlash/constants"
	"OSlash/protos/user/tweet"
	"OSlash/table"
	"context"
	"net/http"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterTweetService(srv *grpc.Server) {
	tweetpb.RegisterTweetServiceServer(srv, &Server{})
}

func (*Server) TweetCreate(ctx context.Context, req *tweetpb.TweetUserRequest) (*tweetpb.CommonResponse, error) {
	err := table.CreateTweet(req)
	if err == nil {
		return &tweetpb.CommonResponse{
			Message: constants.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil
	} else {
		return nil, err
	}
}

func (*Server) TweetList(ctx context.Context, req *tweetpb.CommonRequest) (*tweetpb.TweetListResponse, error) {
	list, err := table.ListTweets(req)
	if err == nil {
		return &tweetpb.TweetListResponse{
			TweetList: list,
		}, nil
	} else {
		return nil, err
	}
}

func (*Server) DeleteTweet(ctx context.Context, req *tweetpb.DeleteTweetRequest) (*tweetpb.CommonResponse, error) {
	err := table.DeleteTweet(req)
	if err == nil {
		return &tweetpb.CommonResponse{
			Message: constants.MSG_SUCCESS,
			Code:    http.StatusOK,
		}, nil
	} else {
		return nil, err
	}
}
