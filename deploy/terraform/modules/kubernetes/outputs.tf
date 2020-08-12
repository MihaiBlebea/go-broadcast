output "loadbalancer_raw" {
    value = kubernetes_service.blog_load_balancer
}

output "load_balancer_name" {
    value = var.load_balancer_name
}