Publishing ports and setting environment variables
=====

Das Ziel
========
In dieser Übung geht es darum den Port eines Containers in das Netzwerk des Hosts zu publishen sowie einem Container 
Environment Variablen mitzugeben.

Für den Container nutzen wir das folgende Image: `wi18/docker-web-page`.
Das Image published eine einfache Website im Container.
Deine Aufgabe ist es nun einen Container auf Basis dieses Images zu starten, welcher den Port 80 des Containers auf 
Port 80 des Hosts, also dein System, veröffentlicht.
Du kannst kontrollieren ob das Netzwerk des Containers richtig konfiguriert ist, indem du versuchst die Webseite 
aufzurufen.
Diese sollte under `localhost` in einem Browser deiner Wahl zur Verfügung stehen.  

--- 

Mit dem Command 
```bash
docker run -e <key>=<value> ... 
``` 
können Environment Variablen gesetzt werden. 
Als Beispiel wird im folgenden Command die Variable `FOO` auf den Wert `bar` gesetzt.
```bash
docker run -e FOO=bar ...
```

Starte einen neuen Container auf dem gleich Image (`wi18/docker-web-page`) und konfiguriere den Port der Applikation
durch das setzten der Environment Variable **`PORT`** auf den Port `8080`.

Denke daran den neuen Container Port auf den Port `80` des Hostsystemes umzulegen. 
Auch hier kannst du erneut unter `localhost` im Browser validieren ob das Containernetzwerk richtig aufgesetzt worden 
ist.

Hints
=====
In den vorherigen Übungen wurde gezeigt wie Container gestartet werden, falls du hiermit Schwierigkeiten hast, greife 
auf diese zurück.

Der Flag `-p <hostPort>:<containerPort>` wird genutzt um den spezifischen Port des Containers auf einen Hostport 
weiterzuleiten. 


