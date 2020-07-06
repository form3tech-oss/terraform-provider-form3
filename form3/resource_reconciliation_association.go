package form3

import (
	"fmt"
	"log"
	"net/http"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3ReconciliationAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceReconciliationAssociationCreate,
		Read:   resourceReconciliationAssociationRead,
		Delete: resourceReconciliationAssociationDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bank_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
			"scheme_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceReconciliationAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createReconciliationNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Reconciliation association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostReconciliation(associations.NewPostReconciliationParams().
		WithCreationRequest(&models.ReconciliationAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create reconciliation_service association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] Reconciliation reconciliation_service association key: %s", d.Id())

	return resourceReconciliationAssociationRead(d, meta)
}

func resourceReconciliationAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	association, err := client.AssociationClient.Associations.GetReconciliationAssociationID(associations.NewGetReconciliationAssociationIDParams().
		WithAssociationID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, http.StatusNotFound) {
			return fmt.Errorf("couldn't find reconciliation_service association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	if err := d.Set("association_id", association.Payload.Data.ID.String()); err != nil {
		return err
	}
	if err := d.Set("name", association.Payload.Data.Attributes.Name); err != nil {
		return err
	}
	if err := d.Set("bank_ids", association.Payload.Data.Attributes.BankIds); err != nil {
		return err
	}
	if err := d.Set("scheme_type", association.Payload.Data.Attributes.SchemeType); err != nil {
		return err
	}

	return nil
}

func resourceReconciliationAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := client.AssociationClient.Associations.GetReconciliationAssociationID(associations.NewGetReconciliationAssociationIDParams().
		WithAssociationID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting reconciliation_service association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteReconciliationAssociationID(associations.NewDeleteReconciliationAssociationIDParams().
		WithAssociationID(association.Payload.Data.ID).
		WithVersion(*association.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting reconciliation_service association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createReconciliationNewAssociationFromResourceData(d *schema.ResourceData) (*models.ReconciliationNewAssociation, error) {
	association := models.ReconciliationNewAssociation{
		Type:       models.ReconciliationAssociationTypeReconciliationAssociations,
		Attributes: &models.ReconciliationAssociationAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
	}

	if attr, ok := d.GetOk("bank_ids"); ok {
		arr := attr.([]interface{})
		for _, v := range arr {
			association.Attributes.BankIds = append(association.Attributes.BankIds, v.(string))
		}
	}

	if attr, ok := d.GetOk("scheme_type"); ok {
		association.Attributes.SchemeType = attr.(string)
	}

	return &association, nil
}
