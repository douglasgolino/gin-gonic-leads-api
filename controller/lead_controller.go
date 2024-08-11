package controller

import (
	"gin-gonic-leads-api/model"
	"gin-gonic-leads-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeadController struct {
	LeadUsecase usecase.LeadUsecase
}

func NewLeadController(usecase usecase.LeadUsecase) LeadController {
	return LeadController{
		LeadUsecase: usecase,
	}
}

func (p *LeadController) Create(ctx *gin.Context) {

	var lead model.Lead

	err := ctx.BindJSON(&lead)
	if err != nil {
		response := model.Response{
			Message: "Solicitação Inválida " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	insertedLead, err := p.LeadUsecase.Create(lead)

	if err != nil {
		response := model.Response{
			Message: "Erro ao processar a solicitação",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, insertedLead)
}
