
module "digital_ocean" {
    source          = "./modules/digital_ocean"

    do_token        = "d4d909d908f0759fbe2c4ab0928c2b9e8e0a79a5dd28929ddbc7421fcfa756a9"
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

    blog_image             = "serbanblebea/go-blog:v0.4"
    broadcast_image        = "serbanblebea/go-broadcast:v0.4"
    kubernetes_host        = module.digital_ocean.kubernetes_host
    kubernetes_token       = module.digital_ocean.kubernetes_token
    cluster_ca_certificate = module.digital_ocean.cluster_ca_certificate
}