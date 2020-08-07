---
Title: How to build a self hosted CI/CD pipeline with Go from scratch
Summary: We are going to build a CI/CD pipeline from scratch. It's probably not going to be production ready, but will work for most individual developers
Image: https://img.pngio.com/cartoon-pipeline-cartoon-clipart-assembly-line-cartoon-creative-assembly-line-png-650_400.png
Tags:
    - ci
    - cd
    - golang
    - hosted
Layout: article
Slug: ci-cd-pipeline-with-golang
Published: "2020-08-07 15:04:05"
Kind: article
---
## Requirements
- [x] run on a different machine, server
- [] clone code from github on every push to master branch
- [] run the unit tests against the code
- [] build the docker container / containers and push the images to dockerhub
- [] use terraform to deploy the code to digital ocean
- [] have all the above steps into a modular configuration file