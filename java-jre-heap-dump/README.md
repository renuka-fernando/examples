### Build

```shell
mvn clean package
```

### Get Heap Dump

```shell
java -Dcom.sun.management.jmxremote \
    -Dcom.sun.management.jmxremote.port=12345 \
    -Dcom.sun.management.jmxremote.authenticate=false \
    -Dcom.sun.management.jmxremote.ssl=false \
    -jar target/java-jre-heap-dump-1.0-SNAPSHOT.jar
```

### Visualize

Open VisualVM
