provider "kubernetes" {
    # Specify the path to the kubernetes config file
    config_path = "/Users/mihaiblebea/.kube/config"
}

resource "kubernetes_namespace" "dev" {
    metadata {
        name = "dev"
    }
}

resource "kubernetes_deployment" "hello-deployment" {
    metadata {
        name = "hello-deployment"
        labels = {
            app = "go-broadcast"
            name = "hello-deployment"
        }
    }

    spec {
        replicas = 1

        selector {
            match_labels = {
                app = "go-broadcast"
                name = "hello-world-app"
            }
        }

        template {
            metadata {
                name = "hello-world-app"
                labels = {
                    app = "go-broadcast"
                    name = "hello-world-app"
                }
            }

            spec {
                container {
                    image = "serbanblebea/go-broadcast:v0.1"
                    name = "hello-world-pod"
                    env {
                        name = "HTTP_PORT"
                        value = 8099
                    }
                    port {
                        container_port = 8099
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "hello-service" {
    metadata {
        name = "hello-world-service"
    }

    spec {
        selector = {
            name = "hello-world-app"
        }

        port {
            port = 80
            target_port = 8099
            node_port = 30011
        }

        type = "NodePort"
    }
}