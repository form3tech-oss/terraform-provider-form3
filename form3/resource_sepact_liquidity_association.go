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

func resourceForm3SepactLiquidityAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepactLiquidityAssociationCreate,
		Read:   resourceSepactLiquidityAssociationRead,
		Delete: resourceSepactLiquidityAssociationDelete,
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
			"reachable_bics": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ForceNew: true,
			},
			"settlement_bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"settlement_iban": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_street": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_building_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_city": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"direct_participant_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
		},
	}
}

func resourceSepactLiquidityAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepactLiquidityNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepact-liquidity association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepactLiquidity(associations.NewPostSepactLiquidityParams().
		WithCreationRequest(&models.SepactLiquidityAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepact-liquidity association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepact-liquidity association key: %s", d.Id())

	return resourceSepactLiquidityAssociationRead(d, meta)
}

func resourceSepactLiquidityAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	association, err := client.AssociationClient.Associations.GetSepactLiquidityAssociationID(associations.NewGetSepactLiquidityAssociationIDParams().
		WithAssociationID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, http.StatusNotFound) {
			return fmt.Errorf("couldn't find sepact-liquidity association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	if err := d.Set("association_id", association.Payload.Data.ID.String()); err != nil {
		return err
	}
	if err := d.Set("organisation_id", association.Payload.Data.OrganisationID.String()); err != nil {
		return err
	}
	if err := d.Set("name", association.Payload.Data.Attributes.Name); err != nil {
		return err
	}
	if err := d.Set("reachable_bics", association.Payload.Data.Attributes.ReachableBics); err != nil {
		return err
	}
	if err := d.Set("settlement_bic", association.Payload.Data.Attributes.SettlementBic); err != nil {
		return err
	}
	if err := d.Set("settlement_iban", association.Payload.Data.Attributes.SettlementIban); err != nil {
		return err
	}
	if err := d.Set("address_street", association.Payload.Data.Attributes.Address.Street); err != nil {
		return err
	}
	if err := d.Set("address_building_number", association.Payload.Data.Attributes.Address.BuildingNumber); err != nil {
		return err
	}
	if err := d.Set("address_city", association.Payload.Data.Attributes.Address.City); err != nil {
		return err
	}
	if err := d.Set("address_country", association.Payload.Data.Attributes.Address.Country); err != nil {
		return err
	}
	if association.Payload.Data.Relationships != nil {
		if err := d.Set("direct_participant_id", association.Payload.Data.Relationships.DirectParticipant.Data.ID.String()); err != nil {
			return err
		}
	}

	return nil
}

func resourceSepactLiquidityAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := client.AssociationClient.Associations.GetSepactLiquidityAssociationID(associations.NewGetSepactLiquidityAssociationIDParams().
		WithAssociationID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepact-liquidity association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteSepactLiquidityAssociationID(associations.NewDeleteSepactLiquidityAssociationIDParams().
		WithAssociationID(association.Payload.Data.ID).
		WithVersion(*association.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepact-liquidity association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createSepactLiquidityNewAssociationFromResourceData(d *schema.ResourceData) (*models.SepactLiquidityNewAssociation, error) {
	association := models.SepactLiquidityNewAssociation{
		Type:       models.SepactLiquidityAssociationTypeSepactLiquidityAssociations,
		Attributes: &models.SepactLiquidityAssociationAttributes{},
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

	if attr, ok := d.GetOk("reachable_bics"); ok {
		arr := attr.(*schema.Set).List()
		for _, v := range arr {
			association.Attributes.ReachableBics = append(association.Attributes.ReachableBics, models.Bic11(v.(string)))
		}
	}

	if attr, ok := d.GetOk("settlement_bic"); ok {
		association.Attributes.SettlementBic = models.Bic8(attr.(string))
	}

	if attr, ok := d.GetOk("settlement_iban"); ok {
		association.Attributes.SettlementIban = attr.(string)
	}

	if attr, ok := d.GetOk("address_street"); ok {
		association.Attributes.Address.Street = attr.(string)
	}

	if attr, ok := d.GetOk("address_building_number"); ok {
		association.Attributes.Address.BuildingNumber = attr.(string)
	}

	if attr, ok := d.GetOk("address_city"); ok {
		association.Attributes.Address.City = attr.(string)
	}

	if attr, ok := d.GetOk("address_country"); ok {
		association.Attributes.Address.Country = attr.(string)
	}

	if attr, ok := GetUUIDOK(d, "direct_participant_id"); ok {
		association.Relationships = &models.SepactLiquidityAssociationRelationships{
			DirectParticipant: models.SepactLiquidityAssociationRelationshipsDirectParticipant{
				Data: models.SepactLiquidityRelationshipData{
					ID:   attr,
					Type: models.SepactLiquidityAssociationTypeSepactLiquidityAssociations,
				},
			},
		}
	}

	return &association, nil
}
