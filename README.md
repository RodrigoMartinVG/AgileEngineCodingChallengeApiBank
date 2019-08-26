# Agile Engine - CodingChallenge - Bank API

AgileEngine - Coding Challenge for Paxos - Bank API
## GO Version

```
go1.11.5 linux/amd64
```
## APIs
#### POST /transactions
Inserts a new transaction. Type field may be "credit" or "debit" literals only.
Returns the inserted transaction ID
```json
{
    "type": "credit",
    "amount": 123.5,
}
```

#### GET /transactions
Returns an array with all transactions
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
Returns a transaction for the given ID
```json
{
    "id": "0e3d92c3-3441-4415-b21e-08c04ed1ca2a",
    "type": "credit",
    "amount": 123.5,
    "date": "2019-08-24T12:47:23.959745699Z",
}
```
#### GET /
Returns the current balance for the account
```json
{
    "currentAccountBalance": 123.5,
 
}
```

#### GET /ui
Returns history HTML interface for all transactions
