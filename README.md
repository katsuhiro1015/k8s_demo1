# k8s_demo1

## Skaffold

### install

```
$ brew install skaffold
```

## Dapr

### install

```
$ brew install dapr/tap/dapr-cli
$ brew link dapr-cli
```

### setup kubernetes

```
$ dapr init -k
```


## Kong

### setup

https://docs.konghq.com/kubernetes-ingress-controller/2.0.x/deployment/minikube/

```
$ helm repo add kong https://charts.konghq.com
$ helm repo update

# Helm 3
$ helm install kong/kong --generate-name --set ingressController.installCRDs=false
```

### manifest

demo1のServiceへのingress作成

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: demo
spec:
  ingressClassName: kong
  rules:
  - http:
      paths:
      - path: /
        pathType: ImplementationSpecific
        backend:
          service:
            name: demo1
            port:
              number: 8080
```

### External IPを設定

```
$ minikube tunnel
```

### 動作確認

```
$ curl http://localhost
```