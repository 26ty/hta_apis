package jwe

import (
	model "eirc.app/internal/v1/structure/jwe"
)

type Service interface {
	Created(input *model.JWE) (output *model.Token, err error)
}

type service struct {
}

func New() Service {
	return &service{}
}
