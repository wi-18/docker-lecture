Erstelle und skaliere eine erste Applikation in Kubernetes
=====

Das Ziel
========
Das Ziel dieser Übung ist es eine einfache Applikation in Kubernetes zu erstellen und dieses Deployment im nachhinein 
zu skalieren.


Hints
=====
```bash
kubectl scale --replicas=5 deployment/name-of-deployment
```
würde das Deployment mit dem Namen `name-of-deployment` auf 5 Replicas skalieren.

