package services

import (
	"bank-api/internal/data"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *BankService) Register(requestPayload data.RegisterRequest) (no_rekening string, err error) {
	// init transaction
	tx, err := s.repository.Begin()
	if err != nil {
		err = fmt.Errorf("failed to begin transaction")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// check isNasabahExist by nohp or nik
	isNasabahExist, err := s.repository.IsNasabahExist(tx, requestPayload)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	if isNasabahExist {
		err = fmt.Errorf("EXIST")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return "", err
	}

	startTime := time.Now()
	no_nasabah, err := s.repository.RegisterNasabah(tx, requestPayload)
	elapsedTime := time.Since(startTime)
	if err != nil {
		s.log.Warn(logrus.Fields{"elapsed_time": elapsedTime}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	s.log.Info(
		logrus.Fields{"elapsed_time": elapsedTime}, nil, "Executed: BankRepository.Register with no error",
	)

	// insert no rekening
	no_rekening, err = s.repository.RegisterRekening(tx, no_nasabah)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	s.repository.Commit(tx)

	return
}

func (s *BankService) Tabung(requestPayload data.TrxRequest) (saldo int, err error) {
	startTime := time.Now()
	// init transaction
	tx, err := s.repository.Begin()
	if err != nil {
		err = fmt.Errorf("failed to begin transaction")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	nominal := requestPayload.Nominal

	// check isRekeningValid by no_rekening
	isRekeningValid, err := s.repository.IsRekeningValid(tx, requestPayload.NoRekening)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	if !isRekeningValid {
		err = fmt.Errorf("INVALID")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return 0, err
	}

	// get saldo rekening
	saldoRekening, err := s.repository.GetSaldoByRekening(tx, requestPayload.NoRekening)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// update saldo rekening
	requestPayload.Nominal += saldoRekening
	err = s.repository.AddSaldoByRekening(tx, requestPayload)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// catat mutasi
	var transaksi data.Transaksi
	transaksi.Id = uuid.New()
	transaksi.NoRekening = requestPayload.NoRekening
	transaksi.KodeTransaksi = "D"
	transaksi.Nominal = nominal
	err = s.repository.AddMutasiTransaksi(tx, transaksi)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// get saldo terbaru
	saldo, err = s.repository.GetSaldoByRekening(tx, requestPayload.NoRekening)
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
		logrus.Fields{"elapsed_time": elapsedTime}, nil, "Executed: BankRepository.Tabung with no error",
	)

	// Commit
	s.repository.Commit(tx)

	return
}

func (s *BankService) Tarik(requestPayload data.TrxRequest) (saldo int, err error) {
	startTime := time.Now()
	// init transaction
	tx, err := s.repository.Begin()
	if err != nil {
		err = fmt.Errorf("failed to begin transaction")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	nominal := requestPayload.Nominal

	// check isRekeningValid by no_rekening
	isRekeningValid, err := s.repository.IsRekeningValid(tx, requestPayload.NoRekening)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}
	if !isRekeningValid {
		err = fmt.Errorf("INVALID")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return 0, err
	}

	// get saldo rekening
	saldoRekening, err := s.repository.GetSaldoByRekening(tx, requestPayload.NoRekening)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// update saldo rekening
	requestPayload.Nominal = saldoRekening - requestPayload.Nominal
	err = s.repository.SubstractSaldoByRekening(tx, requestPayload)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// catat mutasi
	var transaksi data.Transaksi
	transaksi.Id = uuid.New()
	transaksi.NoRekening = requestPayload.NoRekening
	transaksi.KodeTransaksi = "C"
	transaksi.Nominal = nominal
	err = s.repository.AddMutasiTransaksi(tx, transaksi)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repository.Rollback(tx)
		return
	}

	// get saldo terbaru
	saldo, err = s.repository.GetSaldoByRekening(tx, requestPayload.NoRekening)
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
		logrus.Fields{"elapsed_time": elapsedTime}, nil, "Executed: BankRepository.Tarik with no error",
	)

	// Commit
	s.repository.Commit(tx)

	return
}

func (s *BankService) GetSaldo(no_rekening string) (saldo int, err error) {
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
		return 0, err
	}

	// get saldo rekening
	saldo, err = s.repository.GetSaldoByRekening(tx, no_rekening)
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
		logrus.Fields{"elapsed_time": elapsedTime}, nil, "Executed: BankRepository.Tarik with no error",
	)

	// Commit
	s.repository.Commit(tx)

	return
}
