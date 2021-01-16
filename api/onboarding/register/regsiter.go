package register

import (
	"OSlash/constants"
	db "OSlash/db/postgres"
	"OSlash/protos/onboarding/register"
	"OSlash/table"
	"OSlash/utils/pswdmanager"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

func RegisterUserService(srv *grpc.Server) {
	regsiterpb.RegisterRegisterServiceServer(srv, &Server{})
}

func (*Server) RegisterUserService(ctx context.Context, req *regsiterpb.RegisterUserRequest) (*regsiterpb.RegisterUserResponse, error) {
	var err error
	var token string
	token, err = CreateUser(req)
	if err == nil {
		//success
		return &regsiterpb.RegisterUserResponse{
			Message:     constants.MSG_SUCCESS,
			Code:        http.StatusOK,
			AccessToken: token,
		}, nil
	}

	// return &regsiterpb.RegisterUserResponse{
	// 	Message:     appconstant.MSG_FAILURE,
	// 	Code:        http.StatusNoContent,
	// 	AccessToken: "",
	// }, err
	return nil, err
}

type Users struct {
	TableName   struct{}  `sql:"all_users" json:"-"`
	Id          string    `param:"id"`
	Phone       string    `param:"phone"`
	FirstName   string    `param:"first_name"`
	LastName    string    `param:"last_name"`
	Role        string    `param:"role"`
	Password    string    `param:"password"`
	Email       string    `param:"email"`
	Version     string    `param:"version"`
	CaptureTime time.Time `param:"capture_time"`
}

// create user
func CreateUser(req *regsiterpb.RegisterUserRequest) (accessToken string, err error) {
	if req.GetEmail() == "" || req.GetPassword() == "" || req.GetRole() == "" {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(constants.ERR_MSG_INTERNAL_SERVER, err))
	}
	data := Users{
		Email:    req.GetEmail(),
		Password: pswdmanager.GetPswdHash([]byte(req.GetPassword())),
		Role:     req.GetRole(),
	}

	err = db.GetPGClient().Create(&data)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(constants.ERR_MSG_INTERNAL_SERVER, err))
	}
	log.Println("User Data:", data)
	return table.CreateUserAccessToken(data.Id, req.GetEmail())
}
