Erstelle und skaliere eine erste Applikation in Kubernetes
=====

Das Ziel
========
Das Ziel dieser Übung ist es eine einfache Applikation in Kubernetes zu erstellen und dieses Deployment im nachhinein 
zu skalieren.

Erstelle ein Deployment auf Basis des Images `wi18/docker-web-page` mit dem Namen `simple-deployment`.
Das Deployment Template bekommt außerdem das Label `app: simple-deployment-app` um in weiteren Aufgaben identifiziert 
werden zu können.

Um zu validieren, dass die Applikation funktionstüchtig ist, können die Logs des Deployments betrachtet werden.
Hierfür kann `kubectl logs deployment.apps/simple-deployment` genutzt werden.

Hints
=====

Mit
```bash
kubectl scale deployment/your_deployment_name --replicas x
```
kann ein Deployment skaliert werden.

