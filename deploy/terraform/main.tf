
module "digital_ocean" {
    source          = "./modules/digital_ocean"

    do_token        = ""
    loadbalancer_ip = module.helm.loadbalancer_ip.load_balancer_ingress[0].ip
}

module "helm" {
    source = "./modules/helm"

    kubernetes_host        = module.digital_ocean.kubernetes_host
    kubernetes_token       = module.digital_ocean.kubernetes_token
    cluster_ca_certificate = module.digital_ocean.cluster_ca_certificate
}

module "kubernetes" {
    source                 = "./modules/kubernetes"

    blog_image             = "serbanblebea/go-blog:v0.5"
    broadcast_image        = "serbanblebea/go-broadcast:v0.4"
    kubernetes_host        = module.digital_ocean.kubernetes_host
    kubernetes_token       = module.digital_ocean.kubernetes_token
    cluster_ca_certificate = module.digital_ocean.cluster_ca_certificate
}