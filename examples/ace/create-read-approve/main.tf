resource "form3_ace" "ace-create" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "CREATE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-create-approve" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "CREATE_APPROVE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

module "read-approve" {
  source          = "../read-approve"
  organisation_id = "${var.organisation_id}"
  role_id         = "${var.role_id}"
  records         = "${var.records}"
}
