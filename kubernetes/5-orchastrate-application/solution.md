Lösung
======

Das Backend kann durch folgende Datei erzeugt werden:
```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: orch-backend-svc
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: 80
      protocol: "TCP"
  selector:
    app: orch-backend-app
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: orch-backend-deployment
spec:
  selector:
    matchLabels:
      app: orch-backend-app
  replicas: 2
  template:
    metadata:
      labels:
        app: orch-backend-app
    spec:
      containers:
        - name: orch-backend-container
          image: wi18/kubernetes-orch-backend
          ports:
            - containerPort: 80
```

--- 

Das Frontend danach durch:
```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: orch-frontend-svc
spec:
  type: NodePort
  ports:
    - port: 80
      nodePort: 30000
      protocol: "TCP"
  selector:
    app: orch-frontend-app
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: orch-frontend-deployment
spec:
  selector:
    matchLabels:
      app: orch-frontend-app
  replicas: 1
  template:
    metadata:
      labels:
        app: orch-frontend-app
    spec:
      containers:
        - name: orch-frontend-container
          image: wi18/kubernetes-orch-frontend
          ports:
            - containerPort: 80
```
