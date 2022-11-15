package v1

import (
	"github.com/SaidovZohid/blog_db/config"
	"github.com/SaidovZohid/blog_db/storage"
)

type handlerV1 struct {
	cfg *config.Config
	Storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg: options.Cfg,
		Storage: options.Storage,
	}
}