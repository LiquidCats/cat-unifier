package factories

import (
	"cat-unifier/internal/kernel/common/contracts"
	"cat-unifier/internal/kernel/common/library/repositiries"
	"cat-unifier/pkg/reader/eth"
	"errors"
	"fmt"
)

type IReaderFactory interface {
	Make() (contracts.IReader, error)
}

type readerFactory struct {
	cfg repositiries.IConfigRepository
}

func (r *readerFactory) makeEth() contracts.IReader {
	node := r.cfg.Get("reader.eth.node", nil)
	return eth.NewReader(fmt.Sprintf("%s", node))
}

func (r *readerFactory) Make() (contracts.IReader, error) {
	switch r.cfg.Get("app.currency", nil) {
	case "eth":
		return r.makeEth(), nil
	}

	return nil, errors.New("not supported currency")
}

func NewReaderFactory(cfg repositiries.IConfigRepository) IReaderFactory {
	return &readerFactory{cfg: cfg}
}
