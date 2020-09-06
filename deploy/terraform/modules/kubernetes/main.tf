provider "kubernetes" {
    # config_path = "/Users/mihaiblebea/.kube/config"
    load_config_file       = false
    host                   = var.kubernetes_host
    token                  = var.kubernetes_token
    cluster_ca_certificate = base64decode(var.cluster_ca_certificate)
}

resource "kubernetes_namespace" "mihaiblebea" {
    metadata {
        name = "mihaiblebea"
    }
}

# blog microservice
resource "kubernetes_deployment" "blog-deployment" {
    metadata {
        name      = "blog-deployment"
        namespace = "mihaiblebea"
        labels    = {
            app  = "go-broadcast"
            name = "blog-deployment"
        }
    }

    spec {
        replicas = 1

        selector {
            match_labels = {
                app  = "go-broadcast"
                name = "blog-pod"
            }
        }

        template {
            metadata {
                name   = "blog-pod"
                labels = {
                    app  = "go-broadcast"
                    name = "blog-pod"
                }
            }

            spec {
                container {
                    image = var.blog_image
                    name  = "blog-container"
                    env {
                        name  = "HTTP_PORT"
                        value = var.http_port
                    }
                    env {
                        name = "LIST_URL"
                        value = "http://lists/lead"
                    }
                    port {
                        container_port = var.http_port
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "blog_node_port" {
    metadata {
        name      = "blog-service"
        namespace = "mihaiblebea"
    }

    spec {
        selector = {
            name = "blog-pod"
        }

        port {
            port        = 80
            target_port = var.http_port
            node_port   = var.node_port
        }

        type = "NodePort"
    }
}

resource "kubernetes_service" "blog_cluster_ip" {
    metadata {
        name      = "blog-service-cluster"
        namespace = "mihaiblebea"
    }

    spec {
        selector = {
            name = "blog-pod"
        }

        port {
            port        = 80
            target_port = var.http_port
        }

        type = "ClusterIP"
    }
}

# broadcast microservice
resource "kubernetes_deployment" "broadcast-deployment" {
    metadata {
        name      = "broadcast-deployment"
        namespace = "mihaiblebea"
        labels    = {
            app  = "go-broadcast"
            name = "broadcast-deployment"
        }
    }

    spec {
        replicas = 1

        selector {
            match_labels = {
                app  = "go-broadcast"
                name = "broadcast-pod"
            }
        }

        template {
            metadata {
                name   = "broadcast-pod"
                labels = {
                    app  = "go-broadcast"
                    name = "broadcast-pod"
                }
            }

            spec {
                container {
                    image = var.broadcast_image
                    name  = "broadcast-container"
                    env {
                        name  = "LINKEDIN_ACCESS_TOKEN"
                        value = var.linkedin_access_token
                    }

                    env {
                        name  = "TWITTER_CONSUMER_KEY"
                        value = var.twitter_consumer_key
                    }

                    env {
                        name  = "TWITTER_CONSUMER_SECRET"
                        value = var.twitter_consumer_secret
                    }

                    env {
                        name  = "TWITTER_TOKEN"
                        value = var.twitter_token
                    }

                    env {
                        name  = "TWITTER_TOKEN_SECRET"
                        value = var.twitter_token_secret
                    }

                    env {
                        name  = "POCKET_CONSUMER_KEY"
                        value = var.pocket_consumer_key
                    }

                    env {
                        name  = "POCKET_ACCESS_TOKEN"
                        value = var.pocket_access_token
                    }
                }
            }
        }
    }
}

# list microservice
resource "kubernetes_deployment" "list-deployment" {
    metadata {
        name      = "list-deployment"
        namespace = "mihaiblebea"
        labels    = {
            app  = "go-broadcast"
            name = "list-deployment"
        }
    }

    spec {
        replicas = 1

        selector {
            match_labels = {
                app  = "go-broadcast"
                name = "list-pod"
            }
        }

        template {
            metadata {
                name   = "list-pod"
                labels = {
                    app  = "go-broadcast"
                    name = "list-pod"
                }
            }

            spec {
                container {
                    image = var.list_image
                    name  = "list-container"
                    env {
                        name  = "HTTP_PORT"
                        value = var.list_http_port
                    }
                    env {
                        name  = "GOOGLE_CREDENTIALS_FILE"
                        value = var.google_credentials_file
                    }
                    env {
                        name  = "GOOGLE_TOKEN_FILE"
                        value = var.google_token_file
                    }
                    env {
                        name  = "AWS_ACCESS_KEY_ID"
                        value = var.aws_access_key_id
                    }
                    env {
                        name  = "AWS_SECRET_ACCESS_KEY"
                        value = var.aws_secret_access_key
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "list_cluster_ip" {
    metadata {
        name      = "lists"
        namespace = "mihaiblebea"
    }

    spec {
        selector = {
            name = "list-pod"
        }

        port {
            port        = 80
            target_port = var.list_http_port
        }

        type = "ClusterIP"
    }
}