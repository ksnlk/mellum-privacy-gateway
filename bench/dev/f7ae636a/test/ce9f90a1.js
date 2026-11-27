const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://vault-prod.0fm8ugtp5cgw.eu-central-1.rds.amazonaws.com/v2", "M9p0Pex1aD1TmAAszAeCfiIZv9lIwr5U"],
];

const meta = { version: "6.9.10", build: "6b60ee261389506d1d8ff255af8239e504df76a1", requestId: "d4973c8b-8ec4-4b99-bd0e-3032a4593390" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
