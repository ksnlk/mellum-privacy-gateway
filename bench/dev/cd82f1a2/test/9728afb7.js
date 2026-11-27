const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://reports-primary.helix.cloud/v2", "MDVwS0tucU8wSK6eFccBqvkAprxZo1w8"],
];

const meta = { version: "3.17.31", build: "296d0a691a395f6abb6f6899026e4036d7908460", requestId: "0361b8d1-7fd8-48aa-9ca2-4e9ddecb6e49" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
