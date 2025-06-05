package main

import (
	"fmt"
	"go.uber.org/fx"
)

type Cache struct {
	entity string
}

func NewItemCache() Cache {
	return Cache{entity: "item"}
}

func NewAddressCache() Cache {
	return Cache{entity: "address"}
}

type ItemRepository struct {
	itemCache Cache
}

func NewItemRepository(itemCache Cache) *ItemRepository {
	return &ItemRepository{itemCache: itemCache}
}

func main() {
	fx.New(
		fx.Provide(fx.Annotate(NewItemCache, fx.ResultTags(`name:"item_cache"`))),
		fx.Provide(fx.Annotate(NewAddressCache, fx.ResultTags(`name:"address_cache"`))),
		fx.Provide(fx.Annotate(NewItemRepository, fx.ParamTags(`name:"address_cache"`))),
		fx.Invoke(func(repository *ItemRepository) { fmt.Printf("repository cache: %s", repository.itemCache.entity) }),
	).Run()
}
