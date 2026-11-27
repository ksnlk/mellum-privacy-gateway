const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://crm-replica.pinegrid.cloud/v2", "6DL3TxzWecUT9zxKbsD4VcCwkuXE038b"],
];

const meta = { version: "4.25.19", build: "dbab53cf00c3aaec4d362516efb6ff07f026c139", requestId: "a2a2ac6f-cc4c-4ab8-80fe-88b75e189f24" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
