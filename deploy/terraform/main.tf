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
    source     = "./modules/digital_ocean_lb"

    do_token                   = var.do_token
    domain_name                = var.domain_name
    google_search_console_code = var.google_search_console_code
    droplet_id                 = module.digital_ocean.droplet_id

    aws_domain_verification_token = var.aws_domain_verification_token
    aws_domain_key_set            = var.aws_domain_key_set
}

module "kubernetes" {
    source                  = "./modules/kubernetes"

    kubernetes_host         = module.digital_ocean.kubernetes_host
    kubernetes_token        = module.digital_ocean.kubernetes_token
    cluster_ca_certificate  = module.digital_ocean.cluster_ca_certificate

    blog_image              = var.blog_image
    broadcast_image         = var.broadcast_image

    linkedin_access_token   = var.linkedin_access_token
    twitter_consumer_key    = var.twitter_consumer_key
    twitter_consumer_secret = var.twitter_consumer_secret
    twitter_token           = var.twitter_token
    twitter_token_secret    = var.twitter_token_secret
    pocket_consumer_key     = var.pocket_consumer_key
    pocket_access_token     = var.pocket_access_token

    list_image              = var.list_image
    google_credentials_file = var.google_credentials_file
    google_token_file       = var.google_token_file
    aws_access_key_id       = var.aws_access_key_id
    aws_secret_access_key   = var.aws_secret_access_key
}