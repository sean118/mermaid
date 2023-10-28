//go:build linux || darwin

package markdown_test

import (
	"os"

	"github.com/go-spectest/markdown"
	"github.com/go-spectest/mermaid/sequence"
)

// Examle is example code. Skip this test on Windows.
// The newline codes in the comment section where
// the expected values are written are represented as '\n',
// causing failures when testing on Windows.
func Example() {
	diagram := sequence.NewDiagram(os.Stdout).
		Participant("Sophia").
		Participant("David").
		Participant("Subaru").
		LF().
		SyncRequest("Sophia", "David", "Please wake up Subaru").
		SyncResponse("David", "Sophia", "OK").
		LF().
		LoopStart("until Subaru wake up").
		SyncRequest("David", "Subaru", "Wake up!").
		SyncResponse("Subaru", "David", "zzz").
		SyncRequest("David", "Subaru", "Hey!!!").
		BreakStart("if Subaru wake up").
		SyncResponse("Subaru", "David", "......").
		BreakEnd().
		LoopEnd().
		LF().
		SyncResponse("David", "Sophia", "wake up, wake up").
		String()

	markdown.NewMarkdown(os.Stdout).
		H2("Sequence Diagram").
		CodeBlocks(markdown.SyntaxHighlightMermaid, diagram).
		Build() //nolint

	// Output:
	// ## Sequence Diagram
	// ```mermaid
	// sequenceDiagram
	//     participant Sophia
	//     participant David
	//     participant Subaru
	//
	//     Sophia->>David: Please wake up Subaru
	//     David-->>Sophia: OK
	//
	//     loop until Subaru wake up
	//     David->>Subaru: Wake up!
	//     Subaru-->>David: zzz
	//     David->>Subaru: Hey!!!
	//     break if Subaru wake up
	//     Subaru-->>David: ......
	//     end
	//     end
	//
	//     David-->>Sophia: wake up, wake up
	// ```
}
