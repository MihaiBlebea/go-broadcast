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

    default = "serbanblebea/go-blog:v0.2"
}