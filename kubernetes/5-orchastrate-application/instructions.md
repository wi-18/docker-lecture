Erstelle ein zweiteilige Applikation in Kubernetes
=====

Das Ziel
========
In dieser Übung soll eine Applikation, welche aus zwei Services besteht, erstellt werden.
Die Applikation ist aufgeteilt in das Image `wi18/kubernetes-orch-backend`, welches auf Port `80` Daten bereitstellt
und das Image `wi18/kubernetes-orch-frontend`, welches auf Port `80` eine Webseite bereitstellt welche die Daten des Backends
visualisiert.

Erstelle zuerst das Backend Deployment sowie einen einfachen Service für dieses.
Dieser ist kein `NodePort` sondern ein `ClusterIP` Service und spezifiziert dementsprechend einen `targetPort` und 
keinen `nodePort`. Der Wert hierbei ist Port `80`.

---

**Nenne den Service des Backends `orch-backend-svc`!!!!!!**

Danach erstelle, wie in Aufgabe 3, ein Deployment für das Frontend (Image: `wi18/kubernetes-orch-frontend`) sowie 
einen `NodePort` Service für dieses Deployment welches Port `80` auf Port `30000` veröffentlicht.
Die Namensgebung des Frontends ist irrelevant, kann aber der `orch-frontend-` Strategie folgen.

Hints
=====
Good luck.

