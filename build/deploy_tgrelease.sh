#!/bin/bash

kubectl apply -f deployment_tgrelease.yaml
kubectl apply -f service.yaml 
kubectl apply -f ingress_tgrelease.yaml 
