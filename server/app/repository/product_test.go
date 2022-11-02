package repository

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sol-shop-server/ent"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var repo *ProductRepository

func TestMain(m *testing.M) {
	client, _ := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client.Schema.Create(context.Background())
	defer client.Close()

	r := NewRepository(client)
	repo = r.ProductRepository

	os.Exit(m.Run())
}

var dummyProduct = ent.Product{
	Name:        "Water",
	Description: "To Drink",
	Price:       500,
	Stock:       10,
	Edges:       ent.ProductEdges{},
}

func TestProductService_GetProductByID(t *testing.T) {
	created, err := repo.CreateProduct(&dummyProduct, nil)
	require.NoError(t, err)

	got, err := repo.GetProductByID(1)
	require.NoError(t, err)

	checkProduct(t, created, got)
}

func TestProductService_ListProducts(t *testing.T) {
	var err error
	created := make([]*ent.Product, 2)
	created[0], err = repo.CreateProduct(&dummyProduct, nil)
	require.NoError(t, err)
	created[1], err = repo.CreateProduct(&dummyProduct, nil)
	require.NoError(t, err)

	query, err := repo.ListProducts(2, 0)
	require.NoError(t, err)

	for i, got := range query {
		checkProduct(t, created[i], got)
	}
}

func TestProductService_ListProductsByCategoryID(t *testing.T) {
	var err error
	category, err := repo.CreateProductCategory(newRandomProductCategory())
	require.NoError(t, err)

	created := make([]*ent.Product, 2)
	created[0], err = repo.CreateProduct(&dummyProduct, []int{category.ID})
	require.NoError(t, err)
	created[1], err = repo.CreateProduct(&dummyProduct, []int{category.ID})
	require.NoError(t, err)

	query, err := repo.ListProductsByCategoryID(category.ID, 2, 0)
	require.NoError(t, err)

	for i, got := range query {
		checkProduct(t, created[i], got)
	}
}

func TestProductService_CreateProduct(t *testing.T) {
	got, err := repo.CreateProduct(&dummyProduct, nil)
	require.NoError(t, err)

	checkProduct(t, &dummyProduct, got)
}

func TestProductService_UpdateProduct(t *testing.T) {
	created, err := repo.CreateProduct(&dummyProduct, nil)
	require.NoError(t, err)

	created.Description = "New Description"
	got, err := repo.UpdateProduct(created.ID, created, nil)
	require.NoError(t, err)

	checkProduct(t, created, got)
}

func TestProductService_DeleteProduct(t *testing.T) {
	created, err := repo.CreateProduct(&dummyProduct, nil)
	require.NoError(t, err)

	repo.DeleteProduct(created.ID)
	_, err = repo.GetProductByID(created.ID)
	require.Error(t, err)
}

func TestProductService_GetProductCategoryByID(t *testing.T) {
	created, err := repo.CreateProductCategory(newRandomProductCategory())
	require.NoError(t, err)

	got, err := repo.GetProductCategoryByID(created.ID)
	require.NoError(t, err)

	checkProductCategory(t, created, got)
}

func TestProductService_ListProductCategories(t *testing.T) {
	var err error
	created := make([]*ent.ProductCategory, 2)
	created[0], err = repo.CreateProductCategory(newRandomProductCategory())
	require.NoError(t, err)
	created[1], err = repo.CreateProductCategory(newRandomProductCategory())
	require.NoError(t, err)

	query, err := repo.ListProductCategories(2, 0)
	require.NoError(t, err)
	require.Len(t, query, 2)

	query, err = repo.ListProductCategories(1, 0)
	require.NoError(t, err)
	require.Len(t, query, 1)
}

func TestProductService_CreateProductCategory(t *testing.T) {
	gen := newRandomProductCategory()
	got, err := repo.CreateProductCategory(gen)
	require.NoError(t, err)

	checkProductCategory(t, gen, got)
}

func TestProductService_UpdateProductCategory(t *testing.T) {
	created, err := repo.CreateProductCategory(newRandomProductCategory())
	require.NoError(t, err)

	created.Description = "New Description"
	got, err := repo.UpdateProductCategory(created.ID, created)
	require.NoError(t, err)

	checkProductCategory(t, created, got)
}

func TestProductService_DeleteProductCategory(t *testing.T) {
	got, err := repo.CreateProductCategory(newRandomProductCategory())
	require.NoError(t, err)

	repo.DeleteProductCategory(got.ID)

	_, err = repo.GetProductCategoryByID(got.ID)
	require.Error(t, err)
}

func checkProduct(t *testing.T, want *ent.Product, got *ent.Product) {
	// require.Equal(t, dummyProduct.ID, got.ID)
	require.Equal(t, want.Name, got.Name)
	require.Equal(t, want.Description, got.Description)
	require.Equal(t, want.Price, got.Price)
	require.Equal(t, want.Stock, got.Stock)
}

func checkProductCategory(t *testing.T, want *ent.ProductCategory, got *ent.ProductCategory) {
	// require.Equal(t, dummyProduct.ID, got.ID)
	require.Equal(t, want.Name, got.Name)
	require.Equal(t, want.Description, got.Description)
}

func newRandomProductCategory() *ent.ProductCategory {
	return &ent.ProductCategory{
		Name:        randomString(5),
		Description: randomString(10),
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
