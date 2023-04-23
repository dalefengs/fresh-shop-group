package shop

type ServiceGroup struct {
	CategoryService
	BrandService
	BrandCategoryService
	TagsService
	GoodsService
	OrderService
	OrderDetailsService
	OrderDeliveryService
	OrderReturnService
	FavoritesService
}
