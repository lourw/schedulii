# This script requires the following to be installed on your machine:
# - docker
# - kind
# - kubectl

# Delete cluster if it already exists
kind delete cluster

# Create cluster with configurations that allow nginx ingress to run
# https://kind.sigs.k8s.io/docs/user/ingress/
cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF

# Load local docker image into cluster control plane
kind load docker-image schedulii_api
kind load docker-image schedulii-ui

# Apply local kubernetes manifest to deploy app resources (deployment, service, ingress)
kubectl apply -f schedulii_api.yaml
kubectl apply -f schedulii-ui.yaml

# Apply nginx ingress controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for ingress controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

# Before you can access the ingress, you need to add an entry to your /etc/hosts file
# Example: append `127.0.0.1 schedulii.api.dev.com` to /etc/hosts
