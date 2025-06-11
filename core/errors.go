package core

import "errors"

var (
	ErrChainDatabaseClosed    = errors.New("chain's database seems to be closed")
	ErrChainWriteBufferClosed = errors.New("chain's write buffer seems to be closed")
)

var (
	ErrChainSanityCheckFailed = errors.New("the chain is not valid")
)
