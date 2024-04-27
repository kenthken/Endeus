package helper_test

import (
	exceptions "endeus/api/expections"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) exceptions.BaseErrorResponse {
	err := recover()
	if err != nil {
		tx.Rollback()
		logrus.Info(err)
		return exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
		}
	} else {
		tx.Commit()
		return exceptions.BaseErrorResponse{}
	}
}
