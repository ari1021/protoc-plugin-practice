package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/arkuchy/protoc-plugin-practice/protoc-gen-customopt/generated"
	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/types/descriptorpb"
)

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var req plugin.CodeGeneratorRequest
	if err = proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func processReq(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}
	var resp plugin.CodeGeneratorResponse
	for _, fname := range req.FileToGenerate {
		f := files[fname]
		out := fname + ".json"
		messages := f.MessageType
		var whiteMessages []*descriptorpb.DescriptorProto
		for _, message := range messages {
			if isTarget(message) {
				whiteMessages = append(whiteMessages, message)
			}
		}
		b, err := json.Marshal(whiteMessages)
		if err != nil {
			log.Fatalln(err)
		}
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(out),
			Content: proto.String(string(b)),
		})
	}
	return &resp
}

func isTarget(m *descriptorpb.DescriptorProto) bool {
	opt := m.GetOptions()
	if opt == nil {
		return false
	}
	ext, err := proto.GetExtension(opt, generated.E_MessageList)
	if err == proto.ErrMissingExtension {
		return false
	}
	if err != nil {
		log.Fatalln(err)
	}
	mopt := ext.(*generated.MessageListOptions)
	return mopt.Target
}

func emitResp(resp *plugin.CodeGeneratorResponse) error {
	buf, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(buf)
	return err
}

func run() error {
	req, err := parseReq(os.Stdin)
	if err != nil {
		return err
	}
	resp := processReq(req)
	return emitResp(resp)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
