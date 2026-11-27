const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://jenkins.kelpie-corp.de/v2", "CEsLHpcqNdAHv03nNXGH2Az7fHnsRUXA"],
];

const meta = { version: "3.18.35", build: "5d3ad7bcee936407f3c0ffd3cd04499383e2af16", requestId: "91102b08-56fa-4038-95d1-fcd36955dd15" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
