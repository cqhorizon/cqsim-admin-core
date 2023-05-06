package service

import (
	"fmt"

	"cqsim-admin-core/logger"
)

type Service struct {
	Msg   string
	MsgID string
	Log   *logger.Helper
	Error error
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}
