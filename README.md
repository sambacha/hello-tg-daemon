## Building

Build Command:

```
go build
```

Output file: `hello-tg-daemon`

## Running

Run with

```
./hello-tg-daemon --http.api=myNamespace
```

**Note: remember to run turbo-geth with --private.api.addr=localhost:9090**

```
tg --private.api.addr=localhost:9090
```

## Test

```
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"myNamespace_getBlockNumberByHash","params":["ANYHASH"],"id":1}' localhost:8545
```

replace ANYHASH with a blockhash accordingly to the chain being used.