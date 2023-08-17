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

func (x *DouyinUserCreateBaseUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserCreateBaseUserRequest[number], err)
}

func (x *DouyinUserCreateBaseUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinUserCreateBaseUserRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinUserCreateBaseUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserCreateBaseUserResponse[number], err)
}

func (x *DouyinUserCreateBaseUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinUserCreateBaseUserResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinUserCreateBaseUserResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinUserFindBaseUserByNameRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserFindBaseUserByNameRequest[number], err)
}

func (x *DouyinUserFindBaseUserByNameRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinUserFindBaseUserByNameResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserFindBaseUserByNameResponse[number], err)
}

func (x *DouyinUserFindBaseUserByNameResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinUserFindBaseUserByNameResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinUserFindBaseUserByNameResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v BaseUser
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseUser = &v
	return offset, nil
}

func (x *DouyinUserFindBaseUserByIdRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserFindBaseUserByIdRequest[number], err)
}

func (x *DouyinUserFindBaseUserByIdRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinUserFindBaseUserByIdResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserFindBaseUserByIdResponse[number], err)
}

func (x *DouyinUserFindBaseUserByIdResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinUserFindBaseUserByIdResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinUserFindBaseUserByIdResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v BaseUser
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseUser = &v
	return offset, nil
}

func (x *DouyinUserFindBaseUserListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserFindBaseUserListRequest[number], err)
}

func (x *DouyinUserFindBaseUserListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.AuthorId = append(x.AuthorId, v)
			return offset, err
		})
	return offset, err
}

func (x *DouyinUserFindBaseUserListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinUserFindBaseUserListResponse[number], err)
}

func (x *DouyinUserFindBaseUserListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinUserFindBaseUserListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *DouyinUserFindBaseUserListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v BaseUser
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseUser = append(x.BaseUser, &v)
	return offset, nil
}

func (x *BaseUser) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseUser[number], err)
}

func (x *BaseUser) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseUser) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseUser) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Avatar = &tmp
	return offset, err
}

func (x *BaseUser) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.BackgroundImage = &tmp
	return offset, err
}

func (x *BaseUser) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Signature = &tmp
	return offset, err
}

func (x *DouyinUserCreateBaseUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinUserCreateBaseUserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *DouyinUserCreateBaseUserRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *DouyinUserCreateBaseUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinUserCreateBaseUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DouyinUserCreateBaseUserResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DouyinUserCreateBaseUserResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *DouyinUserFindBaseUserByNameRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinUserFindBaseUserByNameRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *DouyinUserFindBaseUserByNameResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinUserFindBaseUserByNameResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DouyinUserFindBaseUserByNameResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DouyinUserFindBaseUserByNameResponse) fastWriteField3(buf []byte) (offset int) {
	if x.BaseUser == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetBaseUser())
	return offset
}

func (x *DouyinUserFindBaseUserByIdRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinUserFindBaseUserByIdRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *DouyinUserFindBaseUserByIdResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinUserFindBaseUserByIdResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DouyinUserFindBaseUserByIdResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DouyinUserFindBaseUserByIdResponse) fastWriteField3(buf []byte) (offset int) {
	if x.BaseUser == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetBaseUser())
	return offset
}

func (x *DouyinUserFindBaseUserListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DouyinUserFindBaseUserListRequest) fastWriteField2(buf []byte) (offset int) {
	if len(x.AuthorId) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 2, len(x.GetAuthorId()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetAuthorId()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *DouyinUserFindBaseUserListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinUserFindBaseUserListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DouyinUserFindBaseUserListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DouyinUserFindBaseUserListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.BaseUser == nil {
		return offset
	}
	for i := range x.GetBaseUser() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetBaseUser()[i])
	}
	return offset
}

func (x *BaseUser) FastWrite(buf []byte) (offset int) {
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

func (x *BaseUser) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *BaseUser) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *BaseUser) fastWriteField3(buf []byte) (offset int) {
	if x.Avatar == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAvatar())
	return offset
}

func (x *BaseUser) fastWriteField4(buf []byte) (offset int) {
	if x.BackgroundImage == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetBackgroundImage())
	return offset
}

func (x *BaseUser) fastWriteField5(buf []byte) (offset int) {
	if x.Signature == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetSignature())
	return offset
}

func (x *DouyinUserCreateBaseUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DouyinUserCreateBaseUserRequest) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *DouyinUserCreateBaseUserRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *DouyinUserCreateBaseUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinUserCreateBaseUserResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DouyinUserCreateBaseUserResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DouyinUserCreateBaseUserResponse) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetUserId())
	return n
}

func (x *DouyinUserFindBaseUserByNameRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinUserFindBaseUserByNameRequest) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *DouyinUserFindBaseUserByNameResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinUserFindBaseUserByNameResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DouyinUserFindBaseUserByNameResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DouyinUserFindBaseUserByNameResponse) sizeField3() (n int) {
	if x.BaseUser == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetBaseUser())
	return n
}

func (x *DouyinUserFindBaseUserByIdRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinUserFindBaseUserByIdRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *DouyinUserFindBaseUserByIdResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinUserFindBaseUserByIdResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DouyinUserFindBaseUserByIdResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DouyinUserFindBaseUserByIdResponse) sizeField3() (n int) {
	if x.BaseUser == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetBaseUser())
	return n
}

func (x *DouyinUserFindBaseUserListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField2()
	return n
}

func (x *DouyinUserFindBaseUserListRequest) sizeField2() (n int) {
	if len(x.AuthorId) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(2, len(x.GetAuthorId()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetAuthorId()[numIdxOrVal])
			return n
		})
	return n
}

func (x *DouyinUserFindBaseUserListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinUserFindBaseUserListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DouyinUserFindBaseUserListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DouyinUserFindBaseUserListResponse) sizeField3() (n int) {
	if x.BaseUser == nil {
		return n
	}
	for i := range x.GetBaseUser() {
		n += fastpb.SizeMessage(3, x.GetBaseUser()[i])
	}
	return n
}

func (x *BaseUser) Size() (n int) {
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

func (x *BaseUser) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *BaseUser) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *BaseUser) sizeField3() (n int) {
	if x.Avatar == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetAvatar())
	return n
}

func (x *BaseUser) sizeField4() (n int) {
	if x.BackgroundImage == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetBackgroundImage())
	return n
}

func (x *BaseUser) sizeField5() (n int) {
	if x.Signature == nil {
		return n
	}
	n += fastpb.SizeString(5, x.GetSignature())
	return n
}

var fieldIDToName_DouyinUserCreateBaseUserRequest = map[int32]string{
	1: "Username",
	2: "Password",
}

var fieldIDToName_DouyinUserCreateBaseUserResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserId",
}

var fieldIDToName_DouyinUserFindBaseUserByNameRequest = map[int32]string{
	1: "Username",
}

var fieldIDToName_DouyinUserFindBaseUserByNameResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "BaseUser",
}

var fieldIDToName_DouyinUserFindBaseUserByIdRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_DouyinUserFindBaseUserByIdResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "BaseUser",
}

var fieldIDToName_DouyinUserFindBaseUserListRequest = map[int32]string{
	2: "AuthorId",
}

var fieldIDToName_DouyinUserFindBaseUserListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "BaseUser",
}

var fieldIDToName_BaseUser = map[int32]string{
	1: "Id",
	2: "Name",
	3: "Avatar",
	4: "BackgroundImage",
	5: "Signature",
}
