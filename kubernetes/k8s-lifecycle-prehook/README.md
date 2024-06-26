### Test

```sh
jmeter -t test.jmx \
    -Jusers=10 \
    -JrampUpPeriod=1 \
    -Jduration=60 \
    -Jpath=/chain \
    -n -l results.jtl
```

### Install Nginx Ingress Controller

1. Community version
2. NGINX version

Diff: https://www.f5.com/company/blog/nginx/guide-to-choosing-ingress-controller-part-4-nginx-ingress-controller-options

#### 1. Community version

```sh
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.10.1
```

##### nginx.ingress.kubernetes.io/service-upstream: "true"

kubectl ingress-nginx backends -n ingress-nginx --backend default-app-8080

```json
{
  "endpoints": [
    {
      "address": "10.43.241.18",
      "port": "8080"
    }
  ],
  "name": "default-app-8080",
  "noServer": false,
  "port": 8080,
  "service": {
    "metadata": {
      "creationTimestamp": null
    },
    "spec": {
      "allocateLoadBalancerNodePorts": true,
      "clusterIP": "10.43.241.18",
      "clusterIPs": [
        "10.43.241.18"
      ],
      "externalTrafficPolicy": "Cluster",
      "internalTrafficPolicy": "Cluster",
      "ipFamilies": [
        "IPv4"
      ],
      "ipFamilyPolicy": "SingleStack",
      "ports": [
        {
          "name": "8080-8080",
          "nodePort": 32392,
          "port": 8080,
          "protocol": "TCP",
          "targetPort": 8080
        }
      ],
      "selector": {
        "app": "app"
      },
      "sessionAffinity": "None",
      "type": "LoadBalancer"
    },
    "status": {
      "loadBalancer": {}
    }
  },
  "sessionAffinityConfig": {
    "cookieSessionAffinity": {
      "name": ""
    },
    "mode": "",
    "name": ""
  },
  "sslPassthrough": false,
  "trafficShapingPolicy": {
    "cookie": "",
    "header": "",
    "headerPattern": "",
    "headerValue": "",
    "weight": 0,
    "weightTotal": 0
  },
  "upstreamHashByConfig": {
    "upstream-hash-by-subset-size": 3
  }
}
```

##### nginx.ingress.kubernetes.io/service-upstream: "false"

kubectl ingress-nginx backends -n ingress-nginx --backend default-app-8080

```json
{
  "endpoints": [
    {
      "address": "10.42.0.58",
      "port": "8080"
    },
    {
      "address": "10.42.0.82",
      "port": "8080"
    }
  ],
  "name": "default-app-8080",
  "noServer": false,
  "port": 8080,
  "service": {
    "metadata": {
      "creationTimestamp": null
    },
    "spec": {
      "allocateLoadBalancerNodePorts": true,
      "clusterIP": "10.43.241.18",
      "clusterIPs": [
        "10.43.241.18"
      ],
      "externalTrafficPolicy": "Cluster",
      "internalTrafficPolicy": "Cluster",
      "ipFamilies": [
        "IPv4"
      ],
      "ipFamilyPolicy": "SingleStack",
      "ports": [
        {
          "name": "8080-8080",
          "nodePort": 32392,
          "port": 8080,
          "protocol": "TCP",
          "targetPort": 8080
        }
      ],
      "selector": {
        "app": "app"
      },
      "sessionAffinity": "None",
      "type": "LoadBalancer"
    },
    "status": {
      "loadBalancer": {}
    }
  },
  "sessionAffinityConfig": {
    "cookieSessionAffinity": {
      "name": ""
    },
    "mode": "",
    "name": ""
  },
  "sslPassthrough": false,
  "trafficShapingPolicy": {
    "cookie": "",
    "header": "",
    "headerPattern": "",
    "headerValue": "",
    "weight": 0,
    "weightTotal": 0
  },
  "upstreamHashByConfig": {
    "upstream-hash-by-subset-size": 3
  }
}
```

#### 2. NGINX version

```sh
helm install my-release oci://ghcr.io/nginxinc/charts/nginx-ingress --version 1.2.2
```

```sh
NAME                                  ENDPOINTS                         AGE
app                                   10.42.0.58:8080,10.42.0.61:8080   3h36m
my-release-nginx-ingress-controller   10.42.0.56:443,10.42.0.56:80      29h
kubernetes                            192.168.8.101:6443                121d
```


Nginx Conf
```conf
# configuration for default/app
upstream default-app-foo.com-app-8080 {zone default-app-foo.com-app-8080 256k;
	random two least_conn;
	server 10.42.0.58:8080 max_fails=1 fail_timeout=10s max_conns=0;
	server 10.42.0.61:8080 max_fails=1 fail_timeout=10s max_conns=0;
}



server {
	listen 80;listen [::]:80;

	server_tokens on;

	server_name foo.com;

	set $resource_type "ingress";
	set $resource_name "app";
	set $resource_namespace "default";
	location / {
		set $service "app";
		proxy_http_version 1.1;

		proxy_connect_timeout 60s;
		proxy_read_timeout 60s;
		proxy_send_timeout 60s;
		client_max_body_size 1m;
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Host $host;
		proxy_set_header X-Forwarded-Port $server_port;
		proxy_set_header X-Forwarded-Proto $scheme;
		proxy_buffering on;
		proxy_pass http://default-app-foo.com-app-8080;
	}
}
```
