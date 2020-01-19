Lösung
======

Als Beispiel Java Datei dient hier:
```java
public class Example {
    public static void main(String[] args) {
        System.out.println("Hello World from inside a docker container!!");
    }
}
```

Die Dockerfile ist folgenderweise aufgebaut:
```dockerfile
# The base image for this docker image will be openjdk12, which contains java 12
FROM adoptopenjdk/openjdk12

# The /app/src directory is our working directory
WORKDIR /app/src

# Copy the entire content of the /src folder (the Example.java file) into the containers working directory (/app/src)
COPY src/ .

# Compile the Example.java file into an Example.class file inside the Image.
RUN javac Example.java

# Specify the start command in which the java file is executed. 
CMD java Example
```

Diese Dateien liegen in folgender Dateistruktur vor:
```
.
├── Dockerfile
└── src
    └── Example.java

1 directory, 2 files
```
