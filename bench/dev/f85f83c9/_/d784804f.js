const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://billing-prod.4lhi40rf5wvq.eu-central-1.rds.amazonaws.com/v2", "rCDsH87HxD32hrrhluGayl5JSfQPC4k9"],
];

const meta = { version: "0.26.37", build: "8f47c499bfa6995bd6a60472a3bfe44bb385d600", requestId: "e71d0173-77ac-4e70-ac92-98969839e6d8" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
