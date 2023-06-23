package services

import (
	"bank-api/internal/data"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *BankService) GetMutasi(no_rekening string) (response []data.Mutasi, err error) {
	startTime := time.Now()
	// init transaction
	tx, err := s.repository.Begin()
	if err != nil {
		err = fmt.Errorf("failed to begin transaction")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// check isRekeningValid by no_rekening
	isRekeningValid, err := s.repository.IsRekeningValid(tx, no_rekening)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	if !isRekeningValid {
		err = fmt.Errorf("INVALID")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return
	}

	// get mutasi rekening
	response, err = s.repository.GetMutasiByRekening(tx, no_rekening)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	elapsedTime := time.Since(startTime)

	if err != nil {
		s.log.Warn(logrus.Fields{"elapsed_time": elapsedTime}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	s.log.Info(
		logrus.Fields{"elapsed_time": elapsedTime}, nil, "Executed: BankRepository.GetMutasi with no error",
	)

	// Commit
	s.repository.Commit(tx)

	return
}
