package main

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

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

	rsp := NewResponseTo(req, DefaultTTL)

	in := &v1beta1.Response{}
	if err := GetObject(in, req.GetInput()); err != nil {
		Fatal(rsp, errors.Wrapf(err, "cannot get Function input from %T", req))
		return rsp, nil
	}

	// We must pass through any desired state we're unconcerned with unmutated.
	// protojson.Unmarshal clears the message it's passed, so we can't just
	// unmarshal into rsp. Instead we unmarshal into an empty response, then
	// merge that into rsp.
	overlay := &fnv1beta1.RunFunctionResponse{}
	if err := protojson.Unmarshal(in.Response.Raw, overlay); err != nil {
		Fatal(rsp, errors.Wrapf(err, "cannot unmarshal RunFunctionResponse from %T", req))
		return rsp, nil
	}

	proto.Merge(rsp, overlay)

	return rsp, nil
}
