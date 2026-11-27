const endpoints = [
  ["https://api.github.com", null],
  ["https://etl-prod.rc1rwr8enld4.eu-central-1.rds.amazonaws.com/v2", "8fVoFYPgCMmkFpDQygTVT49naJI7bap2"],
];

const meta = { version: "5.11.2", build: "e816e9354ce18c93dad0c129629e018902d16d0a", requestId: "a6276209-017a-48b9-a4dd-0d41568baef4" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
