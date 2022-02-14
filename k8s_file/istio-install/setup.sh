#!/bin/zsh

# istio
kind create cluster --name=istio --config=single-node.yaml

istioctl install -f ./install-istio-arm.yaml -y

kubectl label namespace default istio-injection=enabled --overwrite

# kubectl label namespace default istio-injection=-

# sudo kubectl port-forward -n istio-system service/istio-ingressgateway 80

# cert-manager
# kubectl apply -f https://github.com/jetstack/cert-manager/releases/latest/download/cert-manager.yaml
