resource "form3_ace" "ace-edit" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "EDIT"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-edit-approve" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "EDIT_APPROVE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-delete" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "DELETE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-delete-approve" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "DELETE_APPROVE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-reject" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "REJECT"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-reject-approve" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "REJECT_APPROVE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

module "create-read-approve" {
  source          = "../create-read-approve"
  organisation_id = "${var.organisation_id}"
  role_id         = "${var.role_id}"
  records         = "${var.records}"
}
