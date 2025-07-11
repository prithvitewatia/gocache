# GoCache 🧠

`gocache` is a minimal, thread-safe in-memory cache library for Go, built with TTL support and background cleanup.

## Features

- ✅ Set/Get/Delete with optional TTL
- 🧹 Background goroutine for expired key cleanup
- 🔐 Thread-safe with `sync.RWMutex`
- 🧪 Easily extensible for LRU, metrics, persistence

## Installation

```bash
go get github.com/prithvitewatia/gocache
```
