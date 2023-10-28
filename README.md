[![Go Reference](https://pkg.go.dev/badge/github.com/go-spectest/mermaid.svg)](https://pkg.go.dev/github.com/go-spectest/mermaid)
[![LinuxUnitTest](https://github.com/go-spectest/mermaid/actions/workflows/linux_test.yml/badge.svg)](https://github.com/go-spectest/mermaid/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/go-spectest/mermaid/actions/workflows/mac_test.yml/badge.svg)](https://github.com/go-spectest/mermaid/actions/workflows/mac_test.yml)
[![WindowsUnitTest](https://github.com/go-spectest/mermaid/actions/workflows/windows_test.yml/badge.svg)](https://github.com/go-spectest/mermaid/actions/workflows/windows_test.yml)
[![reviewdog](https://github.com/go-spectest/mermaid/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/go-spectest/mermaid/actions/workflows/reviewdog.yml)
[![Gosec](https://github.com/go-spectest/mermaid/actions/workflows/gosec.yml/badge.svg)](https://github.com/go-spectest/mermaid/actions/workflows/gosec.yml)
# What is mermaid package
The Package mermaid is a simple mermaid builder in golang. This library assembles mermaid using method chaining, not uses a template engine like [html/template](https://pkg.go.dev/html/template). Currently, this library supports only sequence diagram. 
  
This library was initially developed to display test results in [go-spectest/spectest](https://github.com/go-spectest/spectest). Therefore, it implements the features required by spectest, but there are no plans to add additional functionalities unless requested by someone. This library is intended to be used in conjunction with [go-spectest/markdown](https://github.com/go-spectest/markdown).
  
Additionally, complex code that increases the complexity of the library, such as generating nested lists, will not be added. I want to keep this library as simple as possible.
  
## Supported OS and Go version
- OS: Linux, macOS, Windows
- Go: 1.18 or later
  
## Supported mermaid syntax (sequence diagram)
- [x] participant
- [ ] participant alias (not supported)
- [x] actor
- [ ] actor alias (not supported)
- [x] activate / deactivate
- [x] note over / left of / right of
- [ ] link
- [x] sync / async request
- [x] sync / async response
- [x] loop
- [x] alt / else / opt
- [x] parrallel
- [x] break
- [x] critical
- [x] group

## Example
### Basic usage
Generate mermaid in markdown.

```go
package main

import (
	"os"

	"github.com/go-spectest/markdown"
	"github.com/go-spectest/mermaid/sequence"
)

//go:generate go run main.go

func main() {
	diagram := sequence.NewDiagram(os.Stdout).
		Participant("Viktoriya").
		Participant("Naohiro").
		Participant("Naoyuki").
		LF().
		SyncRequest("Viktoriya", "Naohiro", "Please wake up Naoyuki").
		SyncResponse("Naohiro", "Viktoriya", "OK").
		LF().
		LoopStart("until Naoyuki wake up").
		SyncRequest("Naohiro", "Naoyuki", "Wake up!").
		SyncResponse("Naoyuki", "Naohiro", "zzz").
		SyncRequest("Naohiro", "Naoyuki", "Hey!!!").
		BreakStart("if Naoyuki wake up").
		SyncResponse("Naoyuki", "Naohiro", "......").
		BreakEnd().
		LoopEnd().
		LF().
		SyncResponse("Naohiro", "Viktoriya", "wake up, wake up").
		String()

	markdown.NewMarkdown(os.Stdout).
		H2("Sequence Diagram").
		CodeBlocks(markdown.SyntaxHighlightMermaid, diagram).
		Build()
}
```

Output: [markdown is here](./doc/generated.md)
````
## Sequence Diagram
```mermaid
sequenceDiagram
    participant Viktoriya
    participant Naohiro
    participant Naoyuki

    Viktoriya->>Naohiro: Please wake up Naoyuki
    Naohiro-->>Viktoriya: OK

    loop until Naoyuki wake up
    Naohiro->>Naoyuki: Wake up!
    Naoyuki-->>Naohiro: zzz
    Naohiro->>Naoyuki: Hey!!!
    break if Naoyuki wake up
    Naoyuki-->>Naohiro: ......
    end
    end

    Naohiro-->>Viktoriya: wake up, wake up
```
````

You can generate mermaid using go generate. Please define code to generate Mermaid first. Then, run "go generate ./..." to generate it.

## Contribution
First off, thanks for taking the time to contribute! Contributions are not only related to development. For example, GitHub Star motivates me to develop! Please feel free to contribute to this project.

## LICENSE
[MIT License](./LICENSE)