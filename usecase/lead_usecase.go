package usecase

import (
	"gin-gonic-leads-api/model"
	"gin-gonic-leads-api/repository"
)

type LeadUsecase struct {
	repository repository.LeadRepository
}

func NewLeadUseCase(repo repository.LeadRepository) LeadUsecase {
	return LeadUsecase{
		repository: repo,
	}
}

func (pu *LeadUsecase) Create(lead model.Lead) (model.Lead, error) {

	leadId, err := pu.repository.Create(lead)
	if err != nil {
		return model.Lead{}, err
	}

	lead.Id = leadId

	return lead, nil
}
