const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://ledger-prod.y1fi2ane5d1b.eu-central-1.rds.amazonaws.com/v2", "4aI9oiy7hw4WzrYtE8MtuFV6z4bcKwgh"],
];

const meta = { version: "4.4.0", build: "9bd90e4a10cde45d42a58ec39b37ea882ee85c42", requestId: "f818dfe2-3198-4cd0-91a9-b727f4a6eb08" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
