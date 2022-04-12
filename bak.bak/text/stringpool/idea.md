# stringpool

> Tricky and fun utilities for Go programs on macOS.

## Goals

-   maintain a collection of strings.Builder objects
-   return a Builder object when requested
-   release a Builder object when work is complete
-   lower the workload on the GC
