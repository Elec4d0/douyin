package protos

import (
	"comment/comment_deploy/commentsql"
	api "comment/server/protos/kitex_gen/api"
	"comment/server/user_info"
	"context"
	"log"
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

		comment := &commentsql.Comment{
			Video_id:    videoId,
			Content:     content,
			Create_date: currentDateString,
			User_id:     userId,
		}
		err := commentsql.CreatComment(comment)
		if err != nil {
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}
		commentsql.CommentCountAdd(videoId)

		userInfo, err := user_info.UserInfo(userId, userId)
		if err != nil {
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}

		commentData := &api.Comment{
			Id:         videoId,
			Content:    content,
			CreateDate: currentDateString,
			User:       userInfo,
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
		comment, err := commentsql.FindComment(videoId, commentId)
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
		commentsql.DeleteComment(comment)
		commentsql.CommentCountDel(videoId)
	}
	return
}

// CommentList implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentList(ctx context.Context, req *api.DouyinCommentListRequest) (resp *api.DouyinCommentListResponse, err error) {
	videoId := req.VideoId
	userId := req.UserId
	commentList, err := commentsql.FindCommentAll(videoId)
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

	CommentList, err := user_info.UserInfoList(userId, commentList)
	if err != nil {
		statusMsg := "Get video comment list unsuccessful"
		resp = &api.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
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
	commentCount, err := commentsql.FindCommentCount(videoId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	resp = &api.DouyinCommentserverCommentcountResponse{
		CommentCount: commentCount,
	}
	return resp, nil
}

// CommentAllCount implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentAllCount(ctx context.Context, req *api.DouyinCommentserverCommentallcountRequest) (resp *api.DouyinCommentserverCommentallcountResponse, err error) {
	videoIds := req.VideoIds
	commentCounts, err := commentsql.FindCommentAllCount(videoIds)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	resp = &api.DouyinCommentserverCommentallcountResponse{
		CommentCounts: commentCounts,
	}
	return resp, nil
}
