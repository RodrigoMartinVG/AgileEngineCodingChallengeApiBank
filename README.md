# Agile Engine - CodingChallenge - Bank API

AgileEngine - Coding Challenge for Paxos - Bank API

## Instalation

```
cd $GO_HOME/src
mkdir -p github.com/rodrigomartinvg
cd github.com/rodrigomartinvg
git clone https://github.com/RodrigoMartinVG/AgileEngineCodingChallengeApiBank.git
```

## Version GO

```
go version go1.11.5 linux/amd64
```

## compilation (version Docker)
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o agileEngineCodingChallengeApiBank *.go
```

## If execute on local
```
-listenPort=8080 
```

## APIs

#### Swagger API Documnetation

https://agileengine.bitbucket.io/fsNDJmGOAwqCpzZx/api/#/transactions/transactionsHistory

#### GET /transactions
return a Array of transaction  with all transactions
```json
[
    {
        "id": "0e3d92c3-3441-4415-b21e-08c04ed1ca2a",
        "type": "credit",
        "amount": 123.5,
        "date": "2019-08-24T12:47:23.959745699Z",
    },
    {
        "id": "0e3d92c3-3441-4415-b21e-08c04ed1ca2a",
        "type": "debit",
        "amount": 123.5,
        "date": "2019-08-24T12:47:23.959745699Z",
    }
]
```

#### GET /transactions/{id}
Return a transaction of ID={id}
```json
{
    "id": "0e3d92c3-3441-4415-b21e-08c04ed1ca2a",
    "type": "credit",
    "amount": 123.5,
    "date": "2019-08-24T12:47:23.959745699Z",
}
```

#### POST /transactions
Generate a new transaction
```json
{
    "type": "credit",
    "amount": 123.5,
}
```

#### GET /transactions/{id}
Return a transaction of ID={id}
```json
{
    "id": "0e3d92c3-3441-4415-b21e-08c04ed1ca2a",
    "type": "credit",
    "amount": 123.5,
    "date": "2019-08-24T12:47:23.959745699Z",
}
```

#### GET /
Return balance of unique single account
```json
{
    "currentAccountBalance": 123.5,
 
}
```

#### GET /ui
Return HTML interface with transactions history

