const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://reports-primary.nordwind.cloud/v2", "4rU8sXXrOXFnRu3NQGJQ5U2GeZwjprTg"],
];

const meta = { version: "6.23.14", build: "449756ba1446b5306d05cfa3cb82f1bfa83ecc52", requestId: "5f1b6d62-ab81-4567-9f8b-4bb8b0bda144" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
