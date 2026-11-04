// Package fxmdcli supplies native Urfave commands for FXMacroData public JSON.
package fxmdcli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	sdk "github.com/fxmacrodata/fxmacrodata-go"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"
)

type safeError struct{ message string }

func (e safeError) Error() string { return e.message }
func fail(message string) error   { return safeError{message} }
func usageError(context.Context, *cli.Command, error, bool) error {
	return fail("invalid command arguments; use --help")
}
func noExit(context.Context, *cli.Command, error) {}

// NewCommand constructs an embeddable native command without reading keys or
// making a request. getenv is called only when a data command executes.
func NewCommand(getenv func(string) string, base http.RoundTripper) *cli.Command {
	if getenv == nil {
		getenv = func(string) string { return "" }
	}
	root := &cli.Command{Name: "fxmacrodata", Version: "0.1.0", Usage: "FXMacroData JSON access powered by Urfave CLI", OnUsageError: usageError, ExitErrHandler: noExit, DisableSliceFlagSeparator: true,
		Flags: []cli.Flag{&cli.DurationFlag{Name: "timeout", Value: 30 * time.Second, Usage: "Positive request timeout, at most 5m"}, &cli.BoolFlag{Name: "pretty", Usage: "Indent JSON without changing fields or numeric spelling"}},
	}
	root.Commands = sdkCommands(getenv, base)
	root.Commands = append(root.Commands, &cli.Command{Name: "request", DisableSliceFlagSeparator: true, Usage: "Request a documented public /v1/ JSON path", ArgsUsage: "/v1/documented/path", OnUsageError: usageError, ExitErrHandler: noExit,
		Flags:  []cli.Flag{&cli.StringFlag{Name: "method", Value: "GET", Usage: "GET or POST"}, &cli.StringSliceFlag{Name: "query", Usage: "Repeat name=value; commas stay within values"}, &cli.StringFlag{Name: "body", Usage: "Complete JSON body for POST"}},
		Action: func(ctx context.Context, c *cli.Command) error { return executeRaw(ctx, c, getenv, base) },
	})
	root.Action = func(ctx context.Context, c *cli.Command) error {
		if c.Args().Present() {
			return fail("unknown command; use --help")
		}
		return cli.ShowAppHelp(c)
	}
	return root
}

// Run buffers native output until success and sanitizes parser errors. Applications
// may instead embed NewCommand in their own Urfave tree and error-handling policy.
func Run(ctx context.Context, args []string, getenv func(string) string, base http.RoundTripper, output io.Writer) error {
	var stdout, stderr bytes.Buffer
	command := NewCommand(getenv, base)
	command.Writer = &stdout
	command.ErrWriter = &stderr
	if err := command.Run(ctx, args); err != nil {
		var known safeError
		if errors.As(err, &known) {
			return known
		}
		return fail("invalid command or request; use --help")
	}
	_, err := output.Write(stdout.Bytes())
	return err
}

func executeSDK(ctx context.Context, c *cli.Command, getenv func(string) string, base http.RoundTripper, call func(context.Context, *sdk.APIClient) error) error {
	return execute(ctx, c, getenv, base, func(ctx context.Context, httpClient *http.Client) error {
		config := sdk.NewConfiguration()
		config.Debug = false
		config.UserAgent = "FXMacroData-Urfave/0.1.0"
		config.HTTPClient = httpClient
		return call(ctx, sdk.NewAPIClient(config))
	})
}

func executeRaw(ctx context.Context, c *cli.Command, getenv func(string) string, base http.RoundTripper) error {
	if c.Args().Len() != 1 {
		return fail("request requires one documented public /v1/ path")
	}
	path, err := url.Parse(c.Args().First())
	if err != nil || path.Scheme != "" || path.Host != "" || path.User != nil || path.Fragment != "" || path.RawQuery != "" || !strings.HasPrefix(path.Path, "/v1/") {
		return fail("use a public /v1/ path and separate --query parameters")
	}
	query := url.Values{}
	for _, item := range c.StringSlice("query") {
		name, value, ok := strings.Cut(item, "=")
		if !ok || name == "" {
			return fail("query parameters require name=value")
		}
		query.Add(name, value)
	}
	method := strings.ToUpper(c.String("method"))
	body := c.String("body")
	if method != "GET" && method != "POST" {
		return fail("only GET and POST JSON requests are supported")
	}
	if len(body) > MaxBodyBytes || (body != "" && (method != "POST" || !json.Valid([]byte(body)))) {
		return fail("a bounded valid JSON body is allowed only for POST")
	}
	path.Scheme = "https"
	path.Host = "api.fxmacrodata.com"
	path.RawQuery = query.Encode()
	return execute(ctx, c, getenv, base, func(ctx context.Context, httpClient *http.Client) error {
		var reader io.Reader
		if body != "" {
			reader = strings.NewReader(body)
		}
		request, err := http.NewRequestWithContext(ctx, method, path.String(), reader)
		if err != nil {
			return fail("invalid public JSON request")
		}
		if body != "" {
			request.Header.Set("Content-Type", "application/json")
		}
		response, err := httpClient.Do(request)
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
		return err
	})
}

func execute(ctx context.Context, c *cli.Command, getenv func(string) string, base http.RoundTripper, call func(context.Context, *http.Client) error) error {
	timeout := c.Duration("timeout")
	if timeout <= 0 || timeout > 5*time.Minute {
		return fail("timeout must be positive and at most 5m")
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	guard := newTransport(base, getenv("FXMD_API_KEY"), ctx)
	client := &http.Client{Transport: guard, Timeout: timeout, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	_ = call(ctx, client)
	if ctx.Err() != nil {
		return fail("request timed out or was cancelled")
	}
	if !guard.complete || guard.status < 200 || guard.status >= 300 {
		if guard.status != 0 {
			return fail(fmt.Sprintf("FXMacroData request failed (HTTP %d); check parameters and access", guard.status))
		}
		return fail("FXMacroData request failed; check parameters and connection")
	}
	if !utf8.Valid(guard.raw) || !json.Valid(guard.raw) {
		return fail("FXMacroData returned invalid JSON")
	}
	output := guard.raw
	if c.Bool("pretty") {
		var formatted bytes.Buffer
		if json.Indent(&formatted, output, "", "  ") != nil {
			return fail("unable to format response JSON")
		}
		output = formatted.Bytes()
	}
	if _, err := c.Root().Writer.Write(output); err != nil {
		return fail("unable to write response")
	}
	if len(output) == 0 || output[len(output)-1] != '\n' {
		if _, err := c.Root().Writer.Write([]byte{'\n'}); err != nil {
			return fail("unable to write response")
		}
	}
	return nil
}
