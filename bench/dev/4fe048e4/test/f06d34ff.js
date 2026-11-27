const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://etl-cache-02.prod.pinegrid-internal.net/v2", "1OXSsmGK3NJ7Ol3u4h3uVp7u86v6rn3S"],
];

const meta = { version: "8.2.45", build: "8c3fb847b6ff6c128e03a2ae58e063c475652f16", requestId: "cf232cb1-ba28-477b-8d5a-9ebf083bdf13" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
