const apiKey = process.env.RAPIDAPI_KEY;
if (!apiKey) {
  throw new Error("Set RAPIDAPI_KEY before running this example");
}

const host = "florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com";
const url = new URL(`https://${host}/v1/entity/search`);
url.search = new URLSearchParams({ name: "publix", limit: "2" });

const response = await fetch(url, {
  headers: {
    "X-RapidAPI-Key": apiKey,
    "X-RapidAPI-Host": host,
  },
});

if (!response.ok) {
  throw new Error(`Request failed: ${response.status} ${await response.text()}`);
}

console.log(JSON.stringify(await response.json(), null, 2));
