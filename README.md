# learn-grpc

learn grpc exposed via Istio using grpc & grpc-web

## Usage

1. `kind create cluster --name istio-testing`
2. `istioctl install --set profile=demo -y`
5. `kubectl -n istio-system create secret tls meera-tech-manual --cert ~/.ssh/meera.tech/fullchain.pem --key ~/.ssh/meera.tech/privkey.pem`
	updated with kubectl -n istio-system create secret tls meera-tech-manual --cert /etc/letsencrypt/live/meera.tech/fullchain.pem --key /etc/letsencrypt/live/meera.tech/privkey.pem -o yaml --dry-run --save-config | kaf -
3. `export KO_DOCKER_REPO=tespkg` `ko apply -f config/`
4. expose ingress gateway in kind: https://kind.sigs.k8s.io/docs/user/loadbalancer/
