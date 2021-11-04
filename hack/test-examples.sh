#!/usr/bin/env bash
set -eu -o pipefail

# Load the configmaps that contains the parameter values used for certain examples.
kubectl apply -f examples/configmaps/simple-parameters-configmap.yaml

grep -lR 'workflows.argoproj.io/test' examples/* | while read f ; do
  kubectl delete workflow -l workflows.argoproj.io/test
  echo "RUN $f"
  kubectl create -f $f
  name=$(kubectl get workflow -o name)
  kubectl wait --for=condition=Completed $name
  phase="$(kubectl get $name -o 'jsonpath={.status.phase}')"
  test Succeeded == $phase
done
