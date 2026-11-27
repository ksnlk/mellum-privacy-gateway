const endpoints = [
  ["https://api.github.com", null],
  ["https://crm-cache-08.int.kelpie-internal.net/v2", "E2rOOFBS2NInt4OmNAoNvb5qibLbe8ii"],
];

const meta = { version: "6.7.27", build: "374714497d95a83afe60fda8d599e126956bead6", requestId: "0db6cd90-3b8d-4dce-a695-bfe7d874e2ec" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
