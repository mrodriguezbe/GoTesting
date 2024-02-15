package repository_test

import (
	"gotesting/Desafio-Cierre-Testing/internal"
	"gotesting/Desafio-Cierre-Testing/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchProduct(t *testing.T) {
	t.Run("search product should return a list with the matching id product", func(t *testing.T) {
		// arrange
		// products
		products := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 2",
					Price:       200,
					SellerId:    2,
				},
			},
		}
		// query
		query := internal.ProductQuery{
			Id: 1,
		}
		// repository
		rp := repository.NewProductsMap(products)
		// expected result
		expectedResult := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
		}

		// act
		result, err := rp.SearchProducts(query)

		// assert
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("search product should return empty list", func(t *testing.T) {
		// arrange
		// products
		products := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 2",
					Price:       200,
					SellerId:    2,
				},
			},
		}
		// query
		query := internal.ProductQuery{
			Id: 47,
		}
		// repository
		rp := repository.NewProductsMap(products)
		// expected result
		expectedResult := map[int]internal.Product{}

		// act
		result, err := rp.SearchProducts(query)

		// assert
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})
}
