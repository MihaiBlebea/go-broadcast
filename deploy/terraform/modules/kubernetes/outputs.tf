output "loadbalancer_ip" {
    value = kubernetes_service.blog_load_balancer.ip
}

output "loadbalancer_hostname" {
    value = kubernetes_service.blog_load_balancer.hostname
}