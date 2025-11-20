package ui

import (
	"path/filepath"
	"strings"
)

// CodeFenceLanguageFromPath returns a glamour-friendly language identifier
// derived from a file path extension. An empty string means "no preference".
func CodeFenceLanguageFromPath(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".go":
		return "go"
	case ".ts", ".tsx":
		return "typescript"
	case ".js", ".jsx", ".mjs", ".cjs":
		return "javascript"
	case ".py":
		return "python"
	case ".rb":
		return "ruby"
	case ".rs":
		return "rust"
	case ".java":
		return "java"
	case ".kt", ".kts":
		return "kotlin"
	case ".swift":
		return "swift"
	case ".php":
		return "php"
	case ".cs":
		return "csharp"
	case ".cc", ".cpp", ".cxx", ".hpp", ".hh":
		return "cpp"
	case ".c", ".h":
		return "c"
	case ".m", ".mm":
		return "objective-c"
	case ".sh", ".bash", ".zsh":
		return "bash"
	case ".ps1":
		return "powershell"
	case ".yaml", ".yml":
		return "yaml"
	case ".json":
		return "json"
	case ".toml":
		return "toml"
	case ".tf":
		return "hcl"
	case ".sql":
		return "sql"
	case ".md", ".markdown":
		return "markdown"
	default:
		return ""
	}
}
