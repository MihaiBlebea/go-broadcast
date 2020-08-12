variable "do_token" {
    description = "digital ocean api token"
    type        = string
}

variable "kubeconfig_path" {
    description = "the path to save the kubeconfig file to"
    default     = "~/.kube/test-config"
}

variable "cluster_tags" {
    description = "default cluster tags"
    default     = ["k8s-cluster"]
}

variable "node_tags" {
    description = "default node tags"
    default     = ["worker-node"]
}


