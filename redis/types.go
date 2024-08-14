package urchin_util

import (
	"errors"
)

const (
	StoragePrefix = "urchin"
)

var (
	ErrorInvalidParameter = errors.New("invalid parameters")
	ErrorNotExists        = errors.New("not exists")
	ErrorNotAllowed       = errors.New("too much task, wait a moment and try again")
	ErrorInternal         = errors.New("internal error")
	ErrorUnknown          = errors.New("unknown error")
	ErrorNotImplement     = errors.New("not implement")
)
