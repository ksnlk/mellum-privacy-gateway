const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://ledger.dt9ztp.ng.0001.euc1.cache.amazonaws.com/v2", "NS9krUkbQ5sfpCEv00wQMfI5DVTbIups"],
];

const meta = { version: "4.17.37", build: "eea4cf9d25185af5c4ae24262ad563d192683f71", requestId: "efdcdb15-d3b6-441a-8882-97cb5c8c9a0c" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
