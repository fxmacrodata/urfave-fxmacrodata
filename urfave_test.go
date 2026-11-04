package fxmdcli

import (
	"bytes"
	"context"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTrip func(*http.Request) (*http.Response, error)

func (f roundTrip) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func response(status int, raw string) *http.Response {
	return &http.Response{StatusCode: status, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(raw))}
}
func fixtureKey(string) string { return "fictional-fixture-key" }
func run(args []string, getenv func(string) string, base http.RoundTripper) (string, error) {
	var out bytes.Buffer
	err := Run(context.Background(), append([]string{"fxmacrodata"}, args...), getenv, base, &out)
	return out.String(), err
}

func TestNativeHelpDoesNotReadKeyOrRequest(t *testing.T) {
	out, err := run([]string{"--help"}, func(string) string { t.Fatal("credential read for help"); return "" }, roundTrip(func(*http.Request) (*http.Response, error) { t.Fatal("request during help"); return nil, nil }))
	if err != nil || !strings.Contains(out, "research-panel") || !strings.Contains(out, "request") {
		t.Fatal(err, out)
	}
}
func TestCompleteJSONAndNoKeyFlow(t *testing.T) {
	raw := ` {"status":"ok","service":"fixture","future":{"n":9007199254740993,"decimal":1.2300,"missing":null}} `
	out, err := run([]string{"health"}, nil, roundTrip(func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("X-API-Key") != "" {
			t.Fatal("unexpected key")
		}
		return response(200, raw), nil
	}))
	if err != nil || out != raw+"\n" {
		t.Fatal(err, out)
	}
	out, err = run([]string{"calendar", "USD"}, nil, roundTrip(func(*http.Request) (*http.Response, error) {
		return response(403, `{"private":"fictional-fixture-key"}`), nil
	}))
	if err == nil || out != "" || strings.Contains(err.Error(), "fictional") {
		t.Fatal(err, out)
	}
}
func TestFalseZeroAndSDKDefaults(t *testing.T) {
	calls := 0
	_, err := run([]string{"announcements", "USD", "cpi", "--official-only=false", "--limit=0"}, fixtureKey, roundTrip(func(req *http.Request) (*http.Response, error) {
		calls++
		q := req.URL.Query()
		if !q.Has("official_only") || q.Get("official_only") != "false" || q.Get("limit") != "0" || q.Get("offset") != "0" || q.Has("page") {
			t.Fatal("parameters changed", q)
		}
		return response(400, `{}`), nil
	}))
	if err == nil || calls != 1 {
		t.Fatal(err, calls)
	}
}
func TestRawNativeRequestPreservesQueryAndBody(t *testing.T) {
	body := `{"series":[],"decision_times":[],"unknown_future_field":null}`
	raw := `{"data":[{"period":"fictional","revision":0,"value":1},{"period":"fictional","revision":1,"value":null}],"provenance":{"fixture":true}}`
	out, err := run([]string{"request", "/v1/research/panel", "--method=POST", "--body=" + body, "--query=scope=a,b", "--query=scope=c", "--query=limit=0"}, fixtureKey, roundTrip(func(req *http.Request) (*http.Response, error) {
		got, _ := io.ReadAll(req.Body)
		if string(got) != body || req.Method != "POST" || req.URL.Query().Get("limit") != "0" || strings.Join(req.URL.Query()["scope"], "|") != "a,b|c" || req.Header.Get("X-API-Key") != "fictional-fixture-key" {
			t.Fatal("request changed", string(got), req.URL)
		}
		return response(200, raw), nil
	}))
	if err != nil || out != raw+"\n" {
		t.Fatal(err, out)
	}
}
func TestInvalidInputIsRedactedBeforeNetwork(t *testing.T) {
	cases := [][]string{{"request", "https://foreign.example/v1/data"}, {"request", "/v1/../private"}, {"request", "/v1/health", "--query=api_key=fictional-private-marker"}, {"request", "/v1/research/panel", "--method=DELETE"}, {"request", "/v1/research/panel", "--method=POST", "--body=not-json"}, {"announcements", "USD", "cpi", "--limit=fictional-private-marker"}, {"fx-intraday-reference-rates", "USD", "EUR", "--start-time=fictional-private-marker"}, {"health", "--timeout=0s"}}
	for _, args := range cases {
		out, err := run(args, fixtureKey, roundTrip(func(*http.Request) (*http.Response, error) { t.Fatal("invalid input dispatched"); return nil, nil }))
		if err == nil || out != "" || strings.Contains(err.Error(), "fictional-private-marker") {
			t.Fatal(args, err, out)
		}
	}
}
func TestRedirectResponseLimitsAndCancellation(t *testing.T) {
	for _, status := range []int{302, 401, 429, 500} {
		calls := 0
		out, err := run([]string{"health"}, fixtureKey, roundTrip(func(*http.Request) (*http.Response, error) {
			calls++
			r := response(status, `{"key":"fictional-fixture-key"}`)
			r.Header.Set("Location", "https://foreign.example/")
			return r, nil
		}))
		if err == nil || out != "" || calls != 1 || strings.Contains(err.Error(), "fictional") {
			t.Fatal(status, err, out, calls)
		}
	}
	for _, raw := range []string{"not JSON", strings.Repeat("x", MaxBodyBytes+1)} {
		out, err := run([]string{"health"}, nil, roundTrip(func(*http.Request) (*http.Response, error) { return response(200, raw), nil }))
		if err == nil || out != "" {
			t.Fatal("invalid response accepted")
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var out bytes.Buffer
	err := Run(ctx, []string{"fxmacrodata", "health"}, nil, roundTrip(func(req *http.Request) (*http.Response, error) { return nil, req.Context().Err() }), &out)
	if err == nil || out.Len() != 0 {
		t.Fatal("cancellation ignored")
	}
}
func TestCommandFactoryEmbedsInNativeUrfaveTree(t *testing.T) {
	var out bytes.Buffer
	parent := &cli.Command{Name: "customer-app", Writer: &out, ErrWriter: io.Discard, Commands: []*cli.Command{NewCommand(nil, roundTrip(func(*http.Request) (*http.Response, error) {
		return response(200, `{"status":"ok","service":"fixture"}`), nil
	}))}}
	if err := parent.Run(context.Background(), []string{"customer-app", "fxmacrodata", "health"}); err != nil || !strings.Contains(out.String(), `"service":"fixture"`) {
		t.Fatal(err, out.String())
	}
}

func TestEverySDKOperationAndAllParametersViaNativeParser(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{name: "GetAnnouncementChanges", args: []string{"announcement-changes", "--currencies=fixture", "--indicators=fixture", "--since=fixture", "--limit=0", "--payload=fixture"}},
		{name: "GetAnnouncements", args: []string{"announcements", "USD", "USD", "--start-date=fixture", "--end-date=fixture", "--series-mode=fixture", "--limit=0", "--offset=0", "--page=0", "--seasonality=fixture", "--frequency=fixture", "--annualization=fixture", "--period-aggregation=fixture", "--revisions=fixture", "--basis=fixture", "--official-only=false", "--as-of=fixture", "--availability=fixture", "--dataset-version=fixture", "--value-mode=fixture"}},
		{name: "GetCalendar", args: []string{"calendar", "USD", "--indicator=fixture", "--start-date=fixture", "--end-date=fixture", "--timezone=fixture"}},
		{name: "GetCommodities", args: []string{"commodities", "USD", "--start-date=fixture", "--end-date=fixture", "--limit=0", "--offset=0", "--page=0"}},
		{name: "GetCot", args: []string{"cot", "USD", "--start-date=fixture", "--end-date=fixture", "--limit=0", "--offset=0", "--page=0"}},
		{name: "GetCurves", args: []string{"curves", "USD", "--curve-family=fixture", "--metric=fixture", "--view=fixture", "--method=fixture", "--date=fixture", "--as-of=fixture", "--availability=fixture", "--max-staleness-days=0", "--require-common-date=false"}},
		{name: "GetDataCatalogue", args: []string{"data-catalogue", "USD", "--include-capabilities=false", "--include-coverage=false", "--indicator=fixture"}},
		{name: "GetFactor", args: []string{"factor", "USD", "USD", "--start-date=fixture", "--end-date=fixture", "--include-components=false", "--include-sources=false", "--limit=0", "--offset=0"}},
		{name: "GetFinancialPrices", args: []string{"financial-prices", "USD", "--source=fixture", "--start-date=fixture", "--end-date=fixture", "--measure=fixture", "--instrument-id=fixture", "--issuer=fixture", "--limit=0", "--offset=0"}},
		{name: "GetForex", args: []string{"forex", "USD", "USD", "--start-date=fixture", "--end-date=fixture", "--limit=0", "--offset=0", "--page=0", "--indicators=fixture", "--dataset-version=fixture"}},
		{name: "GetFxIntradayReferenceRates", args: []string{"fx-intraday-reference-rates", "USD", "USD", "--start-time=2026-01-01T00:00:00Z", "--end-time=2026-01-01T00:00:00Z"}},
		{name: "GetFxReferenceRates", args: []string{"fx-reference-rates", "USD", "USD", "--date=fixture", "--start-date=fixture", "--end-date=fixture", "--source=fixture", "--timestamp-type=fixture", "--native-only=false", "--include-source-legs=false", "--order=fixture", "--limit=0", "--offset=0", "--page=0", "--dataset-version=fixture"}},
		{name: "GetFxSessionRates", args: []string{"fx-session-rates", "USD", "USD", "--session=fixture", "--date=fixture", "--start-date=fixture", "--end-date=fixture", "--order=fixture", "--limit=0", "--offset=0", "--page=0", "--dataset-version=fixture"}},
		{name: "GetFxSourceUniverse", args: []string{"fx-source-universe", "--currency=fixture", "--source=fixture"}},
		{name: "GetFxSources", args: []string{"fx-sources"}},
		{name: "GetHealth", args: []string{"health"}},
		{name: "GetLatestAnnouncements", args: []string{"latest-announcements", "USD"}},
		{name: "GetLatestCommodities", args: []string{"latest-commodities"}},
		{name: "GetMarketSessions", args: []string{"market-sessions", "--at=fixture"}},
		{name: "GetPredictionCoverage", args: []string{"prediction-coverage"}},
		{name: "GetPredictionCoverageForCurrency", args: []string{"prediction-coverage-for-currency", "USD"}},
		{name: "GetPredictions", args: []string{"predictions", "USD", "USD", "--prediction-class=fixture", "--prediction-type=fixture", "--prediction-source=fixture", "--official-limit=0", "--pre-release-only=false", "--seasonality=fixture", "--frequency=fixture", "--annualization=fixture", "--period-aggregation=fixture", "--basis=fixture", "--start-date=fixture", "--end-date=fixture", "--limit=0", "--offset=0", "--page=0", "--before-date=fixture", "--as-of=fixture", "--event-compatible-only=false"}},
		{name: "GetPressReleases", args: []string{"press-releases", "USD", "--limit=0", "--offset=0"}},
		{name: "GetRateDifferentials", args: []string{"rate-differentials", "USD", "USD", "--measure=fixture", "--rate-type=fixture", "--curve-family=fixture", "--start-tenor-years=1.5", "--end-tenor-years=1.5", "--start-date=fixture", "--end-date=fixture", "--limit=0", "--offset=0", "--as-of=fixture", "--availability=fixture"}},
		{name: "GetRiskSentiment", args: []string{"risk-sentiment", "--start-date=fixture", "--end-date=fixture", "--limit=0", "--offset=0"}},
		{name: "Ping", args: []string{"ping"}},
		{name: "ResearchPanel", args: []string{"research-panel", "--research-panel-request={\"series\":[],\"decision_times\":[],\"start_date\":\"2026-01-01\",\"end_date\":\"2026-01-02\"}"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			out, err := run(tc.args, fixtureKey, roundTrip(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.URL.Host != "api.fxmacrodata.com" || req.Header.Get("X-API-Key") == "" {
					t.Fatal("incorrect authenticated request")
				}
				return response(400, `{"fixture":true}`), nil
			}))
			if err == nil || out != "" || calls != 1 {
				t.Fatal("native request not dispatched", err, out, calls)
			}
		})
	}
}
