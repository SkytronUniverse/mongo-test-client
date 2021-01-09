package products

import (
	config "dev/interview-craft/configs"
	"dev/interview-craft/tools"

	"github.com/labstack/gommon/log"
)

type ProductManager interface {
	AddNewProduct(product ProductDetails) error
	GetInventory() ([]ProductDetails, error)
}

type ProductManagerDriver struct {
	MongoClient  tools.MongoClient
	MongoDetails config.Details
}

type ProductDetails struct {
	ID    string `json:"id,omitempty" bson:"id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Price string `json:"price,omitempty" bson:"price,omitempty"`
}

type ProductDetailsList []ProductDetails

func NewProductManager(mgoClient tools.MongoClient, mgoDetails config.Details) ProductManager {
	return &ProductManagerDriver{
		MongoClient:  mgoClient,
		MongoDetails: mgoDetails,
	}
}

func (p *ProductManagerDriver) AddNewProduct(product ProductDetails) error {
	log.Info("products-add-new-product")
	defer log.Info("products-add-new-product-complete")

	_, err := p.MongoClient.AddProductToInventory(p.MongoDetails.Name, p.MongoDetails.Collection, product)
	if err != nil {
		log.Error("products-add-new-product-error ", err)
		return err
	}

	return nil
}

func (p *ProductManagerDriver) GetInventory() ([]ProductDetails, error) {
	log.Info("products-get-inventory")
	defer log.Info("products-get-inventory-complete")

	var productList ProductDetailsList

	err := p.MongoClient.DisplayInventory(p.MongoDetails.Name, p.MongoDetails.Collection, &productList)
	if err != nil {
		log.Error("products-get-inventory-error ", err)
		return nil, err
	}

	return productList, nil
}
