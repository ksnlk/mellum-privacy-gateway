const endpoints = [
  ["https://pypi.org", null],
  ["https://ledger.hu98wk.ng.0001.euc1.cache.amazonaws.com/v2", "cvSQCPPIAwbyiuG1qyOTIXFu7uoH5SrE"],
];

const meta = { version: "4.19.42", build: "b2b0d7bb01779e342c12f35ba21f96ea0107d5a3", requestId: "d66c542b-0f29-4b53-9825-4f532bce7d6b" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
