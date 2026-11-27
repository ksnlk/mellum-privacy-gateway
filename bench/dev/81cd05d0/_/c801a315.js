const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://payments-db-09.corp.nordwind-internal.net/v2", "KpImzXcLwhGooOtjK3kPoKiSr3SZmoog"],
];

const meta = { version: "5.22.9", build: "1e0cd58aa74f92a4b8d1b315a9da19a67433ec82", requestId: "f791f7e6-94d3-4cb7-9f0c-67e064eb70b8" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
