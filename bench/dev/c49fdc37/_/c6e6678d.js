const endpoints = [
  ["https://pypi.org", null],
  ["https://billing.vnnggv.ng.0001.euc1.cache.amazonaws.com/v2", "piTva3QyjloLSeFIT1tkn2EuvVszRThL"],
];

const meta = { version: "0.15.4", build: "9697594704d1020dd8ee422d51f8708deb1891c3", requestId: "5d83e60a-522b-453f-8940-1f74fca74546" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
