# cosine_similarity
Cosine similarity implementation in Go.
Can be used in the terminal as well as a go package.

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

### In the terminal
```
git clone https://github.com/omgupta1608/cosine_similarity.git
```

#### Run it
```
./cosine_similarity "Who are you?" "Who is that?"
```
#### Output
```
Cosine Similarity Score for "Who are you?" and "Who is that?" is: 0.5773502691896258
```