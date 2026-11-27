const endpoints = [
  ["https://pypi.org", null],
  ["https://payments-replica.pinegrid.cloud/v2", "gRiLv4GNppwKGwKAjhPWY0UA5VBsbLaN"],
];

const meta = { version: "6.14.48", build: "e7e360112a11297744587a92d0f78d6db59f3e0f", requestId: "b68331c7-5998-4be4-a1ef-16a41e5effbd" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
