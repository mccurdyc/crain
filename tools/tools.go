// +build tools

// This follows the agreed-upon current best approach for adding developer tooling to your Module.
//
// For more details, refer to this section in the Go Wiki:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/quasilyte/go-ruleguard/cmd/ruleguard"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
