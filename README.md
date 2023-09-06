# function-dummy

A [Crossplane] Composition Function that returns what you tell it to.

## What is this?

This [Composition Function][function-design] just returns whatever
`RunFunctionResponse` you tell it to. You provide a `RunFunctionResponse` in the
Function's input. This Function is mostly useful for testing Crossplane itself.

Note that this is a beta-style Function. It won't work with Crossplane v1.13 or
earlier - it targets the [implementation of Functions][function-pr] coming with
Crossplane v1.14 in late October.

Here's an example:

```yaml
apiVersion: apiextensions.crossplane.io/v1
kind: Composition
metadata:
  name: test-crossplane
spec:
  compositeTypeRef:
    apiVersion: database.example.com/v1alpha1
    kind: NoSQL
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
              apiVersion: example.org/v1
              kind: XR
              spec: {} # etc...
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

## Developing

This Function doesn't use the typical Crossplane build submodule and Makefile,
since we'd like Functions to have a less heavyweight developer experience.
It mostly relies on regular old Go tools:

```shell
# Run code generation - see input/generate.go
$ go generate ./...

# Run tests
$ go test -cover ./...
?       github.com/negz/function-patch-and-transform/input/v1beta1      [no test files]
ok      github.com/negz/function-patch-and-transform    0.021s  coverage: 76.1% of statements

# Lint the code
$ docker run --rm -v $(pwd):/app -v ~/.cache/golangci-lint/v1.54.2:/root/.cache -w /app golangci/golangci-lint:v1.54.2 golangci-lint run

# Build a Docker image - see Dockerfile
$ docker build .
```

[Crossplane]: https://crossplane.io
[function-design]: https://github.com/crossplane/crossplane/blob/3996f20/design/design-doc-composition-functions.md
[function-pr]: https://github.com/crossplane/crossplane/pull/4500
[docs-composition]: https://docs.crossplane.io/v1.13/getting-started/provider-aws-part-2/#create-a-deployment-template
[#2581]: https://github.com/crossplane/crossplane/issues/2581
