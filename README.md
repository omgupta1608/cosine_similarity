# cosine_similarity
Cosine similarity implementation in Go.

## How to use
```
go get github.com/omgupta1608/cosine_similarity
```

```go
import (
    "fmt"
    cs "github.com/omgupta1608/cosine_similarity"
)

func main() {
    a := "Who are you?"
    b := "Who is that?" 
    score := cs.Cosine(a, b)
    fmt.Println("Score: ", score)
}
```
#### Output
```
Score: 0.5773502691896258
```