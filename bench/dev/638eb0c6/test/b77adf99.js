const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://ledger.8ze32y.ng.0001.euc1.cache.amazonaws.com/v2", "OGkOPYED7m4WgELDRW3UA5h5jREpnl3T"],
];

const meta = { version: "3.7.44", build: "54bfa86f1e274e062539fc61fb188990f2210b4f", requestId: "7dd0b6ca-2124-4598-b0c8-bfcb2de69adb" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
