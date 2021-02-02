
locals {
  organisation_id          = "%s"
  parent_organisation_id   = "%s"
  association_id           = "%s"
  sponsored_association_id = "%s"
}
resource "form3_organisation" "organisation" {
  organisation_id        = "${local.organisation_id}"
  parent_organisation_id = "${local.parent_organisation_id}"
  name                   = "terraform-provider-form3-test-organisation"
}

resource "form3_sepainstant_association" "association" {
  organisation_id      = "${form3_organisation.organisation.organisation_id}"
  association_id       = "${local.association_id}"
  business_user_dn     = "cn=testbic8,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu"
  transport_profile_id = "TEST_PROFILE_1"
  bic                  = "TESTBIC8"
  simulator_only       = true
}

resource "form3_sepainstant_association" "association_sponsored" {
  organisation_id      = "${form3_organisation.organisation.organisation_id}"
  association_id       = "${local.sponsored_association_id}"
  business_user_dn     = ""
  transport_profile_id = ""
  bic                  = "TESTBIC8"
  simulator_only       = true
  sponsor_id           = "${form3_sepainstant_association.association.association_id}"

  depends_on = [
    "form3_sepainstant_association.association"
  ]
}
