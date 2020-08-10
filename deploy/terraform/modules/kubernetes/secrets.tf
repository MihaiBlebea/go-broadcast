resource "kubernetes_secret" "prod_secrets" {
    metadata {
        name      = "prod-secrets"
        namespace = "mihaiblebea"
    }

    data = {
        LINKEDIN_ACCESS_TOKEN   = var.linkedin_access_token
        TWITTER_CONSUMER_KEY    = var.twitter_consumer_key
        TWITTER_CONSUMER_SECRET = var.twitter_consumer_secret
        TWITTER_TOKEN           = var.twitter_token
        TWITTER_TOKEN_SECRET    = var.twitter_token_secret
    }
}