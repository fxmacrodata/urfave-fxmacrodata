# Native capability coverage

Every capability uses native Urfave commands/flags/actions, optional customer-owned API keys and API-enforced access. `urfave.go` supplies the embeddable factory, SDK/raw request dispatch and response output; `commands_generated.go` supplies all named request builders; `client.go` supplies transport restrictions; `urfave_test.go` exercises the actual parser and request path. Status is **native** for every listed family. No operator account or data is distributed.

| FXMacroData capability | Native surface | Auth model | Implementation | Tests | Status |
| --- | --- | --- | --- | --- | --- |
| Catalogue, macro history, releases/revisions and calendar | SDK commands and public v1 request command | Optional key, server-enforced access | Factory/builders/transport | All SDK cases, no-key, raw JSON | native |
| Predictions/coverage, central-bank press releases | Full SDK builders or public JSON requests | API entitlement | Factory/builders/transport | All corresponding SDK cases | native |
| FX, reference/session/source metadata | Full SDK builders or public JSON requests | API entitlement | Factory/builders/transport | All FX SDK cases | native |
| COT, commodities, sessions and financial prices | Full SDK builders or public JSON requests | API entitlement | Factory/builders/transport | All corresponding SDK cases | native |
| Curves, factors, differentials and sentiment | Full SDK builders or public JSON requests | API entitlement | Factory/builders/transport | All corresponding SDK cases | native |
| Research panels | Published SDK model or raw JSON POST | API entitlement | Factory/builders/transport | Body/query and research builder cases | native |
| Other documented public v1 JSON, including seasonality/news where entitled | Endpoint-agnostic native request command | API entitlement | Raw request/transport | Generic complete JSON/body/query fixture | native |

Named operations and parameters:

| SDK operation | Command | Path arguments | Request flags |
| --- | --- | --- | --- |
| GetAnnouncementChanges | announcement-changes | none | --currencies, --indicators, --since, --limit, --payload |
| GetAnnouncements | announcements | currency, indicator | --start-date, --end-date, --series-mode, --limit, --offset, --page, --seasonality, --frequency, --annualization, --period-aggregation, --revisions, --basis, --official-only, --as-of, --availability, --dataset-version, --value-mode |
| GetCalendar | calendar | currency | --indicator, --start-date, --end-date, --timezone |
| GetCommodities | commodities | indicator | --start-date, --end-date, --limit, --offset, --page |
| GetCot | cot | currency | --start-date, --end-date, --limit, --offset, --page |
| GetCurves | curves | currency | --curve-family, --metric, --view, --method, --date, --as-of, --availability, --max-staleness-days, --require-common-date |
| GetDataCatalogue | data-catalogue | currency | --include-capabilities, --include-coverage, --indicator |
| GetFactor | factor | currency, factor | --start-date, --end-date, --include-components, --include-sources, --limit, --offset |
| GetFinancialPrices | financial-prices | currency | --source, --start-date, --end-date, --measure, --instrument-id, --issuer, --limit, --offset |
| GetForex | forex | base, quote | --start-date, --end-date, --limit, --offset, --page, --indicators, --dataset-version |
| GetFxIntradayReferenceRates | fx-intraday-reference-rates | base, quote | --start-time, --end-time |
| GetFxReferenceRates | fx-reference-rates | base, quote | --date, --start-date, --end-date, --source, --timestamp-type, --native-only, --include-source-legs, --order, --limit, --offset, --page, --dataset-version |
| GetFxSessionRates | fx-session-rates | base, quote | --session, --date, --start-date, --end-date, --order, --limit, --offset, --page, --dataset-version |
| GetFxSourceUniverse | fx-source-universe | none | --currency, --source |
| GetFxSources | fx-sources | none | none |
| GetHealth | health | none | none |
| GetLatestAnnouncements | latest-announcements | currency | none |
| GetLatestCommodities | latest-commodities | none | none |
| GetMarketSessions | market-sessions | none | --at |
| GetPredictionCoverage | prediction-coverage | none | none |
| GetPredictionCoverageForCurrency | prediction-coverage-for-currency | currency | none |
| GetPredictions | predictions | currency, indicator | --prediction-class, --prediction-type, --prediction-source, --official-limit, --pre-release-only, --seasonality, --frequency, --annualization, --period-aggregation, --basis, --start-date, --end-date, --limit, --offset, --page, --before-date, --as-of, --event-compatible-only |
| GetPressReleases | press-releases | currency | --limit, --offset |
| GetRateDifferentials | rate-differentials | base, quote | --measure, --rate-type, --curve-family, --start-tenor-years, --end-tenor-years, --start-date, --end-date, --limit, --offset, --as-of, --availability |
| GetRiskSentiment | risk-sentiment | none | --start-date, --end-date, --limit, --offset |
| Ping | ping | none | none |
| ResearchPanel | research-panel | none | --research-panel-request |

There are 25 required path arguments and 131 query/body setters. Research-panel's JSON flag is required. The generic transport test does not certify each live endpoint; use exact published API documentation for paths and shapes. Streaming, binary/non-JSON formats and automatic pagination are outside this JSON client. No data is reshaped or combined to fit a reduced model.
