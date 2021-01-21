#!/bin/bash

kubectl apply -f deployment_dev.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress_dev.yaml
