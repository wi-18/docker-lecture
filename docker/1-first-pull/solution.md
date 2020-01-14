Lösung
======

Das Image kann durch:
````bash
docker pull ubuntu:latest
````
auf die lokale Maschinen heruntergeladen werden.
Der ``:latest`` Tag nach dem Namen des Images (`ubuntu`), signalisiert Docker die neuste Version des Images 
herunterzuladen.
Dieser Tag ist allerdings der Standard und kann so entfallen.
Eine andere Lösung wäre dementsprechend: 
```bash
docker pull ubuntu
```
