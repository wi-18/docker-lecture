Lösung
======

Das Deployment kann mit der YAML Datei
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: simple-deployment
spec:
  selector:
    matchLabels:
      app: simple-deployment-pod
  replicas: 1
  template:
    metadata:
      labels:
        app: simple-deployment-pod
    spec:
      containers:
        - name: simple-deployment-container
          image: wi18/docker-web-page
          ports:
            - containerPort: 80
```
erstellt werden.

----

Dieses kann im Nachhinein mit 
```
kubectl scale deployment/simple-deployment --replicas 10
```
auf 10 Instanzen hoch skaliert werden. 
