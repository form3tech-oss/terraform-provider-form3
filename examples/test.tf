variable "api_host" {}
variable "client_id" {}
variable "client_secret" {}
variable "organisation_id" {}

provider "form3" {
  api_host      = "${var.api_host}"
  client_id     = "${var.client_id}"
  client_secret = "${var.client_secret}"
}


resource "form3_user" "admin_user" {
  organisation_id = "${var.organisation_id}"
  user_id = "44247ebb-fe01-44ab-887d-7f344481712f"
  user_name = "terraform-user"
  email = "terraform-user@form3.tech"
}