# FXMacroData for Urfave CLI

A native Urfave CLI v3 command package and executable for FXMacroData JSON access. It includes all 27 operations and 156 parameters from the published Go SDK v0.1.0, plus a standard JSON request command for other documented public v1 endpoints.

The package is importable as `github.com/fxmacrodata/urfave-fxmacrodata`. `NewCommand` returns a native `*cli.Command` you can embed in your application's command tree. `Run` provides buffered output and sanitized errors for a standalone executable. Install the `cmd/fxmacrodata` package at v0.1.0 with a supported Go toolchain.

```text
fxmacrodata --help
fxmacrodata calendar USD --help
fxmacrodata announcements USD cpi --limit 10
fxmacrodata forex USD EUR --pretty
fxmacrodata request /v1/data_catalogue/USD
```

Examples demonstrate request syntax, not data availability. Public requests may omit a key. For protected access, supply your own `FXMD_API_KEY` through your normal secret environment. The API enforces coverage and entitlements. The command never accepts keys as flags or stores them in a file; help does not read keys or call the API.

## Native embedding

```go
command := fxmdcli.NewCommand(os.Getenv, nil)
app := &cli.Command{Name: "my-app", Commands: []*cli.Command{command}}
err := app.Run(ctx, os.Args)
```

Use your application's own output and error-handling policy when embedding. To use the standalone sanitized runner, call `fxmdcli.Run(ctx, args, getenv, transport, writer)`. A nil transport uses the restricted HTTPS client; the optional transport is useful for application-owned testing.

## Full response and parameter access

Named commands use the original SDK request builders. Flags are applied only when explicitly set, including false and zero. Unspecified flags preserve the SDK's published defaults. Path arguments are positional; research-panel's complete JSON model is supplied through `--research-panel-request`. Each command's help lists its parameters.

`request` takes a documented `/v1/` path, optional repeated `--query=name=value`, `--method=GET|POST` and a bounded JSON `--body` for POST. Commas and duplicate query names are preserved. Do not put a query string, host or credentials in the path. This command does not invent endpoint names, change request schemas or imply access to undocumented routes.

Successful JSON retains original bytes and numeric spelling, unknown fields, provenance, quality flags, nulls, revisions and timestamps. A final newline is added when needed. `--pretty` changes whitespace only. SDK strict-decoder errors do not remove future fields from a successful complete HTTP JSON response. No pagination, date-window changes, missing values or future release times are inferred; request additional pages explicitly.

## Request boundaries

Only HTTPS requests to the FXMacroData public v1 API origin are permitted. Redirects, other hosts, traversal, credential query parameters, forwarded cookies/authorization/trailers and extra upstream calls are refused. Errors omit private bodies, URLs, credentials and input values. Requests/responses are bounded to 4 MiB. The default timeout is 30 seconds; `--timeout` accepts a positive duration no greater than five minutes. Cancellation propagates to the request.

Tests use fictional data and an in-memory transport with the actual Urfave parser. They cover all named SDK operations/parameters, raw JSON queries/bodies, native embedding, no-key/server-denial behavior, original number spelling, request restrictions, limits and error redaction. They do not contact a live API account. Code is MIT licensed. Each customer's plan and data rights govern access and redistribution.

[FXMacroData](https://fxmacrodata.com/?utm_source=github&utm_medium=integration&utm_campaign=open_source&utm_content=urfave_fxm) supplies official-source macroeconomic data. This companion is independent of the Urfave project and implies no endorsement.
