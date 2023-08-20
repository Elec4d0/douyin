package main

import (
	"comment/comment_deploy/comment_mysql"
	api "comment/server/protos/kitex_gen/api"
	"comment/server/user_info"
	"context"
	"time"
)

// CommentServerImpl implements the last service interface defined in the IDL.
type CommentServerImpl struct{}

// CommentAction implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentAction(ctx context.Context, req *api.DouyinCommentActionRequest) (resp *api.DouyinCommentActionResponse, err error) {
	userId := req.UserId
	videoId := req.VideoId
	actionType := req.ActionType
	if actionType == 1 {
		//评论日期
		currentDateString := time.Now().Format("01-02")
		content := *req.CommentText

		if content == "" {
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}

		comment := &comment_mysql.Comment{
			Video_id:    videoId,
			Content:     content,
			Create_date: currentDateString,
			User_id:     userId,
		}
		err := comment_mysql.CreatComment(comment)
		if err != nil {
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}
		userInfo := user_info.UserInfo(userId)
		user := ToApiUser(userInfo)
		commentData := &api.Comment{
			Id:         videoId,
			Content:    content,
			CreateDate: currentDateString,
			User:       user,
		}
		statusMsg := "Comment successful"
		resp = &api.DouyinCommentActionResponse{
			StatusCode: int32(0),
			StatusMsg:  &statusMsg,
			Comment:    commentData,
		}
		return resp, nil
	} else if actionType == 2 {
		commentId := *req.CommentId
		comment, err := comment_mysql.FindComment(videoId, commentId)
		if err != nil {
			statusMsg := "Delete comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}
		statusMsg := "Delete comment successful"
		statusCode := 0
		resp = &api.DouyinCommentActionResponse{
			StatusCode: int32(statusCode),
			StatusMsg:  &statusMsg,
		}
		comment_mysql.DeleteComment(comment)
	}
	return
}

// CommentList implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentList(ctx context.Context, req *api.DouyinCommentListRequest) (resp *api.DouyinCommentListResponse, err error) {
	videoId := req.VideoId
	commentList, err := comment_mysql.FindCommentAll(videoId)
	if err != nil {
		statusMsg := "Get video comment list unsuccessful"
		resp = &api.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}
	if len(commentList) == 0 {
		statusMsg := "Get video comment list unsuccessful"
		resp = &api.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}
	//转换为api.comment
	var CommentList []*api.Comment
	for _, comment := range commentList {
		CommentList = append(CommentList, ToApiComment(comment))
	}

	statusMsg := "Get video comment list successful"
	statusCode := 0
	resp = &api.DouyinCommentListResponse{
		StatusCode:  int32(statusCode),
		StatusMsg:   &statusMsg,
		CommentList: CommentList,
	}
	return resp, nil
}

// CommentCount implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentCount(ctx context.Context, req *api.DouyinCommentserverCommentcountRequest) (resp *api.DouyinCommentserverCommentcountResponse, err error) {
	videoId := req.VideoId
	commentCount := comment_mysql.FindCommentCount(videoId)
	resp = &api.DouyinCommentserverCommentcountResponse{
		CommentCount: commentCount,
	}
	return resp, nil
}

// CommentAllCount implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentAllCount(ctx context.Context, req *api.DouyinCommentserverCommentallcountRequest) (resp *api.DouyinCommentserverCommentallcountResponse, err error) {
	videoIds := req.VideoIds
	commentCounts := comment_mysql.FindCommentAllCount(videoIds)
	resp = &api.DouyinCommentserverCommentallcountResponse{
		CommentCounts: commentCounts,
	}
	return resp, nil
}

func ToApiComment(comment *comment_mysql.Comment) *api.Comment {
	userInfo := user_info.UserInfo(comment.User_id)
	user := ToApiUser(userInfo)
	return &api.Comment{
		Id:         comment.Id,
		User:       user,
		Content:    comment.Content,
		CreateDate: comment.Create_date,
	}
}

func ToApiUser(user user_info.User) *api.User {
	return &api.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     &user.Follow_count,
		FollowerCount:   &user.Follower_count,
		IsFollow:        user.Is_follow,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.Background_image,
		Signature:       &user.Signature,
		TotalFavorited:  &user.Total_favorited,
		WorkCount:       &user.Work_count,
		FavoriteCount:   &user.Favorite_count,
	}
}
