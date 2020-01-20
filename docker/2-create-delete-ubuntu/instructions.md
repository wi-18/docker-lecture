Erstelle und lösche einen Container
=====

Das Ziel
========
In dieser Übung sollst du einen neuen Container, auf der Basis des Ubuntu Images aus Übung 1, erstellen und wieder 
löschen.

Erstelle den Container
----------------------

Erstelle einen neuen Container auf Basis des Ubuntu Images.

Nutze hierfür den Command [docker run](https://docs.docker.com/engine/reference/run/).
Um direkt im Container arbeiten zu können, müssen noch die Flags `-i` und `-t` spezifiziert werden.

`-t` öffnet eine Terminal Session in dem Container, mit welcher interagiert werden kann.
`-i` erhält den Container am Leben solange wie Nutzer Eingaben erstellen. 

Sobald der Container läuft kannst du den Command `uname -s` nutzten um den Namen des Betriebssystem herauszufinden.
Dieser sollte `Linux` zurückgeben.

Danach kannst du mit `exit` deine Terminal Session beenden, was gleichzeitig den Container stoppt.

Lösche den Container
====================

Um den Container zu löschen müssen wir zuerst die **ID** des Containers heraus finden.
Hierzu können wir den:
````bash
docker ps
````
Command nutzten.
Da der Container allerdings bereits gestoppt ist wird dieser nicht mehr aufgezeigt.
Um all existierenden Container aufzuzeigen muss die `--all` Flag gesetzt werden.

Der Output des Commands sollte folgendermaßen aussehen:
```
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                     PORTS               NAMES
ba2a99c85e7d        ubuntu              "/bin/bash"         x seconds ago       Exited (0) x seconds ago                       admiring_visvesvaraya
```

Da wir keinen Namen spezifiziert haben hat Docker hat dem Container einen eigenen gegeben.
In diesem Beispiel heißt der Container `admiring_visvesvaraya`.
Nun lösche diesen Container mit Hilfe des:
```bash
docker rm
```` 
Commands.

Hints
=====
```bash
docker run -it IMAGE_NAME
```
startet einen neuen Container auf der Basis des spezifizierten Images und erlaubt es in diesem mit einem Terminal zu 
arbeiten.

----

```bash
docker ps --all
```
zeigt alle Container an, welche auf dem Host existieren.

----

```bash
docker rm CONTAINER_NAME
```
löscht den Container mit dem spezifizierten Namen.

