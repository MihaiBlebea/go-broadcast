terraform {
    required_version = "0.12.20"

    backend "remote" {
        organization = "PurpleTreeTech"

        workspaces {
            name = "go-broadcast"
        }
    }
}

module "digital_ocean" {
    source   = "./modules/digital_ocean"
    do_token = var.do_token
}

module "digital_ocean_lb" {
    # depends_on         = [module.kubernetes]
    source             = "./modules/digital_ocean_lb"

    do_token           = var.do_token
    load_balancer_name = "blogloadbalancer"
    # loadbalancer_ip   = module.kubernetes.loadbalancer_raw.load_balancer_ingress[0].ip
}

module "kubernetes" {
    source                  = "./modules/kubernetes"

    kubernetes_host         = module.digital_ocean.kubernetes_host
    kubernetes_token        = module.digital_ocean.kubernetes_token
    cluster_ca_certificate  = module.digital_ocean.cluster_ca_certificate

    # certificate_id          = module.digital_ocean_lb.certificate_id
    load_balancer_name      = "blogloadbalancer"

    blog_image              = var.blog_image
    broadcast_image         = "serbanblebea/go-broadcast:v0.4"

    linkedin_access_token   = var.linkedin_access_token
    twitter_consumer_key    = var.twitter_consumer_key
    twitter_consumer_secret = var.twitter_consumer_secret
    twitter_token           = var.twitter_token
    twitter_token_secret    = var.twitter_token_secret
}