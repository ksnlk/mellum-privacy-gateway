const endpoints = [
  ["https://api.github.com", null],
  ["https://payments-replica.helix.cloud/v2", "sbE7TOsUfSW0eylBfyzEqcE2DDgsOgq2"],
];

const meta = { version: "1.26.50", build: "7825f2ff0dfc192ed95c5d0c14ce9c950d03256c", requestId: "50960db0-5f08-42cc-a0ee-bc0d2ff54dd3" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
