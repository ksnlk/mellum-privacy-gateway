const endpoints = [
  ["https://api.github.com", null],
  ["https://billing-prod.wsloo3bqmdzo.eu-central-1.rds.amazonaws.com/v2", "GqVdSZveluiGlvOVBo6h9nA7x02QpbSa"],
];

const meta = { version: "0.11.8", build: "d355196830cb9d760e8432a17bf265c9299a1655", requestId: "bbbe819d-fa83-46be-8a12-49e21bd9d040" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
