#!/bin/bash

# Script to deploy Calico CNI on Kubernetes cluster
# Usage: ./deploy_calico.sh <cluster-name>

# Set the default cluster name if not provided
CLUSTER_NAME=${1:-my-cluster}

# Function to deploy Calico CNI
deploy_calico() {
    kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.25.0/manifests/calico.yaml
}

# Main execution
echo "Deploying Calico CNI on the cluster: $CLUSTER_NAME"
deploy_calico

echo "Calico CNI deployment completed."
