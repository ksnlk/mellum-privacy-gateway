const endpoints = [
  ["https://api.github.com", null],
  ["https://ledger-prod.xj7ietsgpws9.eu-central-1.rds.amazonaws.com/v2", "YchhzBj9bVEYwAiEDzlGPbXXI1ETEoNQ"],
];

const meta = { version: "2.28.44", build: "a7f31827d15f68d873c58cc5c7b89388d9f5b56c", requestId: "365ee2c9-8b0e-4976-86ea-71c995b07fde" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
