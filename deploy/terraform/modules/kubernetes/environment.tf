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

variable "certificate_id" {}