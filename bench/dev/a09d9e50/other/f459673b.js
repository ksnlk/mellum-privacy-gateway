const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://payments-prod.p43rigugwyep.eu-central-1.rds.amazonaws.com/v2", "hWBGRDKYkI6Vm8kv6LaH4aVt4HqovBJB"],
];

const meta = { version: "9.9.31", build: "5840c19ffd09bc0349dd0e15cde973b347b3fcb2", requestId: "e0d7ed31-6baf-42d1-a91a-7b8643ee0650" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
