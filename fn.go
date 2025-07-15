package main

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/crossplane/function-sdk-go/errors"
	"github.com/crossplane/function-sdk-go/logging"
	fnv1 "github.com/crossplane/function-sdk-go/proto/v1"
	"github.com/crossplane/function-sdk-go/request"
	"github.com/crossplane/function-sdk-go/response"

	"github.com/crossplane-contrib/function-dummy/input/v1beta1"
)

// Function returns whatever response you ask it to.
type Function struct {
	fnv1.UnimplementedFunctionRunnerServiceServer

	log logging.Logger
}

// RunFunction runs the Function.
func (f *Function) RunFunction(_ context.Context, req *fnv1.RunFunctionRequest) (*fnv1.RunFunctionResponse, error) {
	f.log.Info("Running Function", "tag", req.GetMeta().GetTag())

	rsp := response.To(req, response.DefaultTTL)

	in := &v1beta1.Response{}
	if err := request.GetInput(req, in); err != nil {
		response.Fatal(rsp, errors.Wrapf(err, "cannot get Function input from %T", req))
		return rsp, nil
	}

	// We must pass through any desired state we're unconcerned with unmutated.
	// protojson.Unmarshal clears the message it's passed, so we can't just
	// unmarshal into rsp. Instead we unmarshal into an empty response, then
	// merge that into rsp.
	overlay := &fnv1.RunFunctionResponse{}
	if err := protojson.Unmarshal(in.Response.Raw, overlay); err != nil {
		response.Fatal(rsp, errors.Wrapf(err, "cannot unmarshal RunFunctionResponse from %T", req))
		return rsp, nil
	}

	proto.Merge(rsp, overlay)

	return rsp, nil
}
