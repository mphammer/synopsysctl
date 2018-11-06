#!/bin/bash

unset DYLD_INSERT_LIBRARIES
echo "args = Namespace, Reg_key, branch"

NS=$1
REG_KEY=$2
BRANCH=$3

echo "Using the secret encoded in this file.  Change it before running, or press enter..."
read x

cat << EOF > /tmp/secret
apiVersion: v1
data:
  ADMIN_PASSWORD: YmxhY2tkdWNr
  POSTGRES_PASSWORD: YmxhY2tkdWNr
  USER_PASSWORD: YmxhY2tkdWNr
  HUB_PASSWORD: YmxhY2tkdWNr
kind: Secret
metadata:
  name: blackduck-secret
type: Opaque
EOF

oc new-project $NS

oc create -f /tmp/secret -n $NS

cat ../blackduck-operator.yaml | sed 's/${REGISTRATION_KEY}/'$REG_KEY'/g' | sed 's/${NAMESPACE}/'$NS'/g' | sed 's/${BRANCHNAME}/'${BRANCH}'/g' | oc create --namespace=$NS -f -

#oc expose rc blackduck-operator --port=8080 --target-port=8080 --name=blackduck-operator-np --type=NodePort --namespace=$NS

#oc expose rc blackduck-operator --port=8080 --target-port=8080 --name=blackduck-operator-lb --type=LoadBalancer --namespace=$NS
