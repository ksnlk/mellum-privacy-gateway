const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://etl-api-07.corp.kelpie-internal.net/v2", "L9T439vDlu9dktXvmqmJxoRtyhDLQGQD"],
];

const meta = { version: "5.22.11", build: "dce6d6368a05330253accaf5474c416ec9a72f5b", requestId: "a72e001e-4580-4e61-a71f-51626978a6f2" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
