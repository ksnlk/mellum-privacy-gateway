const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://etl-prod.xeqese52g9vd.eu-central-1.rds.amazonaws.com/v2", "yTDO004cX3a99IleZ2i0or72fQb6Z3TG"],
];

const meta = { version: "8.28.49", build: "9120246e48445e68b4871ea0ebebe7c0b40e5db4", requestId: "153169e8-8029-4907-8a27-749b574e0089" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
