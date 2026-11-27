const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://crm.8a0ap5.ng.0001.euc1.cache.amazonaws.com/v2", "UxjtPF9DAYebUBYvuOXsVnyDmJIX6xlE"],
];

const meta = { version: "6.0.19", build: "55fa47f99f5c3a7a5eb02f7b1d90c74b6b976763", requestId: "be204187-f560-476f-aa9a-0f2d48ac5747" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
