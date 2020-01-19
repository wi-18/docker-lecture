Lösung
======

Beispiel einer `index.html` Datei:
```html
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>An nginx hosted webpage</title>
</head>
<body>
    <h1>Great! NginX seems to be running inside docker :D</h1>
</body>
</html>
```

--- 
In diesem Beispiel befindet sich die `index.html` Datei im Order `Users/DHWB/Documents/Docker/html`.

```bash
docker run -p 80:80 -v Users/DHWB/Documents/Docker/html:/usr/share/nginx/html nginx
```

