// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DouyinCommentActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentActionRequest[number], err)
}

func (x *DouyinCommentActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CommentText = &tmp
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.CommentId = &tmp
	return offset, err
}

func (x *DouyinCommentActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentActionResponse[number], err)
}

func (x *DouyinCommentActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinCommentActionResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comment = &v
	return offset, nil
}

func (x *Comment) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Comment[number], err)
}

func (x *Comment) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *Comment) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CreateDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinCommentListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentListRequest[number], err)
}

func (x *DouyinCommentListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentListResponse[number], err)
}

func (x *DouyinCommentListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinCommentListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommentList = append(x.CommentList, &v)
	return offset, nil
}

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FollowCount = &tmp
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FollowerCount = &tmp
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *User) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Avatar = &tmp
	return offset, err
}

func (x *User) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.BackgroundImage = &tmp
	return offset, err
}

func (x *User) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Signature = &tmp
	return offset, err
}

func (x *User) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.TotalFavorited = &tmp
	return offset, err
}

func (x *User) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.WorkCount = &tmp
	return offset, err
}

func (x *User) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FavoriteCount = &tmp
	return offset, err
}

func (x *DouyinCommentserverCommentcountRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentserverCommentcountRequest[number], err)
}

func (x *DouyinCommentserverCommentcountRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentserverCommentcountResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentserverCommentcountResponse[number], err)
}

func (x *DouyinCommentserverCommentcountResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.CommentCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentserverCommentallcountRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentserverCommentallcountRequest[number], err)
}

func (x *DouyinCommentserverCommentallcountRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.VideoIds = append(x.VideoIds, v)
			return offset, err
		})
	return offset, err
}

func (x *DouyinCommentserverCommentallcountResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentserverCommentallcountResponse[number], err)
}

func (x *DouyinCommentserverCommentallcountResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.CommentCounts = append(x.CommentCounts, v)
			return offset, err
		})
	return offset, err
}

func (x *DouyinCommentActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVideoId())
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetActionType())
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField4(buf []byte) (offset int) {
	if x.CommentText == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCommentText())
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField5(buf []byte) (offset int) {
	if x.CommentId == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetCommentId())
	return offset
}

func (x *DouyinCommentActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinCommentActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DouyinCommentActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DouyinCommentActionResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Comment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetComment())
	return offset
}

func (x *Comment) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Comment) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Comment) fastWriteField2(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetUser())
	return offset
}

func (x *Comment) fastWriteField3(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetContent())
	return offset
}

func (x *Comment) fastWriteField4(buf []byte) (offset int) {
	if x.CreateDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCreateDate())
	return offset
}

func (x *DouyinCommentListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinCommentListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *DouyinCommentListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVideoId())
	return offset
}

func (x *DouyinCommentListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinCommentListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DouyinCommentListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DouyinCommentListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.CommentList == nil {
		return offset
	}
	for i := range x.GetCommentList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetCommentList()[i])
	}
	return offset
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFollowCount())
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetFollowerCount())
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.GetIsFollow())
	return offset
}

func (x *User) fastWriteField6(buf []byte) (offset int) {
	if x.Avatar == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetAvatar())
	return offset
}

func (x *User) fastWriteField7(buf []byte) (offset int) {
	if x.BackgroundImage == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetBackgroundImage())
	return offset
}

func (x *User) fastWriteField8(buf []byte) (offset int) {
	if x.Signature == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetSignature())
	return offset
}

func (x *User) fastWriteField9(buf []byte) (offset int) {
	if x.TotalFavorited == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.GetTotalFavorited())
	return offset
}

func (x *User) fastWriteField10(buf []byte) (offset int) {
	if x.WorkCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, x.GetWorkCount())
	return offset
}

func (x *User) fastWriteField11(buf []byte) (offset int) {
	if x.FavoriteCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 11, x.GetFavoriteCount())
	return offset
}

func (x *DouyinCommentserverCommentcountRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinCommentserverCommentcountRequest) fastWriteField1(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVideoId())
	return offset
}

func (x *DouyinCommentserverCommentcountResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinCommentserverCommentcountResponse) fastWriteField1(buf []byte) (offset int) {
	if x.CommentCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetCommentCount())
	return offset
}

func (x *DouyinCommentserverCommentallcountRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinCommentserverCommentallcountRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.VideoIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetVideoIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetVideoIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *DouyinCommentserverCommentallcountResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinCommentserverCommentallcountResponse) fastWriteField1(buf []byte) (offset int) {
	if len(x.CommentCounts) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetCommentCounts()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetCommentCounts()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *DouyinCommentActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *DouyinCommentActionRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *DouyinCommentActionRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVideoId())
	return n
}

func (x *DouyinCommentActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetActionType())
	return n
}

func (x *DouyinCommentActionRequest) sizeField4() (n int) {
	if x.CommentText == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetCommentText())
	return n
}

func (x *DouyinCommentActionRequest) sizeField5() (n int) {
	if x.CommentId == nil {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetCommentId())
	return n
}

func (x *DouyinCommentActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinCommentActionResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DouyinCommentActionResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DouyinCommentActionResponse) sizeField3() (n int) {
	if x.Comment == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetComment())
	return n
}

func (x *Comment) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *Comment) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *Comment) sizeField2() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetUser())
	return n
}

func (x *Comment) sizeField3() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetContent())
	return n
}

func (x *Comment) sizeField4() (n int) {
	if x.CreateDate == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCreateDate())
	return n
}

func (x *DouyinCommentListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinCommentListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *DouyinCommentListRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVideoId())
	return n
}

func (x *DouyinCommentListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinCommentListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DouyinCommentListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DouyinCommentListResponse) sizeField3() (n int) {
	if x.CommentList == nil {
		return n
	}
	for i := range x.GetCommentList() {
		n += fastpb.SizeMessage(3, x.GetCommentList()[i])
	}
	return n
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.FollowCount == nil {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFollowCount())
	return n
}

func (x *User) sizeField4() (n int) {
	if x.FollowerCount == nil {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetFollowerCount())
	return n
}

func (x *User) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.GetIsFollow())
	return n
}

func (x *User) sizeField6() (n int) {
	if x.Avatar == nil {
		return n
	}
	n += fastpb.SizeString(6, x.GetAvatar())
	return n
}

func (x *User) sizeField7() (n int) {
	if x.BackgroundImage == nil {
		return n
	}
	n += fastpb.SizeString(7, x.GetBackgroundImage())
	return n
}

func (x *User) sizeField8() (n int) {
	if x.Signature == nil {
		return n
	}
	n += fastpb.SizeString(8, x.GetSignature())
	return n
}

func (x *User) sizeField9() (n int) {
	if x.TotalFavorited == nil {
		return n
	}
	n += fastpb.SizeInt64(9, x.GetTotalFavorited())
	return n
}

func (x *User) sizeField10() (n int) {
	if x.WorkCount == nil {
		return n
	}
	n += fastpb.SizeInt64(10, x.GetWorkCount())
	return n
}

func (x *User) sizeField11() (n int) {
	if x.FavoriteCount == nil {
		return n
	}
	n += fastpb.SizeInt64(11, x.GetFavoriteCount())
	return n
}

func (x *DouyinCommentserverCommentcountRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinCommentserverCommentcountRequest) sizeField1() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVideoId())
	return n
}

func (x *DouyinCommentserverCommentcountResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinCommentserverCommentcountResponse) sizeField1() (n int) {
	if x.CommentCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetCommentCount())
	return n
}

func (x *DouyinCommentserverCommentallcountRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinCommentserverCommentallcountRequest) sizeField1() (n int) {
	if len(x.VideoIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetVideoIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetVideoIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *DouyinCommentserverCommentallcountResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinCommentserverCommentallcountResponse) sizeField1() (n int) {
	if len(x.CommentCounts) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetCommentCounts()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetCommentCounts()[numIdxOrVal])
			return n
		})
	return n
}

var fieldIDToName_DouyinCommentActionRequest = map[int32]string{
	1: "UserId",
	2: "VideoId",
	3: "ActionType",
	4: "CommentText",
	5: "CommentId",
}

var fieldIDToName_DouyinCommentActionResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Comment",
}

var fieldIDToName_Comment = map[int32]string{
	1: "Id",
	2: "User",
	3: "Content",
	4: "CreateDate",
}

var fieldIDToName_DouyinCommentListRequest = map[int32]string{
	1: "UserId",
	2: "VideoId",
}

var fieldIDToName_DouyinCommentListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "CommentList",
}

var fieldIDToName_User = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "FollowCount",
	4:  "FollowerCount",
	5:  "IsFollow",
	6:  "Avatar",
	7:  "BackgroundImage",
	8:  "Signature",
	9:  "TotalFavorited",
	10: "WorkCount",
	11: "FavoriteCount",
}

var fieldIDToName_DouyinCommentserverCommentcountRequest = map[int32]string{
	1: "VideoId",
}

var fieldIDToName_DouyinCommentserverCommentcountResponse = map[int32]string{
	1: "CommentCount",
}

var fieldIDToName_DouyinCommentserverCommentallcountRequest = map[int32]string{
	1: "VideoIds",
}

var fieldIDToName_DouyinCommentserverCommentallcountResponse = map[int32]string{
	1: "CommentCounts",
}
