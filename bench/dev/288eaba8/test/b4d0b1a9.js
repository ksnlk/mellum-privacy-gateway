const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://payments.ttafcl.ng.0001.euc1.cache.amazonaws.com/v2", "iypMJxl02amgSjun21uqoIUVvyfAfZX4"],
];

const meta = { version: "2.0.39", build: "17ac7d29e3567520495b9c1538c7c9b348a442f3", requestId: "26ccdd78-3f5f-4777-87c2-a98dfea4a55f" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
