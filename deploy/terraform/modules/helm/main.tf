provider "helm" {
    version = ">= 1.0.0"
    debug   = true
    alias   = "helm"

    kubernetes {
        config_path = "/Users/mihaiblebea/.kube/config"
    }
}

resource "helm_release" "nginx-ingress" {
    name              = "nginx-ingress"
    repository        = "https://kubernetes-charts.storage.googleapis.com/"
    chart             = "nginx-ingress"
    namespace         = "ingress-nginx"
    wait              = false
}