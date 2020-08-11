# output "loadbalancer_ip" {
#     value = kubernetes_service.blog_load_balancer.ip
# }

# output "loadbalancer_hostname" {
#     value = kubernetes_service.blog_load_balancer.hostname
# }

output "loadbalancer_raw" {
    value = kubernetes_service.blog_load_balancer
}