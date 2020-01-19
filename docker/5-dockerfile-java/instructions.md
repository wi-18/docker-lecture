Erstelle eine Dockerfile für ein Java Applikation
=====

Das Ziel
========
Das Ziel dieser Übung ist es, die Grundlagen einer Dockerfile zu lehren.
Deine Aufgabe ist es eine eigene Java Applikation zu schreiben welche ein einfaches `Hello World` in der Konsole ausgibt
und diese Applikation in ein Docker Image zu verpacken.

Als Basis Image sollte `adoptopenjdk/openjdk12` genutzt werden.

1. Mit `FROM <IMAGE>` wird das Basis Image spezifiziert.
2. Außerdem sollte eine `working directory`, also ein Order als Basis aller weiteren Operationen spezifiziert werden.
Dies kann durch `WORKDIR <PATH>` erreicht werden.
3. Als nächster Schritt muss die entwickelte `.java` Datei in den Container kopiert werden.
Hierfür kann das Kommando `COPY <HOST_PATH> <CONTAINER_PATH>` genutzt werden.
4. Zusätzlich muss der Java Quelltext kompiliert werden. Mit `RUN <COMMAND...>` können Kommandos innerhalb des Containers
ausgeführt werden.
5. Als letzter Schritt muss das Startkommando des Containers festgelegt werden. Dieses Kommando wird ausgeführt wenn der
Container startet. Hier sollte also die Java Datei ausgeführt werden. 

Hints
=====

Um eine Java Datei zu kompilieren kann das `javac YourFileName.java` Kommando genutzt werden.

---

Die kompilierte Java Datei kann mit dem Kommando `java YourFileName` ausgeführt werden. Hierbei wird das Dateiende nicht
übergeben.

