const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://crm.3b7v6z.ng.0001.euc1.cache.amazonaws.com/v2", "kap1nsPFntUopqd6L7td378X8elWI7RT"],
];

const meta = { version: "8.0.2", build: "9592a8ac1e0a99998607090447142cfb53e11d13", requestId: "5180b970-fcee-4733-b3ac-5277ded54efe" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
