const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://vault-prod.85lmyj6ehsrl.eu-central-1.rds.amazonaws.com/v2", "2dNkBMRfYLvNdCAcxq9ON5ibxTA1g8SP"],
];

const meta = { version: "8.16.43", build: "c159b8f84cadf051ac10534c23852930119013bb", requestId: "c877187a-699d-4d73-ba9a-20e30dfe0d50" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
