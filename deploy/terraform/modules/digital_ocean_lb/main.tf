provider "digitalocean" {
    token   = var.do_token
    version = "1.22.0"
}

resource "digitalocean_domain" "mihaiblebea_com" {
    name       = var.domain_name
    ip_address = digitalocean_loadbalancer.public.ip
}

resource "digitalocean_certificate" "mihaiblebea" {
    name    = "mihaiblebea-cert"
    type    = "lets_encrypt"
    domains = [var.domain_name]
}

resource "digitalocean_record" "txt_google_search_console" {
    domain   = var.domain_name
    type     = "TXT"
    name     = "@"
    priority = 10
    value    = var.google_search_console_code
}

# AWS SES register domain records
resource "digitalocean_record" "aws_domain_verification_record" {
    domain   = digitalocean_domain.mihaiblebea_com.name
    type     = "TXT"
    name     = "_amazonses.${digitalocean_domain.mihaiblebea_com.name}"
    priority = 10
    value    = var.aws_domain_verification_token
}

resource "digitalocean_record" "aws_domain_key_set_one" {
    domain   = digitalocean_domain.mihaiblebea_com.name
    type     = "CNAME"
    name     = "64aonjzwf22bgfitrab5nbu6v2rqnard._domainkey.${digitalocean_domain.mihaiblebea_com.name}"
    priority = 10
    value    = "64aonjzwf22bgfitrab5nbu6v2rqnard.dkim.amazonses.com"
}

resource "digitalocean_record" "aws_domain_key_set_two" {
    domain   = digitalocean_domain.mihaiblebea_com.name
    type     = "CNAME"
    name     = "xhrue3rk4uzwfuskugzrngyxt3jloi5m._domainkey.${digitalocean_domain.mihaiblebea_com.name}"
    priority = 10
    value    = "xhrue3rk4uzwfuskugzrngyxt3jloi5m.dkim.amazonses.com"
}

resource "digitalocean_record" "aws_domain_key_set_three" {
    domain   = digitalocean_domain.mihaiblebea_com.name
    type     = "CNAME"
    name     = "i2njichgbwkz2xs2suagqzhckea7bvyd._domainkey.${digitalocean_domain.mihaiblebea_com.name}"
    priority = 10
    value    = "i2njichgbwkz2xs2suagqzhckea7bvyd.dkim.amazonses.com"
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
        entry_port     = 443
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