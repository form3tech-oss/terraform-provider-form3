resource "form3_subscription" "subscription" {
  organisation_id    = "${var.organisation_id}"
  subscription_id    = "${uuid()}"
  callback_transport = "${var.callback_transport}"
  callback_uri       = "${var.callback_uri}"
  event_type         = "${var.event_type}"
  record_type        = "${element(var.record_types, count.index)}"
  count              = "${length(var.record_types)}"
  lifecycle {
    ignore_changes = ["subscription_id"]
  }
}
