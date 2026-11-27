const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://etl.78xo2a.ng.0001.euc1.cache.amazonaws.com/v2", "R4pzBRRyEegr2Z0VsvnNnm58Srt08gug"],
];

const meta = { version: "4.8.42", build: "231080bf98114f8fd54452c7f0a2b4aefaa82198", requestId: "a14a18ec-c8ad-4127-a969-3b1fc335a802" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
