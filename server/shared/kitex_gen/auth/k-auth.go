// Code generated by Kitex v0.11.3. DO NOT EDIT.

package auth

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *DeliverTokenReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUserId bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetUserId = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetUserId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DeliverTokenReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_DeliverTokenReq[fieldId]))
}

func (p *DeliverTokenReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.UserId = _field
	return offset, nil
}

// for compatibility
func (p *DeliverTokenReq) FastWrite(buf []byte) int {
	return 0
}

func (p *DeliverTokenReq) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *DeliverTokenReq) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *DeliverTokenReq) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 1)
	offset += thrift.Binary.WriteI32(buf[offset:], p.UserId)
	return offset
}

func (p *DeliverTokenReq) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I32Length()
	return l
}

func (p *VerifyTokenReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetToken bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetToken = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetToken {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_VerifyTokenReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_VerifyTokenReq[fieldId]))
}

func (p *VerifyTokenReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Token = _field
	return offset, nil
}

// for compatibility
func (p *VerifyTokenReq) FastWrite(buf []byte) int {
	return 0
}

func (p *VerifyTokenReq) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *VerifyTokenReq) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *VerifyTokenReq) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Token)
	return offset
}

func (p *VerifyTokenReq) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Token)
	return l
}

func (p *DeliveryResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetToken bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetToken = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetToken {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DeliveryResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_DeliveryResp[fieldId]))
}

func (p *DeliveryResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Token = _field
	return offset, nil
}

// for compatibility
func (p *DeliveryResp) FastWrite(buf []byte) int {
	return 0
}

func (p *DeliveryResp) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *DeliveryResp) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *DeliveryResp) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Token)
	return offset
}

func (p *DeliveryResp) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Token)
	return l
}

func (p *VerifyResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetRes bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetRes = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetRes {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_VerifyResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_VerifyResp[fieldId]))
}

func (p *VerifyResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field bool
	if v, l, err := thrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Res = _field
	return offset, nil
}

// for compatibility
func (p *VerifyResp) FastWrite(buf []byte) int {
	return 0
}

func (p *VerifyResp) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *VerifyResp) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *VerifyResp) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.BOOL, 1)
	offset += thrift.Binary.WriteBool(buf[offset:], p.Res)
	return offset
}

func (p *VerifyResp) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.BoolLength()
	return l
}

func (p *AuthServiceDeliverTokenByRPCArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AuthServiceDeliverTokenByRPCArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *AuthServiceDeliverTokenByRPCArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0
	_field := NewDeliverTokenReq()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = _field
	return offset, nil
}

// for compatibility
func (p *AuthServiceDeliverTokenByRPCArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *AuthServiceDeliverTokenByRPCArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *AuthServiceDeliverTokenByRPCArgs) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *AuthServiceDeliverTokenByRPCArgs) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 1)
	offset += p.Req.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *AuthServiceDeliverTokenByRPCArgs) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.Req.BLength()
	return l
}

func (p *AuthServiceDeliverTokenByRPCResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AuthServiceDeliverTokenByRPCResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *AuthServiceDeliverTokenByRPCResult) FastReadField0(buf []byte) (int, error) {
	offset := 0
	_field := NewDeliveryResp()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = _field
	return offset, nil
}

// for compatibility
func (p *AuthServiceDeliverTokenByRPCResult) FastWrite(buf []byte) int {
	return 0
}

func (p *AuthServiceDeliverTokenByRPCResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *AuthServiceDeliverTokenByRPCResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *AuthServiceDeliverTokenByRPCResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *AuthServiceDeliverTokenByRPCResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Success.BLength()
	}
	return l
}

func (p *AuthServiceVerifyTokenByRpcArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AuthServiceVerifyTokenByRpcArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *AuthServiceVerifyTokenByRpcArgs) FastReadField2(buf []byte) (int, error) {
	offset := 0
	_field := NewVerifyTokenReq()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = _field
	return offset, nil
}

// for compatibility
func (p *AuthServiceVerifyTokenByRpcArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *AuthServiceVerifyTokenByRpcArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *AuthServiceVerifyTokenByRpcArgs) BLength() int {
	l := 0
	if p != nil {
		l += p.field2Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *AuthServiceVerifyTokenByRpcArgs) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 2)
	offset += p.Req.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *AuthServiceVerifyTokenByRpcArgs) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.Req.BLength()
	return l
}

func (p *AuthServiceVerifyTokenByRpcResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AuthServiceVerifyTokenByRpcResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *AuthServiceVerifyTokenByRpcResult) FastReadField0(buf []byte) (int, error) {
	offset := 0
	_field := NewVerifyResp()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = _field
	return offset, nil
}

// for compatibility
func (p *AuthServiceVerifyTokenByRpcResult) FastWrite(buf []byte) int {
	return 0
}

func (p *AuthServiceVerifyTokenByRpcResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *AuthServiceVerifyTokenByRpcResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *AuthServiceVerifyTokenByRpcResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *AuthServiceVerifyTokenByRpcResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Success.BLength()
	}
	return l
}

func (p *AuthServiceDeliverTokenByRPCArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *AuthServiceDeliverTokenByRPCResult) GetResult() interface{} {
	return p.Success
}

func (p *AuthServiceVerifyTokenByRpcArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *AuthServiceVerifyTokenByRpcResult) GetResult() interface{} {
	return p.Success
}