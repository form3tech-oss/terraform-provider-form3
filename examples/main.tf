variable "api_host" {}
variable "client_id" {}
variable "client_secret" {}
variable "organisation_id" {}

provider "form3" {
  api_host      = "${var.api_host}"
  client_id     = "${var.client_id}"
  client_secret = "${var.client_secret}"
}

resource "form3_organisation" "oganisation" {
  organisation_id        = "2b7b602f-01ff-4845-892a-5a7c185867c6"
  parent_organisation_id = "${var.organisation_id}"
  name                   = "terraform-organisation"
}

resource "form3_user" "admin_user" {
  organisation_id = "${form3_organisation.oganisation.organisation_id}"
  user_id         = "56f36c4b-8df8-4577-8c4b-8d32f32210f0"
  user_name       = "terraform-user"
  email           = "terraform-user@form3.tech"
  roles           = ["${form3_role.role.role_id}"]
}

resource "form3_role" "role" {
  organisation_id = "${form3_organisation.oganisation.organisation_id}"
  role_id         = "81bc779a-620f-4e8c-9915-b8c6c90a5f17"
  name            = "sysadmin"
}

module "ace-all-actions" {
  source          = "ace/all"
  organisation_id = "${form3_organisation.oganisation.organisation_id}"
  records         = ["Organisation", "Role", "User", "Ace", "Account", "Payment", "PaymentSubmission", "Subscription", "FileEndpoint"]
  role_id         = "${form3_role.role.role_id}"
}

module "ace-create-read" {
  source          = "ace/create-read-approve"
  organisation_id = "${form3_organisation.oganisation.organisation_id}"
  records         = ["Return", "ReturnSubmission"]
  role_id         = "${form3_role.role.role_id}"
}

module "ace-read" {
  source          = "ace/read-approve"
  organisation_id = "${form3_organisation.oganisation.organisation_id}"
  records         = ["PaymentAdmission", "ReturnAdmission", "PaymentSubmissionValidation", "PaymentAdmissionValidation", "Reversal", "ReversalAdmission", "ReturnReversal", "ReturnReversalAdmission"]
  role_id         = "${form3_role.role.role_id}"
}

module "subscriptions-created" {
  source             = "subscriptions"
  organisation_id    = "${form3_organisation.oganisation.organisation_id}"
  callback_transport = "queue"
  callback_uri       = "https://sqs.eu-west-1.amazonaws.com/984234431138/terraform-test"
  event_type         = "Created"
  record_types       = ["PaymentAdmission", "ReturnAdmission", "ReversalAdmission"]
}

module "subscriptions-updated" {
  source             = "subscriptions"
  organisation_id    = "${form3_organisation.oganisation.organisation_id}"
  callback_transport = "queue"
  callback_uri       = "https://sqs.eu-west-1.amazonaws.com/984234431138/terraform-test"
  event_type         = "Updated"
  record_types       = ["PaymentSubmission", "ReturnSubmission"]
}
