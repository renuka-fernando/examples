### Install Nginx Ingress Controller

```sh
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.10.1 \
  -f values.yaml
```

### Nginx Configs

https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#configuration-options

https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml

### Apply Configs

```sh
k apply -f k8s/
```
### Test

```sh
curl http://localhost/foo -H 'Host: foo.com' -v
```

