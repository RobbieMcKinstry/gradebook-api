#!/bin/bash

export KUBECONFIG="${KUBECONFIG}:$(pwd)/.deploy/kubeconfig"
kubectl config use-context vagrant-single
