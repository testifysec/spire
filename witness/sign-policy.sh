#/bin/sh

spiffeid="spiffe://dev.testifysec.com/witness-demo/builder"

if which yq > /dev/null; then
    yq --version
else
    echo "installing yq - need sudo"
    VERSION=v4.2.0 BINARY=yq_linux_amd64

    sudo wget https://github.com/mikefarah/yq/releases/download/${VERSION}/${BINARY} -O /usr/bin/yq &&\
    sudo chmod +x /usr/bin/yq
fi

#get trust bundle
kubectl -n=spire get cm spire-bundle -o=json | jq -r '.data["bundle.crt"]' > bundle.crt.tmp
spireca=$(cat bundle.crt.tmp)
fulcioca=$(cat fulcioca.pem)


gcp_b64="$(openssl base64 -A <"gcp.rego")"
gitlab_b64="$(openssl base64 -A <"gitlab.rego")"


bundleb64=$(echo -n "${spireca}" | base64)
fulciob64=$(echo -n "${fulcioca}" | base64)

#gen commit gate policy
cp ci-gate-policy-tmpl.yaml policy.tmp.yaml



fulciorootid=`cat fulcioca.pem | openssl x509 -outform der | sha256sum | awk '{print $1}'`
sed -i "s/{{FULCIO_ROOT_ID}}/$fulciorootid/g" policy.tmp.yaml

yq eval ".roots.${fulciorootid}.certificate = \"${fulciob64}\"" --inplace policy.tmp.yaml



#gen artifact policy
cp artifact-policy-tmpl.yaml policy.tmp.yaml
spirerootid=`cat bundle.crt.tmp | openssl x509 -outform der | sha256sum | awk '{print $1}'`
sed -i "s/{{ROOT_ID}}/$spirerootid/g" policy.tmp.yaml
sed -i "s#{{SPIFFEID}}#$spiffeid#g" policy.tmp.yaml
sed -i "s#{{GCP_IIT_MODULE}}#$gcp_b64#g" policy.tmp.yaml
sed -i "s#{{GITLAB_MODULE}}#$gitlab_b64#g" policy.tmp.yaml
sed -i "s/{{FULCIO_ROOT_ID}}/$fulciorootid/g" policy.tmp.yaml

yq eval ".roots.${spirerootid}.certificate = \"${bundleb64}\"" --inplace policy.tmp.yaml
yq eval ".roots.${fulciorootid}.certificate = \"${fulciob64}\"" --inplace policy.tmp.yaml



#sed -i "s/{{B64_POLICY_MODULE}}/$rego_b64/g" policy.tmp.json

#norm
yq e -j policy.tmp.yaml > policy.tmp.json

#sign policy with witness
witness sign -k testkey.pem -f policy.tmp.json -o artifact-policy.signed.json
