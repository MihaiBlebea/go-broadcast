---
apiVersion: v1
kind: Service
metadata:
  name: tcp-loadbalancer
  annotations:
    # https://developers.digitalocean.com/documentation/v2/#load-balancers
    # https://www.digitalocean.com/docs/kubernetes/how-to/configure-load-balancers/
    service.beta.kubernetes.io/do-loadbalancer-name: "my.example.com"
    service.beta.kubernetes.io/do-loadbalancer-hostname: "my.example.com"
    service.beta.kubernetes.io/do-loadbalancer-protocol: "tcp"
    service.beta.kubernetes.io/do-loadbalancer-tag: "k8s-my-worker"     # remember tag your droplet !!!
    service.beta.kubernetes.io/do-loadbalancer-algorithm: "round_robin" # options: round_robin, least_connections
    service.beta.kubernetes.io/do-loadbalancer-tls-ports: "443"
    service.beta.kubernetes.io/do-loadbalancer-tls-passthrough: "true"
    service.beta.kubernetes.io/do-loadbalancer-enable-proxy-protocol: "true"
    # service.beta.kubernetes.io/do-loadbalancer-certificate-id: "your-certificate-id"
spec:
  type: LoadBalancer
  selector:
    app: traefik
  ports:
    - name: http
      protocol: TCP
      port: 80
      targetPort: 8000
    - name: https
      protocol: TCP
      port: 443
      targetPort: 4443
    - name: postgres-tcp
      protocol: TCP
      port: 5432
      targetPort: 25432
    - name: postgres-adapter-http
      protocol: TCP
      port: 9201
      targetPort: 29201
    - name: traefik-http
      protocol: TCP
      port: 8090
      targetPort: 8090


https://goobar.io/2019/12/07/manually-trigger-a-github-actions-workflow/