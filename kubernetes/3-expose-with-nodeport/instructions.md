Veröffentliche eine Applikation mit einem NodePort Service
=====

Das Ziel
========
Das Ziel dieser Übung ist es die Applikation aus Übung 2, basierend auf dem Image `wi18/docker-web-page` auf einem 
externen Port freizugeben.
Hierfür muss ein Service angelegt werden, ein NodePort Service.
Ein NodePort kann beliebige Ports von anderen Kubernetes Objekten auf eine externe IP Adresse weiterleiten.
Allerdings können extern nur die Ports 30000-32767 genutzt werden. 

Hints
=====

In der YAML kann ein Service unter `specs.type` mit dem Wert `NodePort` als NodePort Service spezifiziert werden.

---

Anstatt eines `targetPort` kann in einem NodePort Service ein `nodePort` spezifiziert werden.

