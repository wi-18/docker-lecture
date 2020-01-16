Lösung
======
```bash
docker run -p 80:80 wi18/docker-web-page
```
startet den Container und leitet den Port 80 des Containers auf den Port 80 des Hostsystems um.

---
```bash
docker run -e PORT=8080 -p 80:8080 wi18/docker-web-page
```
startet den Container, setzt die Environment Variable `PORT` auf `8080` und leitet den Port `8080` des Containers auf
den Port `80` des Hosts um.
