package main

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"google.golang.org/protobuf/encoding/protojson"

	fnv1beta1 "github.com/crossplane/function-sdk-go/proto/v1beta1"

	"github.com/crossplane-contrib/function-dummy/input/v1beta1"
)

// Function returns whatever response you ask it to.
type Function struct {
	fnv1beta1.UnimplementedFunctionRunnerServiceServer

	log logging.Logger
}

// RunFunction runs the Function.
func (f *Function) RunFunction(ctx context.Context, req *fnv1beta1.RunFunctionRequest) (*fnv1beta1.RunFunctionResponse, error) {
	f.log.Info("Running Function", "tag", req.GetMeta().GetTag())

	in := &v1beta1.Response{}
	if err := GetObject(in, req.GetInput()); err != nil {
		rsp := &fnv1beta1.RunFunctionResponse{}
		Fatal(rsp, errors.Wrapf(err, "cannot get Function input from %T", req))
		return rsp, nil
	}

	rsp := &fnv1beta1.RunFunctionResponse{}
	if err := protojson.Unmarshal(in.Response.Raw, rsp); err != nil {
		rsp := &fnv1beta1.RunFunctionResponse{}
		Fatal(rsp, errors.Wrapf(err, "cannot unmarshal RunFunctionResponse from %T", req))
		return rsp, nil
	}

	// Copy over the tag (if any) from our request like a real Function should.
	if rsp.Meta == nil {
		rsp.Meta = &fnv1beta1.ResponseMeta{}
	}
	rsp.Meta.Tag = req.GetMeta().GetTag()

	return rsp, nil
}
