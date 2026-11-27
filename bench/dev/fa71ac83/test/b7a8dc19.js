const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://vpn-gw.acme-corp.de/v2", "UmhrJ60JciAA0P6f4kYL3qg6dgXfULAI"],
];

const meta = { version: "9.4.39", build: "f2a4f14ad05a54cb4fd11d4d9cc2a0ff2a88f2b9", requestId: "9ba58266-7de0-42ec-a2ab-d62e64f49090" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
