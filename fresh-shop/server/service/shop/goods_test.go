package shop

import (
	"mime/multipart"
	"testing"
)

func TestGoodsService_BatchCreateGoodsByExcel(t *testing.T) {
	g := GoodsService{}
	h := multipart.FileHeader{}
	err := g.BatchCreateGoodsByExcel(&h)
	if err != nil {
		panic(err)
	}
}
