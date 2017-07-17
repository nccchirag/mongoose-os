// Code generated by clubbygen.
// GENERATED FILE DO NOT EDIT
// +build clubby_strict

package i2c

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"cesanta.com/common/go/mgrpc"
	"cesanta.com/common/go/mgrpc/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/context"
	"golang.org/x/net/trace"

	"github.com/cesanta/ucl"
	"github.com/cesanta/validate-json/schema"
	"github.com/golang/glog"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://mongoose-iot.com/fwI2C"

type ReadArgs struct {
	Addr *int64 `json:"addr,omitempty"`
	Len  *int64 `json:"len,omitempty"`
}

type ReadResult struct {
	Data_hex *string `json:"data_hex,omitempty"`
}

type ReadRegBArgs struct {
	Addr *int64 `json:"addr,omitempty"`
	Reg  *int64 `json:"reg,omitempty"`
}

type ReadRegBResult struct {
	Value *int64 `json:"value,omitempty"`
}

type ReadRegWArgs struct {
	Addr *int64 `json:"addr,omitempty"`
	Reg  *int64 `json:"reg,omitempty"`
}

type ReadRegWResult struct {
	Value *int64 `json:"value,omitempty"`
}

type WriteArgs struct {
	Addr     *int64  `json:"addr,omitempty"`
	Data_hex *string `json:"data_hex,omitempty"`
}

type WriteRegBArgs struct {
	Addr  *int64 `json:"addr,omitempty"`
	Reg   *int64 `json:"reg,omitempty"`
	Value *int64 `json:"value,omitempty"`
}

type Service interface {
	Read(ctx context.Context, args *ReadArgs) (*ReadResult, error)
	ReadRegB(ctx context.Context, args *ReadRegBArgs) (*ReadRegBResult, error)
	ReadRegW(ctx context.Context, args *ReadRegWArgs) (*ReadRegWResult, error)
	Scan(ctx context.Context) ([]int64, error)
	Write(ctx context.Context, args *WriteArgs) error
	WriteRegB(ctx context.Context, args *WriteRegBArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
}

type _validators struct {
	// This comment prevents gofmt from aligning types in the struct.
	ReadArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ReadResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ReadRegBArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ReadRegBResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ReadRegWArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ReadRegWResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ScanResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	WriteArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	WriteRegBArgs *schema.Validator
}

var (
	validators     *_validators
	validatorsOnce sync.Once
)

func initValidators() {
	validators = &_validators{}

	loader := schema.NewLoader()

	service, err := ucl.Parse(bytes.NewBuffer(_ServiceDefinition))
	if err != nil {
		panic(err)
	}
	// Patch up shortcuts to be proper schemas.
	for _, v := range service.(*ucl.Object).Find("methods").(*ucl.Object).Value {
		if s, ok := v.(*ucl.Object).Find("result").(*ucl.String); ok {
			for kk := range v.(*ucl.Object).Value {
				if kk.Value == "result" {
					v.(*ucl.Object).Value[kk] = &ucl.Object{
						Value: map[ucl.Key]ucl.Value{
							ucl.Key{Value: "type"}: s,
						},
					}
				}
			}
		}
		if v.(*ucl.Object).Find("args") == nil {
			continue
		}
		args := v.(*ucl.Object).Find("args").(*ucl.Object)
		for kk, vv := range args.Value {
			if s, ok := vv.(*ucl.String); ok {
				args.Value[kk] = &ucl.Object{
					Value: map[ucl.Key]ucl.Value{
						ucl.Key{Value: "type"}: s,
					},
				}
			}
		}
	}
	var s *ucl.Object
	_ = s // avoid unused var error
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Read").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Read").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.ReadArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.ReadResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Read").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("ReadRegB").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("ReadRegB").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.ReadRegBArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.ReadRegBResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("ReadRegB").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("ReadRegW").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("ReadRegW").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.ReadRegWArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.ReadRegWResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("ReadRegW").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	validators.ScanResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Scan").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Write").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Write").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.WriteArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("WriteRegB").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("WriteRegB").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.WriteRegBArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
}

func NewClient(i Instance, addr string) Service {
	validatorsOnce.Do(initValidators)
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Read(ctx context.Context, args *ReadArgs) (res *ReadResult, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.Read",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Read: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Read")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ReadResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for Read: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for Read")
			}
		}
	}
	var r *ReadResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) ReadRegB(ctx context.Context, args *ReadRegBArgs) (res *ReadRegBResult, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.ReadRegB",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadRegBArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for ReadRegB: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for ReadRegB")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ReadRegBResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for ReadRegB: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for ReadRegB")
			}
		}
	}
	var r *ReadRegBResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) ReadRegW(ctx context.Context, args *ReadRegWArgs) (res *ReadRegWResult, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.ReadRegW",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadRegWArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for ReadRegW: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for ReadRegW")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ReadRegWResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for ReadRegW: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for ReadRegW")
			}
		}
	}
	var r *ReadRegWResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Scan(ctx context.Context) (res []int64, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.Scan",
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ScanResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for Scan: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for Scan")
			}
		}
	}
	var r []int64
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Write(ctx context.Context, args *WriteArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "I2C.Write",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.WriteArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Write: %+v", err)
				return errors.Annotatef(err, "invalid args for Write")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) WriteRegB(ctx context.Context, args *WriteRegBArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "I2C.WriteRegB",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.WriteRegBArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for WriteRegB: %+v", err)
				return errors.Annotatef(err, "invalid args for WriteRegB")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

//func RegisterService(i *clubby.Instance, impl Service) error {
//validatorsOnce.Do(initValidators)
//s := &_Server{impl}
//i.RegisterCommandHandler("I2C.Read", s.Read)
//i.RegisterCommandHandler("I2C.ReadRegB", s.ReadRegB)
//i.RegisterCommandHandler("I2C.ReadRegW", s.ReadRegW)
//i.RegisterCommandHandler("I2C.Scan", s.Scan)
//i.RegisterCommandHandler("I2C.Write", s.Write)
//i.RegisterCommandHandler("I2C.WriteRegB", s.WriteRegB)
//i.RegisterService(ServiceID, _ServiceDefinition)
//return nil
//}

type _Server struct {
	impl Service
}

func (s *_Server) Read(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Read: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Read")
			}
		}
	}
	var args ReadArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	r, err := s.impl.Read(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ReadResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for Read: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for Read")
			}
		}
	}
	return r, nil
}

func (s *_Server) ReadRegB(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadRegBArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for ReadRegB: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for ReadRegB")
			}
		}
	}
	var args ReadRegBArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	r, err := s.impl.ReadRegB(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ReadRegBResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for ReadRegB: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for ReadRegB")
			}
		}
	}
	return r, nil
}

func (s *_Server) ReadRegW(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.ReadRegWArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for ReadRegW: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for ReadRegW")
			}
		}
	}
	var args ReadRegWArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	r, err := s.impl.ReadRegW(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ReadRegWResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for ReadRegW: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for ReadRegW")
			}
		}
	}
	return r, nil
}

func (s *_Server) Scan(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	r, err := s.impl.Scan(ctx)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ScanResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for Scan: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for Scan")
			}
		}
	}
	return r, nil
}

func (s *_Server) Write(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.WriteArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Write: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Write")
			}
		}
	}
	var args WriteArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Write(ctx, &args)
}

func (s *_Server) WriteRegB(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.WriteRegBArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for WriteRegB: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for WriteRegB")
			}
		}
	}
	var args WriteRegBArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.WriteRegB(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "methods": {
    "Read": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "len": {
          "doc": "Number of bytes to read.",
          "type": "integer"
        }
      },
      "doc": "Read specified number of bytes from the specified device.",
      "result": {
        "properties": {
          "data_hex": {
            "doc": "Hex-encoded data.",
            "type": "string"
          }
        },
        "type": "object"
      }
    },
    "ReadRegB": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "reg": {
          "doc": "Register number.",
          "type": "integer"
        }
      },
      "doc": "Read value of a byte-sized (8-bit) register.",
      "result": {
        "properties": {
          "value": {
            "doc": "Register value read from the device.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "ReadRegW": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "reg": {
          "doc": "Register number.",
          "type": "integer"
        }
      },
      "doc": "Read value of a word-sized (16-bit) register.",
      "result": {
        "properties": {
          "value": {
            "doc": "Register value read from the device.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "Scan": {
      "doc": "Scan the I2C bus, returning addresses of devices that responded.",
      "result": {
        "items": {
          "doc": "List of device addresses for which an ACK was received.",
          "type": "integer"
        },
        "type": "array"
      }
    },
    "Write": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "data_hex": {
          "doc": "Hex-encoded data to write.",
          "type": "string"
        }
      },
      "doc": "Write the specified data to the device with the specified address."
    },
    "WriteRegB": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "reg": {
          "doc": "Register number.",
          "type": "integer"
        },
        "value": {
          "doc": "Register value to write.",
          "type": "integer"
        }
      },
      "doc": "Write value of a word-sized (16-bit) register."
    }
  },
  "name": "I2C",
  "namespace": "http://mongoose-iot.com/fw"
}`))
