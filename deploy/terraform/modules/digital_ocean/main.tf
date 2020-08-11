provider "digitalocean" {
    token   = var.do_token
    version = "1.22.0"
}

# resource "digitalocean_domain" "mihaiblebea_com" {
#     name       = "mihaiblebea.com"
#     ip_address = digitalocean_loadbalancer.public.ip
# }

# resource "digitalocean_certificate" "mihaiblebea" {
#     name    = "mihaiblebea-cert"
#     type    = "lets_encrypt"
#     domains = [digitalocean_domain.mihaiblebea_com.name]
# }

# resource "digitalocean_loadbalancer" "public" {
#     name   = "loadbalancer-1"
#     region = "lon1"

#     forwarding_rule {
#         entry_port     = 80
#         entry_protocol = "http"

#         target_port     = 80
#         target_protocol = "http"
#     }

#     healthcheck {
#         port     = 22
#         protocol = "tcp"
#     }

#     droplet_ids = [digitalocean_kubernetes_cluster.cluster.node_pool[0].nodes[0].droplet_id]
# }

resource "digitalocean_kubernetes_cluster" "cluster" {
    name    = "blog-k8-cluster-1"
    region  = "lon1"
    version = "1.18.6-do.0"
    tags    = var.cluster_tags

    node_pool {
        name       = "worker-pool"
        size       = "s-1vcpu-2gb"
        node_count = 1
        tags       = var.node_tags
    }
}

resource "local_file" "kubeconfig" {
    content  = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
    filename = pathexpand(var.kubeconfig_path)
}