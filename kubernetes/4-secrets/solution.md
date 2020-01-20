Lösung
======

Das Secret kann mit folgender YAML Datei erstellt werden:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: config-secret
type: Opaque
data:
  username: YWRtaW4= #admin in base64
  password: cGFzc3dvcmQxMjM= #password123 in base64
```

um im Folgenden von diesem Deployment genutzt werden zu können:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: secret-deployment
spec:
  selector:
    matchLabels:
      app: secret-deployment-app
  template:
    metadata:
      labels:
        app: secret-deployment-app
    spec:
      containers:
        - name: secret-deployment-container
          image: wi18/kubernetes-secret-app
          volumeMounts:
            - mountPath: "/etc/app/config"
              name: config-secret-volume
              readOnly: true
      volumes:
        - name: config-secret-volume
          secret:
            secretName: config-secret
```
