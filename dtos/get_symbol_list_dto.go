package dtos

type GetSymbolList struct {
	Id string `param:"id"`
}

func NewGetSymbolList() *GetSymbolList {
	return &GetSymbolList{}
}

func CreateGetSymbolList() GetSymbolList {
	return GetSymbolList{}
}
