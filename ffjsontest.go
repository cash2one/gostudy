package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"code.google.com/p/goprotobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pquerna/ffjson/ffjson"
)

var ()

type Test struct {
	*string `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}
type Test1 struct {
	Account *string `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}
type PushAccountVerifyLoginUserPmd struct {
	Accid            *uint64 `protobuf:"varint,1,opt,name=accid" json:"accid,omitempty"`
	Account          *string `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	Zoneid           *uint32 `protobuf:"varint,3,opt,name=zoneid" json:"zoneid,omitempty"`
	Token            *string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Version          *uint32 `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Mid              *string `protobuf:"bytes,6,opt,name=mid" json:"mid,omitempty"`
	Gameversion      *uint32 `protobuf:"varint,7,opt,name=gameversion" json:"gameversion,omitempty"`
	Compress         *string `protobuf:"bytes,8,opt,name=compress" json:"compress,omitempty"`
	Encrypt          *string `protobuf:"bytes,9,opt,name=encrypt" json:"encrypt,omitempty"`
	Encryptkey       *string `protobuf:"bytes,10,opt,name=encryptkey" json:"encryptkey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func BuildProtoFromJson(typ reflect.Type, cmdjson string) []byte {
	proto_cmd := reflect.New(typ).Interface().(proto.Message)
	rawdata := []byte(cmdjson)
	json.Unmarshal(rawdata, proto_cmd)
	sendbuf := proto.NewBuffer(nil)
	sendbuf.Marshal(proto_cmd)
	return sendbuf.Bytes()
}

func BuildJsonFromProto(cmdname string, cmddata []byte) string {
	recvbuf := proto.NewBuffer(cmddata)
	recv := &AccountTokenVerifyLoginUserPmd_CS{} //难点,这里这个结构是不确定的,只能动态描述
	recvbuf.Unmarshal(recv)
	recv_json, _ := json.Marshal(recv)
	return string(recv_json)
}

type RequestUpFrameSyncNullUserPmd_C struct {
	Data             []byte `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func main() {
	account := "wangaijun"
	token := "Token"
	send := &AccountTokenVerifyLoginUserPmd_CS{ //难点,这里这个结构是不确定的,只能动态描述
		Account: &account,
		Token:   &token,
	}
	now := time.Now().UnixNano()
	var send_json []byte
	for i := 0; i < 1000000; i++ {
		send_json, _ = json.Marshal(send)
	}
	fmt.Println("  json:", time.Now().UnixNano()-now)
	now = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		send_json, _ = ffjson.Marshal(send)
		//ffjson.Pool(send_json)
	}
	fmt.Println("ffjson:", time.Now().UnixNano()-now)
	fmt.Println(string(send_json))
	return
	bytes := BuildProtoFromJson(reflect.TypeOf(*send), string(send_json))

	//网络发送部分

	recv := bytes
	recv_json := BuildJsonFromProto("AccountTokenVerifyLoginUserPmd_CS", recv)
	fmt.Println("BuildJsonFromProto", string(recv_json))
	test1 := &Test1{Account: &account}
	test := &Test{}
	anypb := &any.Any{}
	test1_json, _ := json.Marshal(test1)
	json.Unmarshal(test1_json, test)
	fmt.Println(json.Unmarshal(test1_json, anypb))
	fmt.Println(reflect.ValueOf(*test).Field(0).Elem().String())
	fmt.Println(*test)
	fmt.Println("xxxxx:", anypb.Value)
	data := `{"data":"[whj]"}`
	ru := &RequestUpFrameSyncNullUserPmd_C{}
	fmt.Println(json.Unmarshal([]byte(data), ru), ru.Data)
}
