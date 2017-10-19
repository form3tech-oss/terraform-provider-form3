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
  user_id = "56f36c4b-8df8-4577-8c4b-8d32f32210f0"
  user_name = "terraform-user"
  email = "terraform-user@form3.tech"
  roles = ["${form3_role.role.role_id}"]
}

resource "form3_role" "role" {
  organisation_id = "${var.organisation_id}"
  role_id = "81bc779a-620f-4e8c-9915-b8c6c90a5f17"
  name = "sysadmin"
}
