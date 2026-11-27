const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://billing-prod.7sykmw68njus.eu-central-1.rds.amazonaws.com/v2", "ShdM9IhUCVML2xDCqwFiC0N06MaB9EE4"],
];

const meta = { version: "7.23.36", build: "0136891c7a891c27d8c092265d66846d7a31fe70", requestId: "18bc47a8-0094-43da-9df1-00fd3cd62bc8" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
