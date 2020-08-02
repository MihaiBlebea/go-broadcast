resource "kubernetes_secret" "prod_secrets" {
    metadata {
        name      = "prod-secrets"
        namespace = "mihaiblebea"
    }

    data = {
        LINKEDIN_ACCESS_TOKEN   = file("${path.module}/secrets/linkedin_access_token.txt")
        TWITTER_CONSUMER_KEY    = file("${path.module}/secrets/twitter_consumer_key.txt")
        TWITTER_CONSUMER_SECRET = file("${path.module}/secrets/twitter_consumer_secret.txt")
        TWITTER_TOKEN           = file("${path.module}/secrets/twitter_token.txt")
        TWITTER_TOKEN_SECRET    = file("${path.module}/secrets/twitter_token_secret.txt")
    }
}