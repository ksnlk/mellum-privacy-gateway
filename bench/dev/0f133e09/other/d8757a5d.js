const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://billing-primary.kelpie.cloud/v2", "v037wzUz6o46ibbdYgwSMWFz2oajnVw3"],
];

const meta = { version: "4.3.31", build: "6f1dac2df98b44e3ad15b230954b7fa1a0497ede", requestId: "a4787dbf-75b1-430e-bf01-52db55b6a514" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
