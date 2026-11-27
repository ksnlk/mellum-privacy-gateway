const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://crm-prod.vf3s2yhd40eq.eu-central-1.rds.amazonaws.com/v2", "DOXPSbZ8RdFqPqDBbQ50I5kudqtK4lGc"],
];

const meta = { version: "6.28.17", build: "da69d91661738934ebbce498d1c14db102c4190d", requestId: "eb002bb7-cf77-4023-977b-91b274859b84" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
