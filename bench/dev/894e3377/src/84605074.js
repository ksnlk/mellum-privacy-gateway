const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://auth-primary.kelpie.cloud/v2", "Xi1bKsSWv8xiXXyYPXyzoP4eax7EO95u"],
];

const meta = { version: "8.15.3", build: "942d842fa560cfd018ea52df5929776200c36c9d", requestId: "0869a7bf-76d2-48e8-a139-88e4e0efc1be" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
