const endpoints = [
  ["https://api.github.com", null],
  ["https://reports-prod.0omre88zrwf2.eu-central-1.rds.amazonaws.com/v2", "YgBYgwGtp4bfJxsLOnNYJ3EjTz7a0Aou"],
];

const meta = { version: "1.12.10", build: "2dbe1f89e9ae38b452fe1c0a0ee6c746d761643c", requestId: "363d94eb-5ad3-4072-a529-25e4ef78431e" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
