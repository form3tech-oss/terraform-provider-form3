package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceForm3ProductsAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceProductsAssociationCreate,
		Read:   resourceProductsAssociationRead,
		Delete: resourceProductsAssociationDelete,

		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceProductsAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createProductsNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create product association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostProducts(associations.NewPostProductsParams().
		WithCreationRequest(&models.ProductsAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create product association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] product association key: %s", d.Id())

	return resourceProductsAssociationRead(d, meta)
}

func resourceProductsAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	productAssociation, err := client.AssociationClient.Associations.GetProductsID(associations.NewGetProductsIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("couldn't find product association: %s", err)
	}

	_ = d.Set("association_id", productAssociation.Payload.Data.ID.String())
	_ = d.Set("product", productAssociation.Payload.Data.Attributes.Product)

	return nil
}

func resourceProductsAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	productAssociation, err := client.AssociationClient.Associations.GetProductsID(associations.NewGetProductsIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting product association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteProductsID(associations.NewDeleteProductsIDParams().
		WithID(productAssociation.Payload.Data.ID))

	if err != nil {
		return fmt.Errorf("error deleting product association: %s", err)
	}

	return nil
}

func createProductsNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewProductsAssociation, error) {

	association := models.NewProductsAssociation{Attributes: &models.ProductsAssociationAttributes{}}
	association.Type = "product_associations"

	if attr, ok := d.GetOk("organisation_id"); ok {
		association.OrganisationID = strfmt.UUID(attr.(string))
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := d.GetOk("product"); ok {
		association.Attributes.Product = attr.(string)
	}

	return &association, nil
}
