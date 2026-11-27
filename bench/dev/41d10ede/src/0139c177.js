const endpoints = [
  ["https://api.github.com", null],
  ["https://auth-primary.brightloop.cloud/v2", "dj89zE9xQ3l2MafPhV8BZEN5tqxmIdAB"],
];

const meta = { version: "3.29.48", build: "4858a4868202e9ae541788d10dadb13753bc6e25", requestId: "53e9fe42-96ee-49d0-bbc1-c6aae53ced5c" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
