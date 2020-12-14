package main

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/transports/alltransports"
)

func main() {
	err := run()
	fmt.Println(err)
}

func run() error {
	policyContext, err := signature.NewPolicyContext(&signature.Policy{
		Default: []signature.PolicyRequirement{
			// NewPRInsecureAcceptAnything returns a new "insecureAcceptAnything" PolicyRequirement.
			signature.NewPRInsecureAcceptAnything(),
		},
	})
	if err != nil {
		return err
	}
	defer policyContext.Destroy()

	src, err := alltransports.ParseImageName("docker://ghcr.io/cybozu-go/coil:2.0.3")
	// src, err := alltransports.ParseImageName("docker://quay.io/cybozu/squid:3.5.27.1.11")
	if err != nil {
		return err
	}
	dst, err := alltransports.ParseImageName("docker-archive:./tmp.img")
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = copy.Image(ctx, policyContext, dst, src, nil)
	return err
}
