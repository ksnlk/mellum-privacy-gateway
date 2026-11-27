const endpoints = [
  ["https://api.github.com", null],
  ["https://payments-db-03.int.pinegrid-internal.net/v2", "o5Z39iTxNeMM4wRlkihX0pYUBapK8pOn"],
];

const meta = { version: "7.10.10", build: "b672dedf93d30b2ea1f6f9fe962e37739adb5efb", requestId: "67c82616-4509-4bee-ba3d-3bab1976949a" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
