provider "helm" {
    version = ">= 1.0.0"
    debug   = true

    kubernetes {
        # config_path = "/Users/mihaiblebea/.kube/config"
        load_config_file       = false
        host                   = var.kubernetes_host
        token                  = var.kubernetes_token
        cluster_ca_certificate = base64decode(var.cluster_ca_certificate)
    }
}

provider "kubernetes" {
    load_config_file       = false
    host                   = var.kubernetes_host
    token                  = var.kubernetes_token
    cluster_ca_certificate = base64decode(var.cluster_ca_certificate)
}

resource "kubernetes_namespace" "ingress-nginx" {
    metadata {
        name = "ingress-nginx"
    }
}

resource "helm_release" "nginx-ingress" {
    name              = "nginx-ingress"
    repository        = "https://kubernetes-charts.storage.googleapis.com/"
    chart             = "nginx-ingress"
    namespace         = "ingress-nginx"
    wait              = false
    atomic            = true

    set {
        name  = "controller.publishService.enabled"
        value = true
    }
}

data "kubernetes_service" "loadbalancer" {
    metadata {
        name = "nginx-ingress-controller"
        namespace = "ingress-nginx"
    }
}
