const endpoints = [
  ["https://api.github.com", null],
  ["https://argocd.helix-corp.de/v2", "U4fr2jWxaHtqZBlkfsxzDGB6rlbtXrBG"],
];

const meta = { version: "9.21.39", build: "d744b2c072c551ec9b3a27624f954ccb9c59ef75", requestId: "953e2cd3-cdee-4548-885c-4583b5e0a5b0" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
