## Chainlink external adapter starter
 


###clean artefacts
```
make clean
```

### buil all and start server

```
make all
```

### run some go tests

```
make test
```

### simulate adapter http call
```
curl -X POST -H 'Content-Type: application/json' -d '{}' http://localhost:9999/tick
```

### or simply
```make```
