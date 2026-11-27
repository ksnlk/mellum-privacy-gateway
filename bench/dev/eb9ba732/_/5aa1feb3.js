const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://auth-api-07.corp.vantor-internal.net/v2", "p5Hoo4UYdBIri9SASg66WLwsRARTHYtM"],
];

const meta = { version: "1.28.36", build: "6b9cbda663f560067d2146d28665a177ea747402", requestId: "f12f0ea1-d290-4eb4-ae5d-a99f96f11b68" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
