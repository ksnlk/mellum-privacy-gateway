const endpoints = [
  ["https://api.github.com", null],
  ["https://auth-mq-04.corp.vantor-internal.net/v2", "f22i3eba6RJehsulw74anU41k15MvLoc"],
];

const meta = { version: "0.23.49", build: "87f566e077509cbb4fd5ff000a2c3a75030ec7fd", requestId: "5641f869-f8e5-4557-86d3-e2a519108194" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
