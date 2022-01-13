package gqlresolver

import (
	"context"

	"github.com/ipfans/gqlgen-error/business/entity"
	"github.com/ipfans/gqlgen-error/business/usecase/gqlresolver/gqlgen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Queries() gqlgen.QueriesResolver {
	return &queriesResolver{r}
}

type queriesResolver struct{ *Resolver }

func (r *queriesResolver) City(ctx context.Context, id int) (*entity.City, error) {
	panic("not implemented")
}
func (r *queriesResolver) Cities(ctx context.Context, p *entity.ParamCityInput) (*entity.CityConnection, error) {
	panic("not implemented")
}
