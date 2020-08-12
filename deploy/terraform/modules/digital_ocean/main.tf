provider "digitalocean" {
    token   = var.do_token
    version = "1.22.0"
}

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