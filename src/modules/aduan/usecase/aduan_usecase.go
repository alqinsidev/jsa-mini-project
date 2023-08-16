package usecase

import (
	"alqinsidev/jsa-mini-project/aduan/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type aduanUsecase struct {
	aduanRepo domain.AduanRepository
}

func NewAduanUsecase(n domain.AduanRepository) domain.AduanUsecase {
	return &aduanUsecase{
		aduanRepo: n,
	}
}

func (n *aduanUsecase) Fetch(query *domain.RequestQuery) (res []domain.AduanTableResponse, totalData int64, err error) {
	res, totalData, err = n.aduanRepo.Fetch(query)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (n *aduanUsecase) FindById(id uuid.UUID) (res *domain.AduanDetail, err error) {
	res, err = n.aduanRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return
}

func (n *aduanUsecase) FetchSummary() (res *domain.AduanSummaryResponse, err error) {
	res, err = n.aduanRepo.FetchSummary()
	if err != nil {
		log.Error().Err(err).Msg("err summary usecase")
		return nil, err
	}
	return
}

func (n *aduanUsecase) UpdateStatus(p *domain.UpdateStatusPayload) (res interface{}, err error) {
	res, err = n.aduanRepo.UpdateStatus(p)
	if err != nil {
		return nil, err
	}
	return
}
