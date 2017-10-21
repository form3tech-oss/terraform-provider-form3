variable "records" {
  type        = "list"
  description = "list of records to create aces for"
}

variable "organisation_id" {
  type        = "string"
  description = "organisation to create aces for"
}

variable "role_id" {
  type        = "string"
  description = "associated role to create aces for"
}
