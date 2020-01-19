Volumes mit NginX lernen
=====

Das Ziel
========
Das Ziel dieser Übung ist es einen Nginx Instanz in einem Container zu starten um eigene HTML Seiten bereit zu stellen.
Hierfür muss das Docker Image `nginx` genutzt werden.
Die Dokumentation dieses Images findet sich [hier](https://hub.docker.com/_/nginx).

Das Image nutzt Port 80 um die HTML Inhalte zu publizieren sowie den internen Ordern `/usr/share/nginx/html` um die
Struktur der Webseite abzubilden.

Deine Aufgabe ist es nun eine eigene kleine HTML Seite zu schreiben und diese dann, durch das nutzten von Docker Bind 
Mounts, die selbst erstellte HTML Seite in dem Container Ordner `/usr/share/nginx/html` zu linken.

Hints
=====

Der Nginx Container nutzt Port 80, dementsprechend muss, wie in Übung 3, der Port extern freigegeben werden.

---

```bash
docker run -v path/to/host/content/:/path/to/container/content ...
```
kann der Order `path/to/host/content/` des Host Systems mit dem Order `/path/to/container/content` des Containers
verknüpft werden.

