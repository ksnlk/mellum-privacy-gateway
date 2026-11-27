const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://reports.m4pu8x.ng.0001.euc1.cache.amazonaws.com/v2", "XvT4C9khia3sqYxYrWdndWTmShuaVkuQ"],
];

const meta = { version: "5.22.26", build: "41de8f031d12cd8cd1682e250182ca3a066a5343", requestId: "05ece6c0-ed10-494c-84aa-f4ad83ee8161" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
