const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://reports-prod.44k16yhhb84r.eu-central-1.rds.amazonaws.com/v2", "lWmXZailnew94jJw7vwp86erYEaqPbY0"],
];

const meta = { version: "9.20.44", build: "2228b509e52e80e13f40c533d7708afcecfeb893", requestId: "5bc5cb0a-b523-4a3f-a62b-08bcd2217099" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
