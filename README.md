# backoff

## Usage

```Go
import "github.com/christianvozar/backoff"
```

Package backoff provides a library that provides stateless backoff policies for reconnect loops used to maintain persistent connections. Randomized wait times are utilized to avoid the [thundering herd problem](https://en.wikipedia.org/wiki/Thundering_herd_problem).
