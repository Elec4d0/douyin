package protos

import (
	"comment/comment_deploy/commentsql"
	"comment/sensitiveWord"
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
	log.Println("action")
	if actionType == 1 {
		//评论日期
		currentDateString := time.Now().Format("01-02")
		content := *req.CommentText
		content = sensitiveWord.ToInsensitive(content)

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
		log.Println(statusMsg)
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
	log.Println("list")
	videoId := req.VideoId
	userId := req.UserId
	commentList, err := commentsql.FindCommentAll(videoId)
	if err != nil {
		statusMsg := "获取评论列表失败"
		resp = &api.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}
	if len(commentList) == 0 {
		statusMsg := "此视频还没有评论哦，来留下珍贵的评论吧！"
		resp = &api.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}

	CommentList, err := user_info.UserInfoList(userId, commentList)
	if err != nil {
		statusMsg := "获取用户信息失败"
		resp = &api.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}

	statusMsg := ""
	var statusCode int32 = 0
	resp = &api.DouyinCommentListResponse{
		StatusCode:  statusCode,
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
