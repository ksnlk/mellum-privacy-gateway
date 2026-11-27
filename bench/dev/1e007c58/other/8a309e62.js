const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://vault-prod.ujia82m15lfq.eu-central-1.rds.amazonaws.com/v2", "bMneVoucqZMZ83UQQ6OGo4IPDAxUcqec"],
];

const meta = { version: "3.15.41", build: "315253255e0367f25b43de94d1c889e63398834e", requestId: "81b0a1d0-133d-43da-a92c-f4a6cf1d059f" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
