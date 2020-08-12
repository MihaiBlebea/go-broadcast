output "kubernetes_host" {
    value = digitalocean_kubernetes_cluster.cluster.endpoint
}

output "kubernetes_token" {
    value = digitalocean_kubernetes_cluster.cluster.kube_config[0].token
}

output "cluster_ca_certificate" {
    value = digitalocean_kubernetes_cluster.cluster.kube_config[0].cluster_ca_certificate
}

# output "certificate_id" {
#     value = digitalocean_certificate.mihaiblebea.id
# }