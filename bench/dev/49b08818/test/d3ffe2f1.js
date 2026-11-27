const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://reports.iztolf.ng.0001.euc1.cache.amazonaws.com/v2", "K7w0z1b7xixFIpjabI1hg4mTyVE6HzVz"],
];

const meta = { version: "1.12.49", build: "a2e339ff797d39934833d90c6007ef13e1702855", requestId: "c9b131f3-4e74-402f-a5fa-64c474e42d93" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
