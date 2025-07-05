## Install Nginx Ingress Controller and Other Resources

### Option 1: Using Kustomize

#### Install

```sh
kubectl kustomize . --enable-helm | kubectl apply -f -
```

```sh
kubectl wait --timeout=5m -n ingress-nginx deploy/nginx-ic-ingress-nginx-controller --for=condition=Available
```

#### Uninstall

```sh
kubectl kustomize . --enable-helm | kubectl delete -f -
```

### Option 2: Using Helm

```sh
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.10.1 \
  -f values.yaml
```

## Check Logs

### Controller Logs

```sh
kubectl logs -n ingress-nginx -l app.kubernetes.io/component=controller -f
```

### View Nginx Configs

```sh
kubectl ingress-nginx --deployment nginx-ic-ingress-nginx-controller -n ingress-nginx conf
```

### View Backends

```sh
k ingress-nginx --deployment nginx-ic-ingress-nginx-controller -n ingress-nginx backends
```

## Test

```sh
curl 'http://localhost/foo/bar' -H 'Host: foo.com' -d 'hello world!'
```

```sh
curl 'http://localhost/hello' -H 'host: example.com'
```

```sh
curl 'http://localhost/lua' -H 'host: example.com'
```

## Nginx Configs

### Nginx ConfigMap
https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#configuration-options


### Helm Chart Configs
https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml
