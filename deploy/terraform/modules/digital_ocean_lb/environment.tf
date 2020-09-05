variable "do_token" {
    description = "digital ocean api token"
    type        = string
}

variable "google_search_console_code" {
    description = "code copied from google search console for domain validation"
}

variable "droplet_id" {}

variable "domain_name" {}

variable "aws_domain_verification_token" {}

variable "aws_domain_key_set" {
    type = list(string)
}