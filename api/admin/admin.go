package admin

import (
	"OSlash/protos/admin"
	"context"

	"google.golang.org/grpc"
)

type Server struct {
}

func RegisterAdminService(srv *grpc.Server) {
	adminpb.RegisterAdminServiceServer(srv, &Server{})
}

func (*Server) EditUserDetails(ctx context.Context, req *adminpb.EditUserDetailsRequest) (*adminpb.CommonResponse, error) {
	return nil, nil
}

func (*Server) AdminCreateTweet(ctx context.Context, req *adminpb.TweetAdminRequest) (*adminpb.CommonResponse, error) {
	return nil, nil
}

func (*Server) AdminDeleteTweet(ctx context.Context, req *adminpb.DeleteAdminTweetRequest) (*adminpb.CommonResponse, error) {
	return nil, nil
}

func (*Server) AdminUpdateTweet(ctx context.Context, req *adminpb.AdminTweetUpdateRequest) (*adminpb.CommonResponse, error) {
	return nil, nil
}

func (*Server) AdminListTweet(ctx context.Context, req *adminpb.AdminTweetListRequest) (*adminpb.TweetListResponse, error) {
	return nil, nil
}
