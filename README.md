# learn-grpc

learn grpc exposed via Istio using grpc & grpc-web

Goal: use one https port to proxy https, grpc and grpc-web

## Usage

1. `kind create cluster --name istio-testing`
2. `istioctl install --set profile=demo -y`
5. `kubectl -n istio-system create secret tls meera-tech-manual --cert ~/.ssh/meera.tech/fullchain.pem --key ~/.ssh/meera.tech/privkey.pem`
	updated with kubectl -n istio-system create secret tls meera-tech-manual --cert /etc/letsencrypt/live/meera.tech/fullchain.pem --key /etc/letsencrypt/live/meera.tech/privkey.pem -o yaml --dry-run --save-config | kaf -
3. `export KO_DOCKER_REPO=tespkg` `ko apply -f config/`
4. expose ingress gateway in kind: https://kind.sigs.k8s.io/docs/user/loadbalancer/


## Useful debug commands:


```
istioctl proxy-config listeners istio-ingressgateway-76b86f6b45-kk4wj.istio-system
istioctl proxy-config cluster istio-ingressgateway-76b86f6b45-kk4wj --subset hello-grpc.default.svc.cluster.local -o yaml | nvim

istioctl proxy-status
istioctl proxy-status istio-ingressgateway-76b86f6b45-kk4wj.istio-system
istioctl proxy-status istio-ingressgateway-76b86f6b45-kk4wj.istio-system | nvim

istioctl proxy-config listeners istio-ingressgateway-76b86f6b45-kk4wj.istio-system
istioctl proxy-config listeners istio-ingressgateway-76b86f6b45-kk4wj.istio-system --port 8443 -o yaml | nvim
```
