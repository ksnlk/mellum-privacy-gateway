const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://ledger.1in56b.ng.0001.euc1.cache.amazonaws.com/v2", "Q8GLbKUDwhcRietm23boyCz49iQ91vYX"],
];

const meta = { version: "5.24.32", build: "a1bfbfc9334524a24f214ef6308f8220a94aa0df", requestId: "97750dc8-e774-43fa-bdbe-60addfb1892b" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
