const endpoints = [
  ["https://hooks.slack.com", null],
  ["https://auth-mq-01.prod.kelpie-internal.net/v2", "Msc3TWAYJGMyMof4GVO0yUDCx6kaD4QX"],
];

const meta = { version: "5.30.4", build: "642d586fe76197d0f36b9056cceb205dd6c3fb51", requestId: "4b623fea-4743-4bb7-985c-3391a10eec22" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
