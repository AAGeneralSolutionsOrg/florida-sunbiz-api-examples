# Florida Sunbiz Entity & Officer Lookup API

Search Florida business entities and retrieve filing details, registered agents, officers, and managers through RapidAPI.

**[Open the API on RapidAPI](https://rapidapi.com/aa-general-solutions-aa-general-solutions-default/api/florida-sunbiz-entity-and-officer-lookup)** · **[View the product page](https://sunbiz-api.aageneralsolutions.com)**

## Plans

| Plan | Monthly requests | Price |
|---|---:|---:|
| Basic | 100 | Free |
| Pro | 2,500 | $25/month |
| Ultra | 15,000 | $75/month |
| Mega | 50,000 | $150/month |

Every plan has a hard monthly limit, so usage cannot create an unexpected overage charge.

## What you can query

- Search entities by corporate name
- Retrieve a filing by Florida document number
- Read filing status and filing date
- Read principal and mailing addresses
- Read registered-agent information
- Read available officers and managers

The service is backed by the Florida Division of Corporations public bulk dataset and is synchronized periodically. It does not scrape the Sunbiz website during API requests.

## Authentication

Subscribe to a plan on RapidAPI and copy your application key. Set it as an environment variable so it never appears in source control:

```bash
export RAPIDAPI_KEY="your-rapidapi-key"
```

RapidAPI requests require these headers:

```text
X-RapidAPI-Key: your-rapidapi-key
X-RapidAPI-Host: florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com
```

## Search entities

```http
GET /v1/entity/search?name=publix&limit=2
```

At least one search criterion is required. `limit` is optional, defaults to 20, and accepts values from 1 through 50.

```bash
curl --request GET \
  --url 'https://florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com/v1/entity/search?name=publix&limit=2' \
  --header "X-RapidAPI-Key: $RAPIDAPI_KEY" \
  --header 'X-RapidAPI-Host: florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com'
```

Example response:

```json
[
  {
    "document_number": "P01000085676",
    "entity_name": "PUBLIX ADVERTISING, INC.",
    "status": "Inactive",
    "filing_date": "08/29/2001"
  }
]
```

### Search criteria

The search endpoint accepts any of these identity criteria:

| Parameter | Match type | Example |
|---|---|---|
| `name` | Prefix | `publix` |
| `document_number` | Exact | `112252` |
| `fei_ein_number` | Exact | `59-0324412` or `590324412` |
| `registered_agent` | Prefix | `corporate creations` |
| `officer` | Prefix | `murphy` |
| `status` | Optional filter | `Active` or `Inactive` |

At least one identity criterion is required. You can combine criteria; combined filters use AND semantics.

```bash
# Exact FEI/EIN lookup
curl --get \
  --url 'https://florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com/v1/entity/search' \
  --data-urlencode 'fei_ein_number=59-0324412' \
  --header "X-RapidAPI-Key: $RAPIDAPI_KEY" \
  --header 'X-RapidAPI-Host: florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com'

# Active entities associated with an officer or manager name prefix
curl --get \
  --url 'https://florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com/v1/entity/search' \
  --data-urlencode 'officer=murphy' \
  --data-urlencode 'status=Active' \
  --header "X-RapidAPI-Key: $RAPIDAPI_KEY" \
  --header 'X-RapidAPI-Host: florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com'
```

## Get entity details

```http
GET /v1/entity/{doc_number}
```

```bash
curl --request GET \
  --url 'https://florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com/v1/entity/112252' \
  --header "X-RapidAPI-Key: $RAPIDAPI_KEY" \
  --header 'X-RapidAPI-Host: florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com'
```

Example response (abbreviated):

```json
{
  "document_number": "112252",
  "corporate_name": "PUBLIX SUPER MARKETS, INC.",
  "fei_ein_number": "59-0324412",
  "status": "Active",
  "filing_date": "12/27/1921",
  "principal_address": "3300 PUBLIX CORPORATE PKWY\nLAKELAND, FL 33811-3311",
  "mailing_address": "P.O. BOX 407\nLAKELAND, FL 33802",
  "registered_agent": {
    "name": "Corporate Creations Network Inc.",
    "address": "801 US Highway 1\nNorth Palm Beach, FL 33408"
  },
  "officers": [
    {
      "title": "SVP",
      "name": "Example Officer",
      "address": "Lakeland, FL"
    }
  ]
}
```

## Runnable examples

- [cURL](examples/curl.sh)
- [JavaScript](examples/javascript.mjs) — Node.js 18 or newer
- [Python](examples/python.py) — Python 3, no external packages
- [Go](examples/go/main.go) — Go 1.22 or newer

## Errors

| Status | Meaning |
|---:|---|
| 400 | Invalid query or path parameter |
| 401 | Missing or invalid RapidAPI authentication |
| 404 | Document number not found |
| 429 | RapidAPI plan quota or rate limit exceeded |
| 500 | Internal processing error |
| 503 | Local data index unavailable |

## Data freshness

The unauthenticated health endpoint reports the timestamp and filename of the indexed source data:

```bash
curl 'https://api-sunbiz.aageneralsolutions.com/health'
```

```json
{
  "status": "ok",
  "timestamp": "2026-09-30T14:43:40Z",
  "data_updated_at": "2026-09-30T14:43:40Z",
  "last_source": "20260929c.txt"
}
```

## Security

Never commit an `X-RapidAPI-Key`. Use an environment variable or a secrets manager. If a key is exposed, rotate it from your RapidAPI application dashboard.

## Data and disclaimer

Data originates from public Florida Division of Corporations bulk files. This independent API is not affiliated with or endorsed by the Florida Department of State. Data is provided as-is; verify critical information against official records.

## License

The example code in this repository is available under the [MIT License](LICENSE).
