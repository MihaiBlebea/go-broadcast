provider "digitalocean" {
    token   = var.do_token
    version = "1.22.0"
}

data "digitalocean_loadbalancer" "load_balancer" {
    name = var.load_balancer_name
}

resource "digitalocean_domain" "mihaiblebea_com" {
    name       = "mihaiblebea.com"
    ip_address = data.digitalocean_loadbalancer.load_balancer.ip
}

resource "digitalocean_certificate" "mihaiblebea" {
    name    = "mihaiblebea-cert"
    type    = "lets_encrypt"
    domains = [digitalocean_domain.mihaiblebea_com.name]
}