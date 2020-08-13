---
Title: How to build infrastructure with code, using Terraform with Kubernetes - Part 1
Summary: We are going to build a CI/CD pipeline from scratch. It's probably not going to be production ready, but will work for most individual developers
Image: https://www.terraform.io/assets/images/og-image-large-e60c82fe.png
Tags:
    - terraform
    - kubernetes
    - infrastructure
    - devops
Layout: article
Slug: terraform-with-kubernetes-infrastructure-as-code-part-1
Published: "2020-08-12 10:04:05"
Kind: article
---

What is this infrastructure as code all about? And why is everybody talking about it? 

We as developers are starting to rely too much on tools and helpers, just so we can do the same thing in a slightly better way... or not even better, but in a different way.

How do you know where to draw a line in the sand and just say enough is enough?

<img src="/static/img/devops.jpg" />

I would say that infrastructure as code is not a new concept anymore, and the benefits that it can bring to your project, heavily outweight the **pain** of learning another tool.

Terraform and Kubernetes are definitely usefull tools to have in your developer arsenal.


When everything is finished and ready to go live, you start working on deploying this code into production.

But wait...

There is no clear documentation for the current infrastructure.


Then **Alex** points out that the company is using S3 to store credentials for the production environment, but there is no bucket called credentials.

Also, **Mary** says that there might be a load balancer somewhere that rutes traffic between the different staging envs...

Cherry on cake, the main platform engineer who holds the keys to all the infrastructure has decided to leave the company 😂.

You are left with **figuring out on yourself how the current infrastructure works** and how to deploy your code into production.

This is the time when you would want a tool like Terraform to come to the rescue.

**Jack** can bet £5 that there are 8 boxes in a AWS account, but you can find just 3 of them.
## How to become a better developer by using infrastructure as code

## What are we trying to achieve

## Let's build some infrastructure

### Baby steps

If you never used Terraform before, I suggest you start with the official documentation, which is by far one of the best documentation I came across.

<a href="https://www.terraform.io/docs/index.html" target="_blank">Start with the Terraform docs</a>

Now that we got that out of the way, let's have some fun!

First, we will create a ```/deploy/terraform/main.tf``` file in our project root folder.

This will act as the **"Hello world"** in the world of Terraform.

```hcl
provider "digitalocean" {
    token   = xxxxxxxxxxxxxx-xxxxxxx-xxxxxxxxx
    version = "1.22.0"
}
```
This will tell Terraform that we want to deploy our infrastructure on Digital Ocean by using the `digitalocean` provider.

A provider is just a binary that acts as a plugin to the main Terraform binary and adds an abstraction between the Digital Ocean API and our infrastructure as code (if we can call it that, at this stage 😂) 

There are many more `providers` from which you can choose and I highly recommend that you take a look over the list. Btw, did you know that you can create your own provider 😱? (we'll leave that for another time though)

➡️ <a href="https://www.terraform.io/docs/providers/index.html" target="_blank">Terrafrom providers list</a>

The `token` in the provider block is the Digital Ocean API token, that you can generate in your DO account.

Next, we add the version of the provider that we want to use. You may think that it's oki to skip the version, but I have a nice story about how I lost 2 days trying to debug a ci/cd pipeline, just to find out that I was using 2 different versions that didn't get along so well.

Advice to future self: `always use versions if the option is available` 😂

Oki, all good now, we have the provider. What next?

Let's create a Kubernetes cluster resource. We will add this code to our `main.tf` file.

```hcl
resource "digitalocean_kubernetes_cluster" "cluster" {
    name    = "blog-k8-cluster-1"
    region  = "lon1"
    version = "1.18.6-do.0"
    tags    = ["cluster", "app-name"]

    node_pool {
        name       = "worker-pool"
        size       = "s-1vcpu-2gb"
        node_count = 2
        tags       = ["app-node", "app-name"]
    }
}
```

This block is a resource in Terraform world and represents the configuration for a DO kubernetes cluster.

We will give it a name, region, and kubernetes `version` (always use the version 😂).

This cluster wiil need to run on a DO droplet, which will be automaticaly created when we run this main.tf file. You don't have to worry about that.

Use `node_pool` to specify the name, size (s-1vcpu-2gb) and how many cluster nodes do you need (I recommend you go for a node_count > 1, so your app will not go down during updates to one of the nodes).

Tags are optional, so you can skip this step if you don't have a specific case for using them.

If you run this Terraform file right now, with `terraform init` and then `terraform plan`, you will see some data in HCL (Hashicorp Language) with all the resources that you plan to generate.

Next, if you are oki with how things look, run `terraform apply` and you will be prompted to approve the changes with a simple `yes`. Or, in my case you will get an error saying that the DO api token is not valid 😂😂😂.

> Don't forget to replace the dummy Digital Ocean API token with the real one

<img src="/static/img/yeey.gif" />

YEEY! You got a Digital Ocean Kubernetes cluster ready to run your awesome application.

But we are not quite finished yet.

We need to get the config file so we can connect to our cluster from the confort of our living room 😂.

Add this to the `main.tf`

```hcl
resource "local_file" "kubeconfig" {
    content  = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
    filename = pathexpand("~/.kube/test-config")
}
```

What this does, is downloading the `kubeconfig` file from DO to your specified local path.

If you have `kubectl` installed on your machine, then you can now connect to your cluster and start deploying your pods and services.

We have just one single issue to handle... **We need a load balancer** to direct outside traffic to our services and pods.

This can be done with deploying a `LoadBalancer` type of Service, and this will auto-create a Load Balancer in your DO account, but I do not recommend this approach.

Why? Because this auto-generated LB will be outside of the scope of Terraform.

If you wish, it will be like a function with `side-effects`.

This means that we canot control that resource with Terraform and all our hard work would be for nothing, if you are always required to log into the DO platform and configure the Load Balancer by hand.

So we will create a LB with Terraform.

In the same `main.tf` file, add this block.

```hcl
resource "digitalocean_loadbalancer" "public" {
    name   = "loadbalancer-made-up-name-1"
    region = "lon1"

    forwarding_rule {
        entry_port     = 80
        entry_protocol = "http"

        target_port     = 30011
        target_protocol = "http"
    }

    forwarding_rule {
        entry_port     = 443
        entry_protocol = "https"

        target_port     = 30011
        target_protocol = "http"

        certificate_id = 123-123-456-xxxxx
    }

    healthcheck {
        port     = 22
        protocol = "tcp"
    }

    redirect_http_to_https = true

    droplet_ids = [droplet_id_that_we_have_to_add_later]
}
```

Let's take it step by step.

- `name` and `region`, this two are more or less self-explanatory
- `forwarding_rule` is the block where we specify where do we want our traffic to go. For a start, we want to route port 80 on the LB to port 30011 on the droplet (this is the port exposed by our Kubernetes NodePort service - can be anything you like)
- second `forwarding_rule` is for taking the 433 port on the LB to the same 30011 port on the droplet. This config block requires a certificate_id that we will add at the next step. Don't forget to specify the `entry_protocol` and `target_protocol`
- `healthcheck` indicates the config for the health check that DO will perform for you - nothing too fancy here
- `redirect_http_to_https` add this if you want all your http traffic to be redirected to https - why would you not want this 😂
- `droplet_ids` represents a list of droplet ids that this LB will direct traffic to. If you have experience with Kubernetes, imagine this LB as a service that links to a number of pods with the help of labels.

We still have some missing pieces that we need to add before we run this configuration.

If you didn't spotted them already, I am talking about that `domain` and the `certificate_id`.

We will need to add two more blocks to the `main.tf` file.

**Domain**

```hcl
resource "digitalocean_domain" "mihaiblebea_com" {
    name       = "mihaiblebea.com"
    ip_address = digitalocean_loadbalancer.public.ip
}
```

The `name` is the actual domain that you probably bought from GoDaddy or some other third party provider.

`ip_address` is the ip where we want our domain to point to. If you specify this optional config here, an A record will be generated in the DO domain settings. Pretty sweet, no need to touch that anymore. 😂

**Certificate**

```hcl
resource "digitalocean_certificate" "mihaiblebea" {
    name    = "mihaiblebea-cert"
    type    = "lets_encrypt"
    domains = ["mihaiblebea.com"]
}
```

You can generate your own certificate and upload it, but in this case we will use `type = "lets_encrypt"`.

If you run `terraform plan` & `terraform apply` again, you will see that a LB is added to your DO account. (no need to run `terraform init` anymore).

Go ahead and commit this `main.tf` file to your Github or Bitbucket account and you can share it with your team or make incremental changes to it, without ever going back to the DO platform.

When you want to take a break from your DO resources, just run `terraform destroy` and confirm with `yes` when the prompt appears.

You can copy the full `/deploy/terraform/main.tf` file from here:

```hcl
provider "digitalocean" {
    token   = xxxxxxxxxxxxxx-xxxxxxx-xxxxxxxxx
    version = "1.22.0"
}

resource "digitalocean_kubernetes_cluster" "cluster" {
    name    = "blog-k8-cluster-1"
    region  = "lon1"
    version = "1.18.6-do.0"
    tags    = ["cluster", "app-name"]

    node_pool {
        name       = "worker-pool"
        size       = "s-1vcpu-2gb"
        node_count = 2
        tags       = ["app-node", "app-name"]
    }
}

resource "local_file" "kubeconfig" {
    content  = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
    filename = pathexpand("~/.kube/test-config")
}

resource "digitalocean_loadbalancer" "public" {
    name   = "loadbalancer-made-up-name-1"
    region = "lon1"

    forwarding_rule {
        entry_port     = 80
        entry_protocol = "http"

        target_port     = 30011
        target_protocol = "http"
    }

    forwarding_rule {
        entry_port     = 443
        entry_protocol = "https"

        target_port     = 30011
        target_protocol = "http"

        certificate_id = digitalocean_certificate.mihaiblebea.id
    }

    healthcheck {
        port     = 22
        protocol = "tcp"
    }

    redirect_http_to_https = true

    droplet_ids = [digitalocean_kubernetes_cluster.cluster.node_pool[0].nodes[0].droplet_id]
}

resource "digitalocean_domain" "mihaiblebea_com" {
    name       = "mihaiblebea.com"
    ip_address = digitalocean_loadbalancer.public.ip
}

resource "digitalocean_certificate" "mihaiblebea" {
    name    = "mihaiblebea-cert"
    type    = "lets_encrypt"
    domains = ["mihaiblebea.com"]
}
```

In **Part 2** we will add variables, outputs and start using Terraform modules to make our infrastructure as code more modular and to support a stagging and a production environment, without copy-pasting code around.

> There is always a high risk when using Terraform to commit passwords to Github if you are copy-pasting them directly in your `tf` files, so always be very careful about doing so. 

Stay close for part 2, where I will show you how to use Terraform variables to hide those passwords.

<img src="/static/img/home_alone.jpg" />