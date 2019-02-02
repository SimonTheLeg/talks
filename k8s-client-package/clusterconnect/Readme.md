# Building the image yourself

make sure you are in the `k8s-client-package` folder (so go.mod and go.sum are in the Docker build context) and run:

```sh
docker build . -t <your-creative-tag> -f clusterconnect/Dockerfile
```

# Deploying the image to a k8s cluster

## Without auth (default service account)

```sh
kubectl apply -f without-auth
```

## With auth (creates service account, role, and rolebinding)

```sh
kubectl apply -f with-auth
```