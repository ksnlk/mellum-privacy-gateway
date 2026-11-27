const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://crm-primary.brightloop.cloud/v2", "krOQNYVMXnEreLtKDSwttrc7tz2ghv50"],
];

const meta = { version: "8.12.49", build: "e87ecd716e2f2abbb752cad861faa6d2a131a233", requestId: "3e52b490-bfb2-437b-8ef3-bcb3e92cab53" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
