provider "digitalocean" {
    token   = var.do_token
    version = "1.22.0"
}

# data "digitalocean_loadbalancer" "load_balancer" {
#     name = var.load_balancer_name
# }

resource "digitalocean_domain" "mihaiblebea_com" {
    name       = "mihaiblebea.com"
    ip_address = digitalocean_loadbalancer.public.ip
}

resource "digitalocean_certificate" "mihaiblebea" {
    name    = "mihaiblebea-cert"
    type    = "lets_encrypt"
    domains = [digitalocean_domain.mihaiblebea_com.name]
}

resource "digitalocean_loadbalancer" "public" {
    name   = "loadbalancer-1"
    region = "lon1"

    forwarding_rule {
        entry_port     = 80
        entry_protocol = "http"

        target_port     = 30011
        target_protocol = "http"
    }

    forwarding_rule {
        entry_port     = 433
        entry_protocol = "https"

        target_port     = 30011
        target_protocol = "http"

        certificate_id = digitalocean_certificate.mihaiblebea.id
    }

    healthcheck {
        port     = 22
        protocol = "tcp"
    }

    redirect_http_to_https = true

    droplet_ids = [var.droplet_id]
}