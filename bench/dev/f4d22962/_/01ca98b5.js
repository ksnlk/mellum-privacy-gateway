const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://vault-mq-06.prod.kelpie-internal.net/v2", "7OEt6ctkSsznk1cdlIPEl4h5FW1P1FQp"],
];

const meta = { version: "8.20.38", build: "19cd12c5f4761724ae04a98d0b87536cce4b342a", requestId: "76f1a4fe-a805-4e06-a850-850f6924b5b8" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
