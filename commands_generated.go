// Code generated from published FXMacroData Go SDK v0.1.0 request builders.
package fxmdcli

import (
	"context"
	"encoding/json"
	sdk "github.com/fxmacrodata/fxmacrodata-go"
	"github.com/urfave/cli/v3"
	"net/http"
	"time"
)

func sdkCommands(getenv func(string) string, base http.RoundTripper) []*cli.Command {
	commands := []*cli.Command{}
	{
		c := &cli.Command{Name: "announcement-changes", Usage: "Call published GetAnnouncementChanges", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "currencies", Usage: "Published Currencies parameter (string)"}, &cli.StringFlag{Name: "indicators", Usage: "Published Indicators parameter (string)"}, &cli.StringFlag{Name: "since", Usage: "Published Since parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.StringFlag{Name: "payload", Usage: "Published Payload parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetAnnouncementChanges(ctx)
				if c.IsSet("currencies") {
					v := c.String("currencies")
					r = r.Currencies(v)
				}
				if c.IsSet("indicators") {
					v := c.String("indicators")
					r = r.Indicators(v)
				}
				if c.IsSet("since") {
					v := c.String("since")
					r = r.Since(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("payload") {
					v := c.String("payload")
					r = r.Payload(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "announcements", Usage: "Call published GetAnnouncements", ArgsUsage: "currency indicator", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.StringFlag{Name: "series-mode", Usage: "Published SeriesMode parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}, &cli.StringFlag{Name: "seasonality", Usage: "Published Seasonality parameter (string)"}, &cli.StringFlag{Name: "frequency", Usage: "Published Frequency parameter (string)"}, &cli.StringFlag{Name: "annualization", Usage: "Published Annualization parameter (string)"}, &cli.StringFlag{Name: "period-aggregation", Usage: "Published PeriodAggregation parameter (string)"}, &cli.StringFlag{Name: "revisions", Usage: "Published Revisions parameter (string)"}, &cli.StringFlag{Name: "basis", Usage: "Published Basis parameter (string)"}, &cli.BoolFlag{Name: "official-only", Usage: "Published OfficialOnly parameter (bool)"}, &cli.StringFlag{Name: "as-of", Usage: "Published AsOf parameter (string)"}, &cli.StringFlag{Name: "availability", Usage: "Published Availability parameter (string)"}, &cli.StringFlag{Name: "dataset-version", Usage: "Published DatasetVersion parameter (string)"}, &cli.StringFlag{Name: "value-mode", Usage: "Published ValueMode parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetAnnouncements(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("series-mode") {
					v := c.String("series-mode")
					r = r.SeriesMode(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				if c.IsSet("seasonality") {
					v := c.String("seasonality")
					r = r.Seasonality(v)
				}
				if c.IsSet("frequency") {
					v := c.String("frequency")
					r = r.Frequency(v)
				}
				if c.IsSet("annualization") {
					v := c.String("annualization")
					r = r.Annualization(v)
				}
				if c.IsSet("period-aggregation") {
					v := c.String("period-aggregation")
					r = r.PeriodAggregation(v)
				}
				if c.IsSet("revisions") {
					v := c.String("revisions")
					r = r.Revisions(v)
				}
				if c.IsSet("basis") {
					v := c.String("basis")
					r = r.Basis(v)
				}
				if c.IsSet("official-only") {
					v := c.Bool("official-only")
					r = r.OfficialOnly(v)
				}
				if c.IsSet("as-of") {
					v := c.String("as-of")
					r = r.AsOf(v)
				}
				if c.IsSet("availability") {
					v := c.String("availability")
					r = r.Availability(v)
				}
				if c.IsSet("dataset-version") {
					v := c.String("dataset-version")
					r = r.DatasetVersion(v)
				}
				if c.IsSet("value-mode") {
					v := c.String("value-mode")
					r = r.ValueMode(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "calendar", Usage: "Call published GetCalendar", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "indicator", Usage: "Published Indicator parameter (string)"}, &cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.StringFlag{Name: "timezone", Usage: "Published Timezone parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetCalendar(ctx, c.Args().Get(0))
				if c.IsSet("indicator") {
					v := c.String("indicator")
					r = r.Indicator(v)
				}
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("timezone") {
					v := c.String("timezone")
					r = r.Timezone(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "commodities", Usage: "Call published GetCommodities", ArgsUsage: "indicator", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetCommodities(ctx, c.Args().Get(0))
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "cot", Usage: "Call published GetCot", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetCot(ctx, c.Args().Get(0))
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "curves", Usage: "Call published GetCurves", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "curve-family", Usage: "Published CurveFamily parameter (string)"}, &cli.StringFlag{Name: "metric", Usage: "Published Metric parameter (string)"}, &cli.StringFlag{Name: "view", Usage: "Published View parameter (string)"}, &cli.StringFlag{Name: "method", Usage: "Published Method parameter (string)"}, &cli.StringFlag{Name: "date", Usage: "Published Date parameter (string)"}, &cli.StringFlag{Name: "as-of", Usage: "Published AsOf parameter (string)"}, &cli.StringFlag{Name: "availability", Usage: "Published Availability parameter (string)"}, &cli.Int64Flag{Name: "max-staleness-days", Usage: "Published MaxStalenessDays parameter (int64)"}, &cli.BoolFlag{Name: "require-common-date", Usage: "Published RequireCommonDate parameter (bool)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetCurves(ctx, c.Args().Get(0))
				if c.IsSet("curve-family") {
					v := c.String("curve-family")
					r = r.CurveFamily(v)
				}
				if c.IsSet("metric") {
					v := c.String("metric")
					r = r.Metric(v)
				}
				if c.IsSet("view") {
					v := c.String("view")
					r = r.View(v)
				}
				if c.IsSet("method") {
					v := c.String("method")
					r = r.Method(v)
				}
				if c.IsSet("date") {
					v := c.String("date")
					r = r.Date(v)
				}
				if c.IsSet("as-of") {
					v := c.String("as-of")
					r = r.AsOf(v)
				}
				if c.IsSet("availability") {
					v := c.String("availability")
					r = r.Availability(v)
				}
				if c.IsSet("max-staleness-days") {
					v := c.Int64("max-staleness-days")
					r = r.MaxStalenessDays(v)
				}
				if c.IsSet("require-common-date") {
					v := c.Bool("require-common-date")
					r = r.RequireCommonDate(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "data-catalogue", Usage: "Call published GetDataCatalogue", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.BoolFlag{Name: "include-capabilities", Usage: "Published IncludeCapabilities parameter (bool)"}, &cli.BoolFlag{Name: "include-coverage", Usage: "Published IncludeCoverage parameter (bool)"}, &cli.StringFlag{Name: "indicator", Usage: "Published Indicator parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetDataCatalogue(ctx, c.Args().Get(0))
				if c.IsSet("include-capabilities") {
					v := c.Bool("include-capabilities")
					r = r.IncludeCapabilities(v)
				}
				if c.IsSet("include-coverage") {
					v := c.Bool("include-coverage")
					r = r.IncludeCoverage(v)
				}
				if c.IsSet("indicator") {
					v := c.String("indicator")
					r = r.Indicator(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "factor", Usage: "Call published GetFactor", ArgsUsage: "currency factor", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.BoolFlag{Name: "include-components", Usage: "Published IncludeComponents parameter (bool)"}, &cli.BoolFlag{Name: "include-sources", Usage: "Published IncludeSources parameter (bool)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFactor(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("include-components") {
					v := c.Bool("include-components")
					r = r.IncludeComponents(v)
				}
				if c.IsSet("include-sources") {
					v := c.Bool("include-sources")
					r = r.IncludeSources(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "financial-prices", Usage: "Call published GetFinancialPrices", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "source", Usage: "Published Source parameter (string)"}, &cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.StringFlag{Name: "measure", Usage: "Published Measure parameter (string)"}, &cli.StringFlag{Name: "instrument-id", Usage: "Published InstrumentId parameter (string)"}, &cli.StringFlag{Name: "issuer", Usage: "Published Issuer parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFinancialPrices(ctx, c.Args().Get(0))
				if c.IsSet("source") {
					v := c.String("source")
					r = r.Source(v)
				}
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("measure") {
					v := c.String("measure")
					r = r.Measure(v)
				}
				if c.IsSet("instrument-id") {
					v := c.String("instrument-id")
					r = r.InstrumentId(v)
				}
				if c.IsSet("issuer") {
					v := c.String("issuer")
					r = r.Issuer(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "forex", Usage: "Call published GetForex", ArgsUsage: "base quote", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}, &cli.StringFlag{Name: "indicators", Usage: "Published Indicators parameter (string)"}, &cli.StringFlag{Name: "dataset-version", Usage: "Published DatasetVersion parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetForex(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				if c.IsSet("indicators") {
					v := c.String("indicators")
					r = r.Indicators(v)
				}
				if c.IsSet("dataset-version") {
					v := c.String("dataset-version")
					r = r.DatasetVersion(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "fx-intraday-reference-rates", Usage: "Call published GetFxIntradayReferenceRates", ArgsUsage: "base quote", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-time", Usage: "Published StartTime parameter (time.Time)"}, &cli.StringFlag{Name: "end-time", Usage: "Published EndTime parameter (time.Time)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFxIntradayReferenceRates(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("start-time") {
					v, err := time.Parse(time.RFC3339, c.String("start-time"))
					if err != nil {
						return fail("timestamp parameters require RFC3339")
					}
					r = r.StartTime(v)
				}
				if c.IsSet("end-time") {
					v, err := time.Parse(time.RFC3339, c.String("end-time"))
					if err != nil {
						return fail("timestamp parameters require RFC3339")
					}
					r = r.EndTime(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "fx-reference-rates", Usage: "Call published GetFxReferenceRates", ArgsUsage: "base quote", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "date", Usage: "Published Date parameter (string)"}, &cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.StringFlag{Name: "source", Usage: "Published Source parameter (string)"}, &cli.StringFlag{Name: "timestamp-type", Usage: "Published TimestampType parameter (string)"}, &cli.BoolFlag{Name: "native-only", Usage: "Published NativeOnly parameter (bool)"}, &cli.BoolFlag{Name: "include-source-legs", Usage: "Published IncludeSourceLegs parameter (bool)"}, &cli.StringFlag{Name: "order", Usage: "Published Order parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}, &cli.StringFlag{Name: "dataset-version", Usage: "Published DatasetVersion parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFxReferenceRates(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("date") {
					v := c.String("date")
					r = r.Date(v)
				}
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("source") {
					v := c.String("source")
					r = r.Source(v)
				}
				if c.IsSet("timestamp-type") {
					v := c.String("timestamp-type")
					r = r.TimestampType(v)
				}
				if c.IsSet("native-only") {
					v := c.Bool("native-only")
					r = r.NativeOnly(v)
				}
				if c.IsSet("include-source-legs") {
					v := c.Bool("include-source-legs")
					r = r.IncludeSourceLegs(v)
				}
				if c.IsSet("order") {
					v := c.String("order")
					r = r.Order(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				if c.IsSet("dataset-version") {
					v := c.String("dataset-version")
					r = r.DatasetVersion(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "fx-session-rates", Usage: "Call published GetFxSessionRates", ArgsUsage: "base quote", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "session", Usage: "Published Session parameter (string)"}, &cli.StringFlag{Name: "date", Usage: "Published Date parameter (string)"}, &cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.StringFlag{Name: "order", Usage: "Published Order parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}, &cli.StringFlag{Name: "dataset-version", Usage: "Published DatasetVersion parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFxSessionRates(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("session") {
					v := c.String("session")
					r = r.Session(v)
				}
				if c.IsSet("date") {
					v := c.String("date")
					r = r.Date(v)
				}
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("order") {
					v := c.String("order")
					r = r.Order(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				if c.IsSet("dataset-version") {
					v := c.String("dataset-version")
					r = r.DatasetVersion(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "fx-source-universe", Usage: "Call published GetFxSourceUniverse", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "currency", Usage: "Published Currency parameter (string)"}, &cli.StringFlag{Name: "source", Usage: "Published Source parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFxSourceUniverse(ctx)
				if c.IsSet("currency") {
					v := c.String("currency")
					r = r.Currency(v)
				}
				if c.IsSet("source") {
					v := c.String("source")
					r = r.Source(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "fx-sources", Usage: "Call published GetFxSources", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetFxSources(ctx)
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "health", Usage: "Call published GetHealth", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetHealth(ctx)
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "latest-announcements", Usage: "Call published GetLatestAnnouncements", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetLatestAnnouncements(ctx, c.Args().Get(0))
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "latest-commodities", Usage: "Call published GetLatestCommodities", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetLatestCommodities(ctx)
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "market-sessions", Usage: "Call published GetMarketSessions", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "at", Usage: "Published At parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetMarketSessions(ctx)
				if c.IsSet("at") {
					v := c.String("at")
					r = r.At(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "prediction-coverage", Usage: "Call published GetPredictionCoverage", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetPredictionCoverage(ctx)
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "prediction-coverage-for-currency", Usage: "Call published GetPredictionCoverageForCurrency", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetPredictionCoverageForCurrency(ctx, c.Args().Get(0))
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "predictions", Usage: "Call published GetPredictions", ArgsUsage: "currency indicator", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "prediction-class", Usage: "Published PredictionClass parameter (string)"}, &cli.StringFlag{Name: "prediction-type", Usage: "Published PredictionType parameter (string)"}, &cli.StringFlag{Name: "prediction-source", Usage: "Published PredictionSource parameter (string)"}, &cli.Int64Flag{Name: "official-limit", Usage: "Published OfficialLimit parameter (int64)"}, &cli.BoolFlag{Name: "pre-release-only", Usage: "Published PreReleaseOnly parameter (bool)"}, &cli.StringFlag{Name: "seasonality", Usage: "Published Seasonality parameter (string)"}, &cli.StringFlag{Name: "frequency", Usage: "Published Frequency parameter (string)"}, &cli.StringFlag{Name: "annualization", Usage: "Published Annualization parameter (string)"}, &cli.StringFlag{Name: "period-aggregation", Usage: "Published PeriodAggregation parameter (string)"}, &cli.StringFlag{Name: "basis", Usage: "Published Basis parameter (string)"}, &cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.Int64Flag{Name: "page", Usage: "Published Page parameter (int64)"}, &cli.StringFlag{Name: "before-date", Usage: "Published BeforeDate parameter (string)"}, &cli.StringFlag{Name: "as-of", Usage: "Published AsOf parameter (string)"}, &cli.BoolFlag{Name: "event-compatible-only", Usage: "Published EventCompatibleOnly parameter (bool)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetPredictions(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("prediction-class") {
					v := c.String("prediction-class")
					r = r.PredictionClass(v)
				}
				if c.IsSet("prediction-type") {
					v := c.String("prediction-type")
					r = r.PredictionType(v)
				}
				if c.IsSet("prediction-source") {
					v := c.String("prediction-source")
					r = r.PredictionSource(v)
				}
				if c.IsSet("official-limit") {
					v := c.Int64("official-limit")
					r = r.OfficialLimit(v)
				}
				if c.IsSet("pre-release-only") {
					v := c.Bool("pre-release-only")
					r = r.PreReleaseOnly(v)
				}
				if c.IsSet("seasonality") {
					v := c.String("seasonality")
					r = r.Seasonality(v)
				}
				if c.IsSet("frequency") {
					v := c.String("frequency")
					r = r.Frequency(v)
				}
				if c.IsSet("annualization") {
					v := c.String("annualization")
					r = r.Annualization(v)
				}
				if c.IsSet("period-aggregation") {
					v := c.String("period-aggregation")
					r = r.PeriodAggregation(v)
				}
				if c.IsSet("basis") {
					v := c.String("basis")
					r = r.Basis(v)
				}
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("page") {
					v := c.Int64("page")
					r = r.Page(v)
				}
				if c.IsSet("before-date") {
					v := c.String("before-date")
					r = r.BeforeDate(v)
				}
				if c.IsSet("as-of") {
					v := c.String("as-of")
					r = r.AsOf(v)
				}
				if c.IsSet("event-compatible-only") {
					v := c.Bool("event-compatible-only")
					r = r.EventCompatibleOnly(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "press-releases", Usage: "Call published GetPressReleases", ArgsUsage: "currency", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 1 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetPressReleases(ctx, c.Args().Get(0))
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "rate-differentials", Usage: "Call published GetRateDifferentials", ArgsUsage: "base quote", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "measure", Usage: "Published Measure parameter (string)"}, &cli.StringFlag{Name: "rate-type", Usage: "Published RateType parameter (string)"}, &cli.StringFlag{Name: "curve-family", Usage: "Published CurveFamily parameter (string)"}, &cli.Float64Flag{Name: "start-tenor-years", Usage: "Published StartTenorYears parameter (float64)"}, &cli.Float64Flag{Name: "end-tenor-years", Usage: "Published EndTenorYears parameter (float64)"}, &cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}, &cli.StringFlag{Name: "as-of", Usage: "Published AsOf parameter (string)"}, &cli.StringFlag{Name: "availability", Usage: "Published Availability parameter (string)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 2 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetRateDifferentials(ctx, c.Args().Get(0), c.Args().Get(1))
				if c.IsSet("measure") {
					v := c.String("measure")
					r = r.Measure(v)
				}
				if c.IsSet("rate-type") {
					v := c.String("rate-type")
					r = r.RateType(v)
				}
				if c.IsSet("curve-family") {
					v := c.String("curve-family")
					r = r.CurveFamily(v)
				}
				if c.IsSet("start-tenor-years") {
					v := c.Float64("start-tenor-years")
					r = r.StartTenorYears(v)
				}
				if c.IsSet("end-tenor-years") {
					v := c.Float64("end-tenor-years")
					r = r.EndTenorYears(v)
				}
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				if c.IsSet("as-of") {
					v := c.String("as-of")
					r = r.AsOf(v)
				}
				if c.IsSet("availability") {
					v := c.String("availability")
					r = r.Availability(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "risk-sentiment", Usage: "Call published GetRiskSentiment", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "start-date", Usage: "Published StartDate parameter (string)"}, &cli.StringFlag{Name: "end-date", Usage: "Published EndDate parameter (string)"}, &cli.Int64Flag{Name: "limit", Usage: "Published Limit parameter (int64)"}, &cli.Int64Flag{Name: "offset", Usage: "Published Offset parameter (int64)"}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.GetRiskSentiment(ctx)
				if c.IsSet("start-date") {
					v := c.String("start-date")
					r = r.StartDate(v)
				}
				if c.IsSet("end-date") {
					v := c.String("end-date")
					r = r.EndDate(v)
				}
				if c.IsSet("limit") {
					v := c.Int64("limit")
					r = r.Limit(v)
				}
				if c.IsSet("offset") {
					v := c.Int64("offset")
					r = r.Offset(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "ping", Usage: "Call published Ping", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.Ping(ctx)
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	{
		c := &cli.Command{Name: "research-panel", Usage: "Call published ResearchPanel", ArgsUsage: "", OnUsageError: usageError, ExitErrHandler: noExit, Flags: []cli.Flag{&cli.StringFlag{Name: "research-panel-request", Usage: "Published ResearchPanelRequest parameter (ResearchPanelRequest)", Required: true}}}
		c.Action = func(ctx context.Context, c *cli.Command) error {
			if c.Args().Len() != 0 {
				return fail("incorrect positional arguments; use --help")
			}
			return executeSDK(ctx, c, getenv, base, func(ctx context.Context, api *sdk.APIClient) error {
				r := api.FXMacroDataAPI.ResearchPanel(ctx)
				if c.IsSet("research-panel-request") {
					text := c.String("research-panel-request")
					if len(text) > MaxBodyBytes {
						return fail("request body exceeds limit")
					}
					var v sdk.ResearchPanelRequest
					if json.Unmarshal([]byte(text), &v) != nil {
						return fail("request body must match the published research-panel schema")
					}
					r = r.ResearchPanelRequest(v)
				}
				_, response, err := r.Execute()
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
				return err
			})
		}
		commands = append(commands, c)
	}
	return commands
}
