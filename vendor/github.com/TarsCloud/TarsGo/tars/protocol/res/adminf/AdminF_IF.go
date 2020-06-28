//Package adminf comment
// This file war generated by tars2go 1.1
// Generated from AdminF.tars
package adminf

import (
	"context"
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//AdminF struct
type AdminF struct {
	s m.Servant
}

//Shutdown is the proxy function for the method defined in the tars file, with the context
func (_obj *AdminF) Shutdown(_opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "shutdown", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return nil
}

//ShutdownWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AdminF) ShutdownWithContext(ctx context.Context, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "shutdown", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return nil
}

//Notify is the proxy function for the method defined in the tars file, with the context
func (_obj *AdminF) Notify(Command string, _opt ...map[string]string) (ret string, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Command, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "notify", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_string(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//NotifyWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AdminF) NotifyWithContext(ctx context.Context, Command string, _opt ...map[string]string) (ret string, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Command, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "notify", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_string(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *AdminF) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *AdminF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

type _impAdminF interface {
	Shutdown() (err error)
	Notify(Command string) (ret string, err error)
}
type _impAdminFWithContext interface {
	Shutdown(ctx context.Context) (err error)
	Notify(ctx context.Context, Command string) (ret string, err error)
}

func shutdown(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	if withContext == false {
		_imp := _val.(_impAdminF)
		err = _imp.Shutdown()
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impAdminFWithContext)
		err = _imp.Shutdown(ctx)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func notify(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Command string
	err = _is.Read_string(&Command, 1, true)
	if err != nil {
		return err
	}
	if withContext == false {
		_imp := _val.(_impAdminF)
		ret, err := _imp.Notify(Command)
		if err != nil {
			return err
		}

		err = _os.Write_string(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impAdminFWithContext)
		ret, err := _imp.Notify(ctx, Command)
		if err != nil {
			return err
		}

		err = _os.Write_string(ret, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *AdminF) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "shutdown":
		err := shutdown(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "notify":
		err := notify(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
