#!/usr/bin/env python3
import json
import os
import urllib.parse
import urllib.request

api_key = os.environ.get("RAPIDAPI_KEY")
if not api_key:
    raise SystemExit("Set RAPIDAPI_KEY before running this example")

host = "florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com"
query = urllib.parse.urlencode({"name": "publix", "limit": 2})
request = urllib.request.Request(
    f"https://{host}/v1/entity/search?{query}",
    headers={
        "X-RapidAPI-Key": api_key,
        "X-RapidAPI-Host": host,
    },
)

with urllib.request.urlopen(request, timeout=10) as response:
    print(json.dumps(json.load(response), indent=2))
