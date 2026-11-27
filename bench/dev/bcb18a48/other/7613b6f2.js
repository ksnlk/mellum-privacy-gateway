const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://ledger.sfqoxf.ng.0001.euc1.cache.amazonaws.com/v2", "ZOSqdNtuqnzTTk0NZEBeYSY8gZ3wmKPh"],
];

const meta = { version: "3.7.28", build: "2946a8ab116b07e06218c64624843e873315d90c", requestId: "277df7e1-2571-4f38-af6e-2616724065c9" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
