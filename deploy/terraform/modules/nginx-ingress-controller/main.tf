provider "helm" {
    version = ">= 1.0.0"
    debug   = true
    alias   = "helm"

    kubernetes {
        config_path = "/Users/mihaiblebea/.kube/config"
    }
}

# resource "helm_release" "nginx_ingress_controller" {
#     name  = "nginx-ingress-controller"
#     chart = "stable/nginx-ingress"
#     # repository = data.helm_repository.stable.metadata.0.name
#     namespace = "ingress-nginx"
#     atomic = true
    
#     # set {
#     #     name  = "mariadbUser"
#     #     value = "foo"
#     # }
# }

resource "helm_release" "nginx-ingress" {
    name              = "nginx-ingress"
    repository        = "https://kubernetes-charts.storage.googleapis.com/"
    chart             = "nginx-ingress"
    namespace         = "ingress-nginx"
    # version           = var.chart_version[var.environment]
    # dependency_update = true
    # timeout           = 600
    wait              = false

    # values = [
    #     "${file("${var.ci_project_dir}/services-out/nginx-ingress-${var.environment}.yaml")}"
    # ]

    # set {
    #     name  = "history-max"
    #     value = 20
    # }
}