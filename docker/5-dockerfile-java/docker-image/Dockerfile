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
