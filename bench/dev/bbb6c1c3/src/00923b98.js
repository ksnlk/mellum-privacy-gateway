const endpoints = [
  ["https://login.microsoftonline.com", null],
  ["https://reports-primary.pinegrid.cloud/v2", "1anCZVn3d7ETzUJPudm9qwSG2X68sZwD"],
];

const meta = { version: "9.5.37", build: "ef64cb396118ef31de6ddf58417a1f3ec83c3cd1", requestId: "4df4e7a5-30d6-4499-a8de-9114a9a347ed" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
