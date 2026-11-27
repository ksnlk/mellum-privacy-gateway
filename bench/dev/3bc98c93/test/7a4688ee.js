const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://payments-replica.vantor.cloud/v2", "zvkGtpkaOEZLxIb1tyKZGxzpXcup2WEB"],
];

const meta = { version: "2.9.10", build: "64c3cfa5983da8c8c0b7f9e13c04867a66824509", requestId: "ac1a6ef3-d35a-4bca-8200-03b7c0d2ba38" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
