package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceForm3SepaLiquidityAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaLiquidityAssociationCreate,
		Read:   resourceSepaLiquidityAssociationRead,
		Delete: resourceSepaLiquidityAssociationDelete,

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
			"technical_bic": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"settlement_bic": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"settlement_iban": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sponsored_bics": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
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
			"sponsor_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
		},
	}
}

func resourceSepaLiquidityAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaLiquidityNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa liquidity association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepaLiquidity(associations.NewPostSepaLiquidityParams().
		WithCreationRequest(&models.SepaLiquidityAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa liquidity association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa liquidity assocation key: %s", d.Id())

	return resourceSepaLiquidityAssociationRead(d, meta)
}

func resourceSepaLiquidityAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaLiquidityAssociation, err := client.AssociationClient.Associations.GetSepaLiquidityID(associations.NewGetSepaLiquidityIDParams().
		WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find sepa liquidity assocation: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("association_id", sepaLiquidityAssociation.Payload.Data.ID.String())
	_ = d.Set("name", sepaLiquidityAssociation.Payload.Data.Attributes.Name)
	_ = d.Set("technical_bic", sepaLiquidityAssociation.Payload.Data.Attributes.TechnicalBic)
	_ = d.Set("settlement_bic", sepaLiquidityAssociation.Payload.Data.Attributes.SettlementBic)
	_ = d.Set("settlement_iban", sepaLiquidityAssociation.Payload.Data.Attributes.SettlementIban)
	_ = d.Set("sponsored_bics", sepaLiquidityAssociation.Payload.Data.Attributes.SponsoredBics)
	_ = d.Set("address_street", sepaLiquidityAssociation.Payload.Data.Attributes.Address.Street)
	_ = d.Set("address_building_number", sepaLiquidityAssociation.Payload.Data.Attributes.Address.BuildingNumber)
	_ = d.Set("address_city", sepaLiquidityAssociation.Payload.Data.Attributes.Address.City)
	_ = d.Set("address_country", sepaLiquidityAssociation.Payload.Data.Attributes.Address.Country)
	if sepaLiquidityAssociation.Payload.Data.Relationships == nil {
		d.Set("sponsor_id", "")
	} else {
		d.Set("sponsor_id", sepaLiquidityAssociation.Payload.Data.Relationships.Sponsor.Data.ID.String())
	}
	return nil
}

func resourceSepaLiquidityAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	SepaLiquidityAssociation, err := client.AssociationClient.Associations.GetSepaLiquidityID(associations.NewGetSepaLiquidityIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa liquidity assocation: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteSepaLiquidityID(associations.NewDeleteSepaLiquidityIDParams().
		WithID(SepaLiquidityAssociation.Payload.Data.ID).
		WithVersion(*SepaLiquidityAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa liquidity assocation: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createSepaLiquidityNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaLiquidityAssociation, error) {

	association := models.NewSepaLiquidityAssociation{Attributes: &models.SepaLiquidityAssociationAttributes{}}
	association.Type = "sepaliquidity_associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
	}

	if attr, ok := d.GetOk("technical_bic"); ok {
		association.Attributes.TechnicalBic = attr.(string)
	}

	if attr, ok := d.GetOk("settlement_bic"); ok {
		association.Attributes.SettlementBic = attr.(string)
	}

	if attr, ok := d.GetOk("settlement_iban"); ok {
		association.Attributes.SettlementIban = attr.(string)
	}

	if attr, ok := d.GetOk("sponsored_bics"); ok {
		arr := attr.([]interface{})
		var ibans []string
		for _, v := range arr {
			ibans = append(ibans, v.(string))
		}
		association.Attributes.SponsoredBics = ibans
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

	if attr, ok := GetUUIDOK(d, "sponsor_id"); ok {
		association.Relationships = &models.SepaLiquidityAssociationRelationships{
			Sponsor: models.SepaLiquidityAssociationRelationshipsSponsor{
				Data: models.SepaLiquidityRelationshipData{
					ID:   attr,
					Type: "sepaliquidity_associations",
				},
			},
		}
	}
	return &association, nil
}
