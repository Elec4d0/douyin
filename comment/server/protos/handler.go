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
	if actionType == 1 {
		currentDateString := time.Now().Format("01-02") //评论日期
		content := *req.CommentText
		content = sensitiveWord.ToInsensitive(content) //敏感词处理

		if content == "" { //评论为空
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: -1,
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
		err := commentsql.CreatComment(comment) //数据库存评论
		if err != nil {
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: -1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}
		commentsql.CommentCountAdd(videoId) //计数库+1

		userInfo, err := user_info.UserInfo(userId, userId) //获取用户信息
		if err != nil {
			statusMsg := "Comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: -1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}

		commentData := &api.Comment{ //类型转换
			Id:         videoId,
			Content:    content,
			CreateDate: currentDateString,
			User:       userInfo,
		}

		statusMsg := "Comment successful"
		log.Println(statusMsg)
		resp = &api.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  &statusMsg,
			Comment:    commentData,
		}
		return resp, nil
	} else if actionType == 2 {
		commentId := *req.CommentId
		comment, err := commentsql.FindComment(videoId, commentId) //先查是否存在
		if err != nil {
			statusMsg := "Delete comment unsuccessful"
			resp = &api.DouyinCommentActionResponse{
				StatusCode: -1,
				StatusMsg:  &statusMsg,
			}
			return resp, nil
		}
		statusMsg := "Delete comment successful"
		resp = &api.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  &statusMsg,
		}
		commentsql.DeleteComment(comment)   //删除评论
		commentsql.CommentCountDel(videoId) //计数库-1
	}
	return
}

// CommentList implements the CommentServerImpl interface.
func (s *CommentServerImpl) CommentList(ctx context.Context, req *api.DouyinCommentListRequest) (resp *api.DouyinCommentListResponse, err error) {
	videoId := req.VideoId
	userId := req.UserId
	commentList, err := commentsql.FindCommentAll(videoId) //查所有评论
	if err != nil {
		statusMsg := "Get video comment list unsuccessful"
		resp = &api.DouyinCommentListResponse{
			StatusCode: -1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}
	if len(commentList) == 0 { //特判，没有会出错 即查评论表长度是否为0 为0不存在
		statusMsg := "Get video comment list unsuccessful"
		resp = &api.DouyinCommentListResponse{
			StatusCode: -1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}

	CommentList, err := user_info.UserInfoList(userId, commentList) //获取所有评论者信息
	if err != nil {
		statusMsg := "Get video comment list unsuccessful"
		resp = &api.DouyinCommentListResponse{
			StatusCode: -1,
			StatusMsg:  &statusMsg,
		}
		return resp, nil
	}

	statusMsg := "Get video comment list successful"
	resp = &api.DouyinCommentListResponse{
		StatusCode:  0,
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
