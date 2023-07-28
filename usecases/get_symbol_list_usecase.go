package usecases

import (
	"errors"

	"github.com/LucasGabrielBravo/primo-block-store/dtos"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

type GetSymbolList struct {
	app *pocketbase.PocketBase
}

func NewGetSymbolList(app *pocketbase.PocketBase) *GetSymbolList {
	return &GetSymbolList{
		app: app,
	}
}

func CreateGetSymbolList(app *pocketbase.PocketBase) GetSymbolList {
	return GetSymbolList{
		app: app,
	}
}

func (g GetSymbolList) Execute(dto dtos.GetSymbolList) (map[string]any, error) {
	result := map[string]any{}
	symbols := []types.JsonRaw{}

	record, err := g.app.Dao().FindRecordById("lists", dto.Id)

	if err != nil {
		return nil, errors.New("list not found")
	}

	g.app.Dao().ExpandRecord(record, []string{"blocks"}, func(relCollection *models.Collection, relIds []string) ([]*models.Record, error) {
		return g.app.Dao().FindRecordsByIds(relCollection.Id, relIds)
	})

	for _, block := range record.Expand()["blocks"].([]*models.Record) {
		symbols = append(symbols, block.Get("symbol").(types.JsonRaw))
	}

	result["symbols"] = symbols

	return result, nil
}
