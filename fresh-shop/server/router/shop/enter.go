package shop

type RouterGroup struct {
	CategoryRouter
	BrandRouter
	BrandCategoryRouter
	TagsRouter
	GoodsRouter
	OrderRouter
	OrderDetailsRouter
	OrderDeliveryRouter
	OrderReturnRouter
	FavoritesRouter
	CartRouter
	UserAddressRouter
}
