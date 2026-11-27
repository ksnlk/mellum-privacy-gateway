const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://auth-replica.brightloop.cloud/v2", "nopgw5ajNtQyzCgcoSLG2psPq7ivhYNF"],
];

const meta = { version: "1.10.8", build: "98f34ec1264286e35deef096d72b9a39774c7810", requestId: "267fdc64-d9da-4621-96bc-eb2ec96895e7" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
