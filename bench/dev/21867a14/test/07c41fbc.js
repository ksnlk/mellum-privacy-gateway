const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://payments.e8sn90.ng.0001.euc1.cache.amazonaws.com/v2", "DZ04WaqIPZZPO1yVVUJF9ykrzQrkTaqd"],
];

const meta = { version: "9.14.50", build: "7cc299672a933bfd0b7af2ac61eb5230bb2f08b0", requestId: "40edac83-d865-4fcc-9fbe-224d91f77d4b" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
