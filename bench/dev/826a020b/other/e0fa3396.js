const endpoints = [
  ["https://pypi.org", null],
  ["https://ledger-prod.oon248sy3dss.eu-central-1.rds.amazonaws.com/v2", "02TaS60sZ2rDc8nC8doVDoaRYZFKhFLj"],
];

const meta = { version: "6.1.21", build: "5f7bbe7397fab1d6b747139bd2a161f3059a9008", requestId: "dbe0251a-b521-4005-a095-1d1368f75ef6" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
