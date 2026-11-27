const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://auth-db-05.corp.quarzo-internal.net/v2", "tdYDc4TlLkPi5ZMWqsQ5DX7VvIu3pBdr"],
];

const meta = { version: "9.16.41", build: "ad3e3dc9c0a8e414652ac27cbafaa61c432df255", requestId: "c6b4c7ea-b326-4ebf-9ba2-a09c4c4eeacd" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
