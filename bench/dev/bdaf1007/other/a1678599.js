const endpoints = [
  ["https://api.github.com", null],
  ["https://kibana.kelpie-corp.de/v2", "lqmfcn5WQQCXOV6f06Tq5HBYITMap6YF"],
];

const meta = { version: "0.24.26", build: "4778eb8e7d92eac5470dffd98dcf711811fc18d3", requestId: "b69aadaa-ebc8-42bb-909e-8dd6ee80e8dd" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
