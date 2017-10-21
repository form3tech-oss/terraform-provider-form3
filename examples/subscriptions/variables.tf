variable "organisation_id" {
  type        = "string"
  description = "organisation to create aces for"
}

variable "record_types" {
  type        = "list"
  description = "list of records to create subscriptions for"
}

variable "event_type" {
  type        = "string"
  description = "event type to create subscription for"
}

variable "callback_transport" {
  type        = "string"
  description = "either queue or http"
}

variable "callback_uri" {
  type        = "string"
  description = "uri of the callback"
}
