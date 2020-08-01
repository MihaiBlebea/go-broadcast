variable "kubernetes_host" {}

variable "kubernetes_token" {}

variable "cluster_ca_certificate" {}

variable "http_port" {
    description = "http port for exposing the blog server"
    type = string

    default = "8088"
}

variable "node_port" {
    description = "http port for exposing the blog service"
    type = string

    default = "30011"
}

variable "blog_image" {
    description = "docker image for the blog container"
    type = string
}