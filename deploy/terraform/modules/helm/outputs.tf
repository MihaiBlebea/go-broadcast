output "loadbalancer_ip" {
    value = data.kubernetes_service.loadbalancer
}