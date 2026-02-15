# ptrToNew

A Go static analysis tool that finds usages of
`k8s.io/utils/ptr.To()` and suggests replacing them with
the built-in `new()` function.

In Go 1.26, the built-in `new` function now allows its
operand to be an expression, specifying the initial value
of the variable (previously it only accepted a type). This
makes `k8s.io/utils/ptr.To()` unnecessary.

## Before and After

| Before | After |
|---|---|
| `ptr.To(true)` | `new(true)` |
| `ptr.To("hello")` | `new("hello")` |
| `ptr.To(int32(10))` | `new(int32(10))` |
| `ptr.To[int32](10)` | `new(int32(10))` |
| `ptr.To(myFunc())` | `new(myFunc())` |

The tool also handles custom import aliases
(e.g. `import k8sptr "k8s.io/utils/ptr"`).

## Requirements

- Go 1.26 or later

## Installation

```bash
go install ptrtonew/cmd/ptr_to_new@latest
```

Or build from source:

```bash
make build
```

The binary is placed at `bin/ptrToNew`.

## Usage

### Standalone

ptrToNew uses the standard `go/analysis` singlechecker
driver, so it accepts the same flags as other Go
analyzers:

```bash
# Analyze packages
ptrToNew ./...

# Analyze a specific package
ptrToNew ./pkg/mypackage

# Apply suggested fixes automatically
ptrToNew -fix ./...
```

### With golangci-lint

ptrToNew implements the
`golang.org/x/tools/go/analysis.Analyzer` interface and
can be integrated as a custom linter in golangci-lint.

### Check Version

```bash
ptrToNew version
```

## How It Works

ptrToNew scans Go source files for imports of
`k8s.io/utils/ptr`. In files that import it, the tool
walks the AST looking for calls to `ptr.To(value)` or
`ptr.To[Type](value)` and reports a diagnostic with a
suggested fix that replaces the call with `new(value)` or
`new(Type(value))`, respectively.

All diagnostics are reported under the `"modernize"`
category.
