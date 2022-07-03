package restapi

import "github.com/noppawitt/go-petstore/spec"

type Handler struct{}

var _ spec.ServerInterface = (*Handler)(nil)
