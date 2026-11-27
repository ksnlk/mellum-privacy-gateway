const endpoints = [
  ["https://pypi.org", null],
  ["https://crm-replica.kelpie.cloud/v2", "oDOkDjft6E4uXnr5EbZGpTKrCutfrpkB"],
];

const meta = { version: "1.6.5", build: "bb27552209606fc6ff3eb91bd1d74d80a3998fcc", requestId: "7be14462-28bd-43c1-82bd-4e534ee47108" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
