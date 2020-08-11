terraform {
    required_version = "0.12.29"

    backend "remote" {
        organization = "PurpleTreeTech"

        workspaces {
            name = "go-broadcast"
        }
    }
}

module "digital_ocean" {
    source          = "./modules/digital_ocean"

    do_token        = var.do_token
    # loadbalancer_ip = module.helm.loadbalancer_ip.load_balancer_ingress[0].ip
}

# module "helm" {
#     source = "./modules/helm"

#     kubernetes_host        = module.digital_ocean.kubernetes_host
#     kubernetes_token       = module.digital_ocean.kubernetes_token
#     cluster_ca_certificate = module.digital_ocean.cluster_ca_certificate
# }

module "kubernetes" {
    source                  = "./modules/kubernetes"

    kubernetes_host         = module.digital_ocean.kubernetes_host
    kubernetes_token        = module.digital_ocean.kubernetes_token
    cluster_ca_certificate  = module.digital_ocean.cluster_ca_certificate

    blog_image              = var.blog_image
    broadcast_image         = "serbanblebea/go-broadcast:v0.4"

    linkedin_access_token   = var.linkedin_access_token
    twitter_consumer_key    = var.twitter_consumer_key
    twitter_consumer_secret = var.twitter_consumer_secret
    twitter_token           = var.twitter_token
    twitter_token_secret    = var.twitter_token_secret
}