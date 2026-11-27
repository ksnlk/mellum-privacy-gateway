const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://crm-replica.helix.cloud/v2", "uWPOC2a54gdQEbliyBfz2kU0OB2rAfK6"],
];

const meta = { version: "0.8.42", build: "eb131c04eaf23893f1df52f04e682ba98b538c56", requestId: "972e7ac0-d4bd-4100-9490-410fd26c0060" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
