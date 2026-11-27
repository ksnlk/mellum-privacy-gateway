const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://vault-replica.helix.cloud/v2", "TMiZStTF7SKgo1wIGbjgK6zTaoK5DaUW"],
];

const meta = { version: "3.22.34", build: "60150a2a5919b09c1a689e4c8d3bf3e04adfa58d", requestId: "3eb34f1c-f406-45b0-b9ad-65540cac517c" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
