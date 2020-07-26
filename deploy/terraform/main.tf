provider "kubernetes" {
    # Specify the path to the kubernetes config file
    config_path = "/Users/mihaiblebea/.kube/config"
}

variable "http_port" {
    description = "http port for exposing the blog server"
    type = string

    default = "8088"
}

resource "kubernetes_namespace" "mihaiblebea" {
    metadata {
        name = "mihaiblebea"
    }
}

resource "kubernetes_deployment" "blog-deployment" {
    metadata {
        name = "blog-deployment"
        namespace = "mihaiblebea"
        labels = {
            app = "go-broadcast"
            name = "blog-deployment"
        }
    }

    spec {
        replicas = 1

        selector {
            match_labels = {
                app = "go-broadcast"
                name = "blog-pod"
            }
        }

        template {
            metadata {
                name = "blog-pod"
                labels = {
                    app = "go-broadcast"
                    name = "blog-pod"
                }
            }

            spec {
                container {
                    image = "serbanblebea/go-blog:v0.1"
                    name = "blog-container"
                    env {
                        name = "HTTP_PORT"
                        value = var.http_port
                    }
                    port {
                        container_port = var.http_port
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "blog-service" {
    metadata {
        name = "blog-service"
        namespace = "mihaiblebea"
    }

    spec {
        selector = {
            name = "blog-pod"
        }

        port {
            port = 80
            target_port = var.http_port
            node_port = 30011
        }

        type = "NodePort"
    }
}

output "browser_base_url" {
    value = "${kubernetes_service.blog-service}"
}