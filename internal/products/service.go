package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	//repo repo.Querier
}

// func NewService(repo repo.Querier) Service {
// 	return &svc{repo: repo}
// }


func NewService() Service {
	return &svc{}
}

// func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
// 	return s.repo.ListProducts(ctx)
// }

func (s *svc) ListProducts(ctx context.Context) ( error) {
	return nil
}

