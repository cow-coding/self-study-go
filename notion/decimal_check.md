# Go에서 문자가 숫자인지 확인하기

## 패키지

- [`unicode`](https://pkg.go.dev/unicode)

## 사용 함수

- `IsDigit`
- `IsNumber`

### `IsDigit`

```go
func IsDigit(r rune) bool
```
- IsDigit reports whether the rune is a decimal digit.

공식문서에 따르면 사용법은 아래와 같다.

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsDigit('৩'))
	fmt.Printf("%t\n", unicode.IsDigit('A'))
}
```

### `IsNumber`

```go
func IsNumber(r rune) bool
```
- IsNumber reports whether the rune is a number (category N).

공식문서에 따르면 사용법은 아래와 같다.

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsNumber('Ⅷ'))
	fmt.Printf("%t\n", unicode.IsNumber('A'))
}
```

## `IsDigit` vs `IsNumber`

얼핏보면 두 함수는 동일한 것으로 보이는데, 그렇다면 무엇이 다를까?  
두 함수의 설명이 약간 다르다. 한번 비교해보자

- `IsDigit`의 설명
  - `IsDigit reports whether the rune is a decimal digit.`
- `IsNumber`의 설명
  - `IsNumber reports whether the rune is a number (category N).`

정확히 말하면 `IsDigit`은 decimal, 즉 10진수에 한해서 동작하는 것으로 보인다. 근데 그렇다면 `IsNumber`는 10진수가 아닌 수에서 된다는 걸까? 그건 아닌 것 같다. 일단 설명에 있는 category N이 뭔지 알아보자.  

### Unicode Category N

![](/images/unicode.png)

`unicode` 패키지에서 검색을 해보니 unicode character와 관련이 깊은거 같다.  
unicode의 category N을 알아보니... unicode 사이트로는 감이 안오다가 위키에서 도움이 되는 내용을 발견했다.

![](/images/uni_wiki.png)

쉽게 이해해보면 Nd는 Numeric type이고 decimal이 된다. Nl과 No가 뭐냐?라는 생각이 드는데, 이게 한국에서는 크게 이해가 안되는 상황이라 발생하는 문제이다.  

일반적으로 10진수 체계가 잡힌 현대에는 이런 경우가 별로 없지만 고대 그리스 혹은 고딕 문자에는 100을 나타내는 문자가 따로 존재하는 경우가 있다. 이런 경우에는 letter 숫자가 되는데, 이는 10진 체계를 벗어난 표현이라 유니코드에서는 `Nl`로 분류한다.  

추가로 No는 뭔가 싶은데, 유니코드는 상당히 많은 문자가 존재한다. 예를 들면 우리의 상상을 초월하는... ½ 이런 것들이 있다. 참고로 1/2와 다르다. 이런 경우는 No로 분류한다.

### 결론

결국 `IsDigit`은 10진수로 통용되는 문자를 처리하고 `IsNumber`는 **숫자**라고 표현되는 문자들을 모두 구분해준다. 즉 `IsNumber`의 범위가 더 넓다는 것이다.

## Reference

- [StackOverflow : Differences between IsDigit and IsNumber in unicode in Go](https://stackoverflow.com/questions/25540951/differences-between-isdigit-and-isnumber-in-unicode-in-go)
