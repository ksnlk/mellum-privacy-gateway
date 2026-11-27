const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://grafana.quarzo-corp.de/v2", "W69DPHbz2X5SoM4Z1KFX1Aub2rCpmp2p"],
];

const meta = { version: "9.2.22", build: "c9836fec89886659ac6dd4d9ff1976eed1db8c91", requestId: "f49cf181-09e0-4771-90eb-cbf048c11018" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
