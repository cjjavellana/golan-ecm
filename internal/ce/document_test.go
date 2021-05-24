package ce

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestName(t *testing.T) {

	d := Document{
		Object{
			ObjectId: uuid.New(),
		},
	}

	log.Info(d.ObjectId)
}