#!/bin/bash

kubectl apply -f deployment_tg.yaml -n tienkang
kubectl apply -f service.yaml -n tienkang
kubectl apply -f ingress_tg.yaml -n tienkang
