variable "kubernetes_host" {}

variable "kubernetes_token" {}

variable "cluster_ca_certificate" {}

variable "http_port" {
    description = "http port for exposing the blog server"

    default = "8088"
}

variable "node_port" {
    description = "http port for exposing the blog service"

    default = "30011"
}

variable "blog_image" {
    description = "docker image for the blog container"
}

variable "broadcast_image" {
    description = "docker image for the broadcast cronjob"
}

variable "linkedin_access_token" {}

variable "twitter_consumer_key" {}

variable "twitter_consumer_secret" {}

variable "twitter_token" {}

variable "twitter_token_secret" {}

variable "pocket_consumer_key" {}

variable "pocket_access_token" {}

variable "list_image" {}

variable "google_credentials_file" {}

variable "google_token_file" {}

variable "aws_access_key_id" {}

variable "aws_secret_access_key" {}

variable "list_http_port" {
    description = "http port for exposing the list service"

    default = "8099"
}