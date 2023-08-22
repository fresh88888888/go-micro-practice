package service

import "strconv"

type ProdModel struct {
	ProdID   int    `json:"pid"`
	ProdName string `json:"prod_name"`
}

type ProdRequest struct {
	Size int `form:"size"`
}

func NewProd(id int, name string) *ProdModel {
	return &ProdModel{}
}

func NewProdList(n int) []*ProdModel {
	rets := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		rets = append(rets, &ProdModel{ProdID: i + 100, ProdName: "pro-" + strconv.Itoa(i)})
	}
	return rets
}
