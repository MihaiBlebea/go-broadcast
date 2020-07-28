provider "kubernetes" {
    # Specify the path to the kubernetes config file
    config_path = "/Users/mihaiblebea/.kube/config"
}


resource "kubernetes_namespace" "mihaiblebea" {
    metadata {
        name = "mihaiblebea"
    }
}

resource "kubernetes_namespace" "ingress-nginx" {
    metadata {
        name = "ingress-nginx"
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
                    image = var.blog_image
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

resource "kubernetes_service" "blog_node_port" {
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
            node_port = var.node_port
        }

        type = "NodePort"
    }
}

resource "kubernetes_service" "blog_cluster_ip" {
    metadata {
        name = "blog-service-cluster"
        namespace = "mihaiblebea"
    }

    spec {
        selector = {
            name = "blog-pod"
        }

        port {
            port = 80
            target_port = var.http_port
        }

        type = "ClusterIP"
    }
}

resource "kubernetes_secret" "user_password" {
    metadata {
        name = "user-password"
        namespace = "mihaiblebea"
    }

    data = {
        password = file("${path.cwd}/secrets/password.txt")
    }
}

resource "kubernetes_ingress" "blog-ingress" {
    metadata {
        name = "blog-ingress"
        namespace = "mihaiblebea"
        annotations = {
            "nginx.ingress.kubernetes.io/proxy-body-size" = "20m"
            "kubernetes.io/ingress.class" = "nginx"
            "nginx.ingress.kubernetes.io/rewrite-target" = "/$2"
        }
    }

    spec {

        rule {
            host = "mihaiblebea.com"

            http {
                path {
                    path = "/app1/(/|$)(.*)"

                    backend {
                        service_name = "blog-service-cluster"
                        service_port = 80
                    }
                }
            }
        }
    }
}