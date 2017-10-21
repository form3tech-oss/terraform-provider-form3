resource "form3_ace" "ace-read" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "READ"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}

resource "form3_ace" "ace-read-approve" {
  ace_id          = "${uuid()}"
  role_id         = "${var.role_id}"
  organisation_id = "${var.organisation_id}"
  record_type     = "${element(var.records, count.index)}"
  action          = "APPROVE"
  count           = "${length(var.records)}"
  lifecycle {
    ignore_changes = ["ace_id"]
  }
}
