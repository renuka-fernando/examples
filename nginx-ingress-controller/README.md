### Install Nginx Ingress Controller and Other Resources

```sh
kubectl kustomize . --enable-helm | kubectl apply -f -
```

### Controller Logs

```sh
kubectl logs -n ingress-nginx -l app.kubernetes.io/component=controller -f
```

### Test

```sh
curl 'http://localhost/hello' -H 'host: example.com'
```

```sh
curl 'http://localhost/lua' -H 'host: example.com'
```

### Uninstall

```sh
kubectl kustomize . --enable-helm | kubectl delete -f -
```

#### Install Nginx Ingress Controller using Helm

```sh
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.10.1 \
  -f values.yaml
```

### Nginx Configs

#### Nginx ConfigMap
https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#configuration-options


#### Helm Chart Configs
https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml
