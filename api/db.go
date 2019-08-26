package api

import "sync"

var TransactionsDB map[string]TransactionRecord = make(map[string]TransactionRecord)

var userAccount Account

var lockObj sync.RWMutex
