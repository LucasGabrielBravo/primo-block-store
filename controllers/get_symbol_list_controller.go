package controllers

import (
	"net/http"

	"github.com/LucasGabrielBravo/primo-block-store/dtos"
	"github.com/LucasGabrielBravo/primo-block-store/usecases"
	"github.com/labstack/echo/v5"
)

type GetSymbolList struct {
	getSymbolListUsecase *usecases.GetSymbolList
}

func NewGetSymbolList(getSymbolListUsecase *usecases.GetSymbolList) *GetSymbolList {
	return &GetSymbolList{
		getSymbolListUsecase: getSymbolListUsecase,
	}
}

func CreateGetSymbolList(getSymbolListUsecase usecases.GetSymbolList) GetSymbolList {
	return GetSymbolList{
		getSymbolListUsecase: &getSymbolListUsecase,
	}
}

func (g GetSymbolList) Handler(ctx echo.Context) error {
	dto := dtos.CreateGetSymbolList()

	err := ctx.Bind(&dto)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}

	symbols, err := g.getSymbolListUsecase.Execute(dto)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, symbols)
}
