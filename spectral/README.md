### Spectral

```sh
spectral lint basic-swager.yaml
```

```log
  2:6  warning  info-contact           Info object must have "contact" object.      info
 11:9  warning  operation-operationId  Operation must have "operationId".           paths./users.get
 11:9  warning  operation-tags         Operation must have non-empty "tags" array.  paths./users.get
 24:9  warning  operation-operationId  Operation must have "operationId".           paths./bar.get
 24:9  warning  operation-tags         Operation must have non-empty "tags" array.  paths./bar.get
```
