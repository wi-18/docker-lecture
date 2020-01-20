Lösung
======

```yaml
apiVersion: v1
kind: Service
metadata:
  name: simple-deployment-nodeport-service
spec:
  type: NodePort
  ports:
    - port: 80
      nodePort: 30000
      protocol: "TCP"
  selector:
    app: simple-deployment-app
```
