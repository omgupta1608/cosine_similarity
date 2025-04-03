# Cosine Similarity implementation in Go
Cosine similarity quantifies the resemblance between two vectors in a multi-dimensional setting, often applied to gauge the similarity between textual entities. Conceptually, it evaluates the angle between the directions these vectors point towards, indicating their similarity. When this angle is minimal, approaching zero degrees, the cosine value tends toward 1, denoting a stronger alignment and thus higher textual similarity.

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
    score, err := cs.Cosine(a, b, false) // true if you want to ignore case
    if err != nil {
        fmt.Println("error in calculating cosine similarity: ", err.Error())
        return
    }
    fmt.Println("Score: ", score)
}
```
#### Output
```
Score: 0.5773502691896258
```
