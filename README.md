# LRUCache Package

LRUCache is a simple Go package that provides an implementation of a Least Recently Used (LRU) cache.

## Overview

The LRUCache package efficiently manages a fixed-size cache using the LRU eviction policy. It allows users to store key-value pairs and retrieve values based on keys.

## Installation

To use the LRUCache package in your Go project, import it using the following import path:

```go
import (
    lrucache "github.com/Jaki02003/LRU_CACHE"
)
```

## Usage

### Creating a New LRUCache
``` go
cacheSize := 100
lruCache := lrucache.NewLRUCache(cacheSize)
```

### Adding Key-Value Pairs to the Cache
``` go
lruCache.Put("key1", "value1")
lruCache.Put("key2", "value2")
```

### Retrieving Values from the Cache
``` go
value, exists := lruCache.Get("key1")
if exists {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Key not found")
}
```

### Getting the Current Size of the Cache
``` go
currentSize := lruCache.Len()
fmt.Println("Current Size:", currentSize)
```
## Contributing

Contributions are welcome! Feel free to open issues, propose enhancements, or submit pull requests.
