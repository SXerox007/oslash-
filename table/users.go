package table

import (
	"OSlash/constants"
	db "OSlash/db/postgres"
	"OSlash/protos/user/tweet"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AllTweets struct {
	TableName   struct{}  `sql:"all_tweets" json:"-"`
	Id          string    `param:"id"`
	Tweet       string    `param:"tweet"`
	UserId      string    `param:"user_id"`
	Version     string    `param:"version"`
	CaptureTime time.Time `param:"capture_time"`
}

// create tweet
func CreateTweet(req *tweetpb.TweetUserRequest) error {
	data := GetDetailsUsingAccessToken(req.GetCommonRequest().GetAccessToken())
	if data.AccessToken == "" {
		return status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintln(constants.ERR_MSG_INVALID_ACCESS_TOKEN, nil))
	}
	tweetData := AllTweets{
		Tweet:  req.GetTweet(),
		UserId: data.UserId,
	}
	err := db.GetPGClient().Create(&tweetData)
	return err
}

// list of tweets
func ListTweets(req *tweetpb.CommonRequest) ([]*tweetpb.TweetList, error) {
	data := GetDetailsUsingAccessToken(req.GetAccessToken())
	if data.AccessToken == "" {
		return nil, status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintln(constants.ERR_MSG_INVALID_ACCESS_TOKEN, nil))
	}
	tweetList := []AllTweets{}
	db.GetPGClient().Model(&tweetList).
		Where("user_id = ?", data.UserId).
		Select(&tweetList)

	list := []*tweetpb.TweetList{}
	for _, item := range tweetList {
		list = append(list, &tweetpb.TweetList{
			Tweet:   item.Tweet,
			TweetId: item.Id,
		})
	}

	return list, nil

}

// delete tweet
func DeleteTweet(req *tweetpb.DeleteTweetRequest) error {
	data := GetDetailsUsingAccessToken(req.GetCommonRequest().GetAccessToken())
	if data.AccessToken == "" {
		return status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintln(constants.ERR_MSG_INVALID_ACCESS_TOKEN, nil))
	}

	_, err := db.GetPGClient().Model(&AllTweets{}).
		Set("state = ?", 0).
		Where("id = ?", req.GetTweetId()).
		Update()
	return err
}
