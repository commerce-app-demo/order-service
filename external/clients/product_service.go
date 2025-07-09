package clients

import (
	"context"

	productspb "github.com/commerce-app-demo/product-service/proto"
)

type ProductClient struct {
	pbClient *productspb.ProductServiceClient
	ctx      context.Context
}

func (cl *ProductClient) ValidateProduct(id string) (*productspb.Product, error) {
	product, err := (*cl.pbClient).GetProduct(cl.ctx, &productspb.GetProductRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func NewProductClient(ctx context.Context, pb *productspb.ProductServiceClient) (*ProductClient, error) {
	return &ProductClient{
		pbClient: pb,
		ctx:      ctx,
	}, nil
}
