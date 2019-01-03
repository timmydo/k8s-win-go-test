# k8s-win-go-test

This is a windows test app for testing windows containers on kubernetes.


```
docker build -t timmydo/k8s-win-go-test:win-latest -f Dockerfile.win-1809 .
kubectl apply -f https://raw.githubusercontent.com/timmydo/k8s-win-go-test/master/test.yaml
```