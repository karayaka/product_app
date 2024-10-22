package infrastructure

import (
	"context"
	"fmt"
	"os"
	"product_app/common/postgresql"
	"product_app/domain"
	"product_app/persistence"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var productRepository persistence.IProductRepository

var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Before all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func TestGetAllProducts(t *testing.T) {
	//setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()

		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	//clear(ctx, dbPool)
}

func TestAddProducts(t *testing.T) {
	//setup(ctx, dbPool)

	product := domain.Product{
		Id:       1,
		Name:     "AirFryer",
		Price:    3000.0,
		Discount: 22.0,
		Store:    "ABC TECH",
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		err := productRepository.AddProduct(product)
		if err != nil {
			log.Error("hata")
		}
	})

	//clear(ctx, dbPool)
}

func TestGetById(t *testing.T) {
	//setup(ctx, dbPool)

	t.Run("GetAllProducts", func(t *testing.T) {
		product, err := productRepository.GetById(1)
		if err != nil {
			log.Error("hata")
		}
		fmt.Println(product)
	})

	//clear(ctx, dbPool)
}
