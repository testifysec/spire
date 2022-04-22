#! /bin/bash



nodeIDs=($(kubectl -n spire exec -it spire-server-0 -- /opt/spire/bin/spire-server agent list | grep 'spiffe://dev.testifysec.com/spire/agent/gcp_iit' | cut -f 2- -d ':' | tr -d ' ' | tr -d '\r'))
spiffeid="spiffe://dev.testifysec.com/witness-demo/builder"

#add entry for each nodeID
for node in ${nodeIDs[@]}; do
set +e
kubectl exec -n spire spire-server-0 -- \
    /opt/spire/bin/spire-server entry create \
    -parentID ${node} \
    -spiffeID ${spiffeid} \
    -selector k8s:ns:gitlab-runner
set -e
done
