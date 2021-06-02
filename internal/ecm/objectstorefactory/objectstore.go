package objectstorefactory

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/engine/aws"
)

func GetObjectStore(config cfg.AppConfig) ce.ObjectStore {

	switch config.StoreType {
	case cfg.StoreTypeAWS:
		return aws.GetObjectStore(&config)
	default:
		return nil
	}

}
