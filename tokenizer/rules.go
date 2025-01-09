package tokenizer

var rulesets = map[string]LanguageRule{
	"cpp": {
		Name: "cpp",
		Operators: []string{
			`#include\s*<[^>]+>`, // Include directives
			`\b\w+\s*\([^)]*\)`,  // Function calls
			`[-+*/=<>!&|^%]=?`,   // Basic operators and assignments
			`\+\+|--`,            // Increment/decrement
			`&&|\|\|`,            // Logical operators
			`\b(if|else|while|for|do|switch|case|break|continue|return|new|delete)\b`, // Keywords
			`\b(class|struct|namespace|public|private|protected)\b`,                   // OOP keywords
			`<<|>>`,       // Stream operators
			`::|->|\.|::`, // Scope and member access
			`\[\]`,        // Array subscript
		},
		Operands: []string{
			`"[^"]*"`,            // String literals
			`'[^']*'`,            // Character literals
			`\b\d+(\.\d+)?\b`,    // Numbers (integer and float)
			`\b[a-zA-Z_]\w*\b`,   // Identifiers
			`\btrue\b|\bfalse\b`, // Boolean literals
			`\bnullptr\b`,        // Null pointer
		},
		Skip: []string{
			`\/\/.*`,           // Single line comments
			`\/\*[\s\S]*?\*\/`, // Multi line comments
			`\s+`,              // Whitespace
		},
	},
	"java": {
		Name: "java",
		Operators: []string{
			`[{}()]`, // Braces, parentheses
			`[;,:]`,  // Semicolons, commas
			`\baspect\b|\bpointcut\b|\bexecution\b|\bbefore\b`,
			`\bthisJoinPoint\b`,
			`\b\w+\s*\([^)]*\)`,
			`[-+*/=<>!&|^%]=?`,
			`\+\+|--`,
			`&&|\|\|`,
			`\b(if|else|while|for|do|switch|case|break|continue|return|new)\b`,
			`\b(class|interface|extends|implements|public|private|protected)\b`,
			`\b(try|catch|finally|throw|throws)\b`,
			`->|\.|@`,
			`\[\]`,
		},
		Operands: []string{
			`"[^"]*"`,
			`'[^']*'`,
			`\b\d+(\.\d+)?\b`,
			`\b[a-zA-Z_]\w*\b`,
			`\btrue\b|\bfalse\b`,
			`\bnull\b`,
		},
		Skip: []string{
			`\/\/.*`,
			`\/\*[\s\S]*?\*\/`,
			`\s+`,
		},
	},
	"python": {
		Name: "python",
		Operators: []string{
			`\bimport\b|\bfrom\b`, // Import statements
			`\b\w+\s*\([^)]*\)`,   // Function calls
			`[-+*/=<>!&|^%]=?`,    // Basic operators and assignments
			`\*\*`,                // Power operator
			`and|or|not`,          // Logical operators
			`\b(if|elif|else|while|for|in|break|continue|return|def|class)\b`, // Keywords
			`\b(try|except|finally|raise|with|as)\b`,                          // Exception handling
			`\.|@`,                                                            // Member access, decorators
			`\[\]`,                                                            // List subscript
			`:\s*$`,                                                           // Block delimiter
		},
		Operands: []string{
			`"[^"]*"|'[^']*'`,    // String literals (single/double quotes)
			`\b\d+(\.\d+)?\b`,    // Numbers
			`\b[a-zA-Z_]\w*\b`,   // Identifiers
			`\bTrue\b|\bFalse\b`, // Boolean literals
			`\bNone\b`,           // None literal
		},
		Skip: []string{
			`#.*`,            // Single line comments
			`'''[\s\S]*?'''`, // Multi line comments (triple quotes)
			`"""[\s\S]*?"""`, // Multi line comments (triple double quotes)
			`\s+`,            // Whitespace
		},
	},
}