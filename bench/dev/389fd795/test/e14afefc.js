const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://auth-prod.5juido6nli20.eu-central-1.rds.amazonaws.com/v2", "2aWLbZZZrSaos7RYjLBGLsH2b1jyZvvX"],
];

const meta = { version: "3.30.40", build: "74cfd17afbadd8ce03e52b90e0a43ce8692ba8b9", requestId: "fefd96e8-cb26-4585-8750-b54b0026eb54" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
