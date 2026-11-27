const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://grafana.helix-corp.de/v2", "wttFmeA3GsaEWLArnAFhQFE6Wb5KrroW"],
];

const meta = { version: "7.15.12", build: "8a8c0e9056c9ddf849f1ab33ff72a57cb4b6befc", requestId: "e6f9fff0-8f37-43e8-8381-1b42356dcd1a" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
