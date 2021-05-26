package objectstore

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/core/ce"
	"cjavellana.me/ecm/golan/internal/core/engine/aws"
)

func Get(config cfg.Config) ce.ObjectStore {

	switch config.StoreType {
	case cfg.StoreTypeAWS:
		return aws.GetObjectStore()
	default:
		return nil
	}

}
