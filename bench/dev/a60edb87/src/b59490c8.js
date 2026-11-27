const endpoints = [
  ["https://pypi.org", null],
  ["https://kibana.pinegrid-corp.de/v2", "dJWa9M7cDrSWdnCBK1XYEn0gbBFuECLC"],
];

const meta = { version: "6.23.3", build: "025985054aeb5beca78f2c0af84d6b894e6c5b0d", requestId: "462b49d2-1f65-4614-9c52-bd35336f8bae" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
