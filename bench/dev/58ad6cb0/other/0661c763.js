const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://reports-replica.nordwind.cloud/v2", "7qaZ39h4NxqXU5tvTqRje44mAzuT8afn"],
];

const meta = { version: "6.24.17", build: "6bf2ae94401494694aec11c0833aef0c308541e9", requestId: "281b268e-31fa-443e-ba46-803317661e41" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
