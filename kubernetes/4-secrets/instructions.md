Erstelle eine Applikation mit Secrets
=====

Das Ziel
========
Das Ziel dieser Übung ist das Erzeugen einer Applikation welche Kubernetes Secrets nutzt um Daten zu erhalten.
Erstelle hierfür ein Secret mit den Daten `username` und `password`.
Die Werte für die Daten des Secrets müssen mit Base64 encoded werden.

Erstelle danach ein Deployment auf Basis des Images `wi18/kubernetes-secret-app`.
Das erstellte Secret muss in den Order `/etc/app/config` verlinkt werden, damit die App die Daten 
`/etc/app/config/username` und `/etc/app/config/password` lesen kann.

Mit `kubectl logs deployment.apps/<name-of-the-deployment>` kann geprüft werden, ob die Applikation die Secrets korrekt 
ausgelesen hat.

Hints
=====
Provide a simple hint to ease execution of the task for students.

