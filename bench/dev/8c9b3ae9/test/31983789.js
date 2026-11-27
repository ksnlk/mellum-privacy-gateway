const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://argocd.acme-corp.de/v2", "rmC5q5RXBcnjgHdAyoAliCY17fhfQtZx"],
];

const meta = { version: "0.13.31", build: "740ad2e32d9198d5916cc3cbcc5d785535705906", requestId: "c5341a74-9d75-4ff4-843b-70cb3ef07013" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
