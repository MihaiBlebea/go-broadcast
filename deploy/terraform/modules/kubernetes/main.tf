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

# resource "kubernetes_service" "blog_load_balancer" {
#     metadata {
#         name      = "blog-load-balancer"
#         namespace = "mihaiblebea"

#         annotations = {
#             "service.beta.kubernetes.io/do-loadbalancer-name"                   = var.load_balancer_name
#             "service.beta.kubernetes.io/do-loadbalancer-protocol"               = "http"
#             "service.beta.kubernetes.io/do-loadbalancer-tls-ports"              = "443"
#             "service.beta.kubernetes.io/do-loadbalancer-redirect-http-to-https" = "true"
#             # "service.beta.kubernetes.io/do-loadbalancer-certificate-id"         = var.certificate_id
#         }
#     }

#     spec {
#         selector = {
#             name = "blog-pod"
#         }

#         port {
#             name        = "http"
#             port        = 80
#             target_port = var.http_port
#         }

#         port {
#             name        = "https"
#             port        = 443
#             target_port = var.http_port
#         }

#         type = "LoadBalancer"
#     }
# }

resource "kubernetes_cron_job" "broadcast" {
    metadata {
        name      = "broadcast"
        namespace = "mihaiblebea"
    }

    spec {
        concurrency_policy            = "Replace"
        failed_jobs_history_limit     = 5
        schedule                      = "* * * * *"
        starting_deadline_seconds     = 10
        successful_jobs_history_limit = 10
        suspend                       = true

        job_template {
            metadata {}

            spec {
                backoff_limit              = 2
                ttl_seconds_after_finished = 10

                template {
                    metadata {}

                    spec {
                        container {
                            name    = "broadcast"
                            image   = var.broadcast_image
                            # command = ["/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"]

                            env_from {
                                secret_ref {
                                    name = "prod-secrets"
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}