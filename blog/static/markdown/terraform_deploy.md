---
Title: How to build infrastructure with code, using Terraform with Kubernetes - Part 1
Summary: What is this infrastructure as code all about? If you ever had to deal with infrastructure built by somebody else and passed onto you without documentation, then this article is for you.
Image: https://www.terraform.io/assets/images/og-image-large-e60c82fe.png
Tags:
    - terraform
    - kubernetes
    - infrastructure
    - devops
Slug: terraform-with-kubernetes-infrastructure-as-code-part-1
Published: "2020-08-12 10:04:05"
---

What is this infrastructure as code all about?

Why can't we just do infrastructure the old way as we have done for ages?

There are some developers that think that we somehow rely too much on tools in our day to day jobs. **And they may be right.**

Just take a look at how many different tools a modern developer must know just to stay on top of things...

<img src="/static/img/devops.jpg" />

It's **trully amazing** how developers can master them all...

## Infrastructure as code is nothing new

I'm not going to go into a history lesson of AWS or anything like that, but let me know if this scenario sounds familiar to you?

<img src="https://media.giphy.com/media/WoWm8YzFQJg5i/giphy.gif">

Let's say you are joining a new company and you are excited to start contributing to the code base.

You want to impress your new manager so you over-deliver on this new shiny microservice that everybody expects to be a hit.

When everything is finished and ready to go live, you start working on deploying this code into production.

But wait...

There is no clear documentation for the current infrastructure.

**Jack**, the senior dev, can bet £5 that there are 8 boxes in a AWS account, but you can find just 3 of them.

Then **Alex** points out that the company is using S3 to store credentials for the production environment, but there is no bucket called credentials.

Also, **Mary** says that there might be a load balancer somewhere that routes traffic between the different staging envs...

Cherry on the cake, the main platform engineer who holds the keys to all the infrastructure has decided to leave the company 😂.

You are left with **figuring out on yourself how the current infrastructure works** and how to deploy your code into production.

<img src="/static/img/head_in_palm_statue.jpg">

This is the time when you would want a tool like **Terraform** to come to the rescue.

## How to become a better developer by using Terraform to build and deploy infrastructure as code

I would say that infrastructure as code is not a new concept anymore, and the benefits that it can bring to your project, heavily outweight the **pain** of learning a new tool.

In this article, I will walk you towards deploying a Golang application with Kubernetes and Terraform.

If keep on reading till the end, you will:

- Get a better understanding on how Terraform works, so **you can create infrastructure from config files, destroy it and build it again**, all in under 5 minutes

- Find out how to create a Kubernetes cluster on Digital Ocean so **you can improve your platform in incremental steps,** without the need to touch the web interface of your cloud provider of choice

- Learn how to use Terraform modules to make your infrastructure more modular and **create multiple environments from the same config files**

## Before we start you will need to have this prepared

- Digital Ocean accunt - the only time you will need to touch the web platform

- Digital Ocean API token - easy to get once you finished the step above

- Terraform installation on your laptop - <a href="https://learn.hashicorp.com/tutorials/terraform/install-cli" target="_blank">TF installation guide</a>

- Kubectl on your local laptop - <a href="https://kubernetes.io/docs/tasks/tools/install-kubectl" target="_blank">Kubectl installation guide</a>

Also you need to know what are the different parts that compose Terraform.

I'll keep this short so we can start building some infrastructure as soon as possible - please jump over this part if you already know this.

Terraform is a tool that help you define infrastructure as code to manage the full lifecycle — create new resources, manage existing ones, and destroy those no longer needed.

The 3 main building blocks of TF are:

- **Providers**, that tells TF where to deploy your infrastructure: AWS, Kubernetes, DigitalOcean, GCP, etc.

- **Resources**, that define what you want to deploy: LoadBalancer, Droplet, KubernetesCluster, etc

- **Data sources**, this are resources that are not managed by TF, but still you can read data from them.

Terraform works with declarative configuration, this means that you don't need to tell TF what steps it needs to take to achive the end result (like Ansible) - you just need to tell it what the expected end result is, and it will take care of the rest.

Also...

> Terraform commands are idempotent 🤷

...this means that you can use the TF command over the same config over and over again and get the same expected result. Maybe it's not making a lot of sense now, but this is one amazing feature.

## It's time to build some infrastructure

If you never used Terraform before, I suggest you start with the official documentation, which is by far one of the best docs I came across.

<a href="https://www.terraform.io/docs/index.html" target="_blank">Start with the Terraform docs</a>

Now that we got that out of the way, let's have some fun!

First, go to your application folder and create this file ```/deploy/terraform/main.tf```. *(Notice that I put it inside deploy and terraform folders)*

This file will be the base of our infrastructure configuration.

Think of it as the **Hello world** of the Terraform world.

We will write our configuration in a the **HCL** language - Hashicorp Configuration Language.

```hcl
provider "digitalocean" {
    token   = xxxxxxxxxxxxxx-xxxxxxx-xxxxxxxxx
    version = "1.22.0"
}
```
This will tell Terraform that we want to deploy our infrastructure on Digital Ocean by using the `digitalocean` provider.

A provider is just an abstraction between the Digital Ocean API and our **configuration file** (if we can call it that, at this stage 😂) 

There are many more `providers` from which you can choose and I highly recommend that you take a look over the list. Btw, did you know that you can create your own provider 😱? (we'll leave that for another time though)

➡️ <a href="https://www.terraform.io/docs/providers/index.html" target="_blank">Terrafrom providers list</a>

You probably noticed that inside the provider block we have a couple of other configs.

The `token` is the Digital Ocean API token, that you can generate in your DO account.

Next, we add the version of the provider that we want to use. You may think that it's oki to skip the version, but I have a nice story about how I lost 2 days trying to debug a ci/cd pipeline, just to find out that I was using 2 different versions that didn't get along so...

Advice for future self:

> Always specify the version of the tool that you are using.

Oki, all good now, we have the provider. What next?

## Create a Kubernetes cluster

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

This block is a Terraform **resource** and represents the configuration for a DO kubernetes cluster.

We will give it a name, region, and kubernetes `version` (always use the version 😂).

This cluster wiil need to run on a DO droplet, which will be automaticaly created when we run this main.tf file. You don't have to worry about that.

Use `node_pool` to specify the name, size (**s-1vcpu-2gb**) and the nodes count (I recommend you go for a node_count > 1, so your app will not go down during updates).

Tags are optional, so **you can skip this step** if you don't have a specific case for using them.

## Let's run the plan and see what we got

If you run this Terraform file right now, with `terraform init` and then `terraform plan`, you will see some data in HCL with all the resources that we plan to generate.

Next, after you checked that everything looks oki-ish 😂, run `terraform apply` and you will be prompted to approve the changes with a simple `yes`. 

Or, in my case you will get an error saying that the DO api token is not valid 😂😂😂.

> Don't forget to replace the dummy Digital Ocean API token with the real one

<img src="/static/img/yeey.gif" />

YEEY! You got a Digital Ocean Kubernetes cluster ready to host your amazing app.

**But we are not quite finished yet.**

We need to get the config file so we can connect to our cluster from the confort of our living room 😂.

Add this to the `main.tf`

```hcl
resource "local_file" "kubeconfig" {
    content  = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
    filename = pathexpand("~/.kube/test-config")
}
```

What this does, it downloads the `kubeconfig` file from DO to your specified local path.

If you have `kubectl` installed on your machine, then you can now use the config file to connect to your cluster and start deploying your pods and services.

We have just one single issue to handle... **We need a load balancer** to direct outside traffic to our services and pods.

## Deploy a Load Balancer with the Kubernetes TF provider 

There are 2 ways of doing this in Digital Ocean:

- Create a LoadBalancer with Kubernetes and DO will auto generate a LB for you - not ideal as we cannot control it with Terraform

- Create a LoadBalancer directly with Terraform so we have 100% control of it and it's configuration

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

## Add your domain to Digital Ocean with TF

If you have a domain that you want to associate with your new load balancer, then add this block of config code.

```hcl
resource "digitalocean_domain" "mihaiblebea_com" {
    name       = "mihaiblebea.com"
    ip_address = digitalocean_loadbalancer.public.ip
}
```

The `name` is the actual domain that you probably bought from GoDaddy or some other third party provider.

`ip_address` is the ip where we want our domain to point to. If you specify this optional config here, an A record will be generated in the DO domain settings. Pretty sweet, no need to touch that anymore. 😂

## Create a new certificate with TF

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

## Run everything together

Now, save the file, and run the commands again:

1. `terraform plan`

2. `terraform apply -auto-approve`

You can go to your Digital Ocean account and inspect the resources to see if you have the cluster, load balancer, domain and certificate.

Also, remember that we saved the kubernetes config file to our local `.kube` folder. Open a terminal window and type this command `export KUBECONFIG=$HOME/.kube/config-file-name` to use the config file and start working with your kubernetes cluster.

Maybe you noticed that Terraform created a couple of files in your folder. Those are related to the state that TF needs to keep track of so it knows what it needs to create and what already exists in your current infrastructure.

Just leave them be for now, we will go over them in a future article.

## What's next?

This all looks great, but if you plan to use this in production, you will quickly encounter some issues:

- You will need to copy paste the `main.tf` file to create the `dev`, `staging` and `production` environments

- All the secrets and arguments will be hardcoded into the TF file and deployed to Github

- There only way to add new resources to our infrastructure is to extend the `main.tf` file, which will resemble the `node_modules` folder very quickly - we need a better way to do this

In **Part 2** we will add variables, outputs and start using Terraform modules to make our infrastructure as code more modular and to support a stagging and a production environment, without copy-pasting code around.

> There is always a high risk when using Terraform to commit passwords to Github if you are copy-pasting them directly in your `tf` files, so always be very careful about doing so. 

Stay close for part 2, where I will show you how to use Terraform variables to hide those passwords.

<img src="/static/img/home_alone.jpg" />