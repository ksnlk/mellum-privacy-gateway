const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://auth-primary.vantor.cloud/v2", "6EwQufuV87AHcfDSstsrBMbehmisKExp"],
];

const meta = { version: "8.26.36", build: "6d752d5ce92e009dcd7ab3616925022cdbffb790", requestId: "47ca9b00-4f49-43c2-924b-ffe48d8c0761" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
