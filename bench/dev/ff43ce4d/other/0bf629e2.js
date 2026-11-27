const endpoints = [
  ["https://api.github.com", null],
  ["https://admin.nordwind-corp.de/v2", "OFVJbKZ0EVuqeznnKUW2NaaiTUJuXuDx"],
];

const meta = { version: "5.12.39", build: "4072a14c6ab1d6df6ee26aebb9f5dd00dd10f8a4", requestId: "44f93308-1f4c-4976-b2e5-e3089b949bb9" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
