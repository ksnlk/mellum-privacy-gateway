const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://auth-prod.4g2peaaarn1y.eu-central-1.rds.amazonaws.com/v2", "rDNC5Ig8oNWoDNA7zxkFGHUhruzX5huw"],
];

const meta = { version: "3.18.11", build: "948291b553d8c5096a04bd6b797117a713245baa", requestId: "26516161-b066-4769-906a-54735112fc40" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
