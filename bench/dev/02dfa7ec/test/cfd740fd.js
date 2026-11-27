const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://etl.7fenbz.ng.0001.euc1.cache.amazonaws.com/v2", "ONkcc35wupltYb7F7v4tgYc1Kdq6V4FN"],
];

const meta = { version: "9.21.5", build: "b19115f87069410edc7ff25167fdcf675538eee9", requestId: "85dbb816-4a5e-4d41-b92b-114b1b0930f7" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
