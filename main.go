// Package main implements a Composition Function.
package main

import (
	"github.com/alecthomas/kong"

	"github.com/crossplane/function-sdk-go"
)

// CLI of this Function.
type CLI struct {
	Debug bool `help:"Emit debug logs in addition to info logs." short:"d"`

	Network     string `default:"tcp"                                                                                        help:"Network on which to listen for gRPC connections."`
	Address     string `default:":9443"                                                                                      help:"Address at which to listen for gRPC connections."`
	TLSCertsDir string `env:"TLS_SERVER_CERTS_DIR"                                                                           help:"Directory containing server certs (tls.key, tls.crt) and the CA used to verify client certificates (ca.crt)"`
	Insecure    bool   `help:"Run without mTLS credentials. If you supply this flag --tls-server-certs-dir will be ignored."`
}

// Run this Function.
func (c *CLI) Run() error {
	log, err := function.NewLogger(c.Debug)
	if err != nil {
		return err
	}

	return function.Serve(&Function{log: log},
		function.Listen(c.Network, c.Address),
		function.MTLSCertificates(c.TLSCertsDir),
		function.Insecure(c.Insecure))
}

func main() {
	ctx := kong.Parse(&CLI{}, kong.Description("A Crossplane Composition Function that returns what you tell it to."))
	ctx.FatalIfErrorf(ctx.Run())
}
