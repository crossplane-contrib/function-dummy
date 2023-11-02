# function-dummy
[![CI](https://github.com/crossplane-contrib/function-dummy/actions/workflows/ci.yml/badge.svg)](https://github.com/crossplane-contrib/function-dummy/actions/workflows/ci.yml) ![GitHub release (latest SemVer)](https://img.shields.io/github/release/crossplane-contrib/function-dummy)


This [composition function][docs-functions] returns whatever you tell it to.

Provide a YAML-serialized [`RunFunctionResponse`][bsr] as the function's input.
The response's `desired` object will be merged onto any desired state that was
passed to the function using [`proto.Merge`][merge] semantics.

Here's an example:

```yaml
apiVersion: apiextensions.crossplane.io/v1
kind: Composition
metadata:
  name: test-crossplane
spec:
  compositeTypeRef:
    apiVersion: example.crossplane.io/v1beta1
    kind: XR
  pipeline:
  - step: return-some-results
    functionRef:
      name: function-dummy
    input:
      apiVersion: dummy.fn.crossplane.io/v1beta1
      kind: Response
      response:
        desired:
          composite:
            resource:
              status:
                widgets: 200
            connectionDetails:
              very: secret
          resources:
            cool-resource:
              resource:
                apiVersion: example.org/v1
                kind: Composed
                spec: {} # etc...
              ready: READY_TRUE
        results:
        - severity: SEVERITY_NORMAL
          message: "I did the thing!"
```

See the [example](example) directory for an example you can run locally using
the Crossplane CLI:

```shell
$ crossplane beta render xr.yaml composition.yaml functions.yaml
```

See the [composition functions documentation][docs-functions] to learn more
about `crossplane beta render`.

## Developing this function

This function uses [Go][go], [Docker][docker], and the [Crossplane CLI][cli] to
build functions.

```shell
# Run code generation - see input/generate.go
$ go generate ./...

# Run tests - see fn_test.go
$ go test ./...

# Build the function's runtime image - see Dockerfile
$ docker build . --tag=runtime

# Build a function package - see package/crossplane.yaml
$ crossplane xpkg build -f package --embed-runtime-image=runtime
```

[docs-functions]: https://docs.crossplane.io/v1.14/concepts/composition-functions/
[bsr]: https://buf.build/crossplane/crossplane/docs/main:apiextensions.fn.proto.v1beta1#apiextensions.fn.proto.v1beta1.RunFunctionResponse
[merge]: https://pkg.go.dev/github.com/golang/protobuf/proto#Merge
[go]: https://go.dev
[docker]: https://www.docker.com
[cli]: https://docs.crossplane.io/latest/cli