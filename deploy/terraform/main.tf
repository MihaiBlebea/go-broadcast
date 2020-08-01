
module "digital_ocean" {
    source = "./modules/digital_ocean"

    do_token        = "37e43fa89c6321a7952a7ceade620ac11b007604d112a9b21e71687fb3578265"
    loadbalancer_ip = module.helm.loadbalancer_ip.load_balancer_ingress[0].ip
}

module "helm" {
    source = "./modules/helm"

    kubernetes_host        = module.digital_ocean.kubernetes_host
    kubernetes_token       = module.digital_ocean.kubernetes_token
    cluster_ca_certificate = module.digital_ocean.cluster_ca_certificate
}

module "kubernetes" {
    source = "./modules/kubernetes"

    kubernetes_host        = module.digital_ocean.kubernetes_host
    kubernetes_token       = module.digital_ocean.kubernetes_token
    cluster_ca_certificate = module.digital_ocean.cluster_ca_certificate
}