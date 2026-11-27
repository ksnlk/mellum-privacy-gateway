const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://vault-prod.rofd1w74x6j3.eu-central-1.rds.amazonaws.com/v2", "YsaQB6iG9V6iniQTSCMmecR3pz8U6epd"],
];

const meta = { version: "3.8.24", build: "816ed2ec35c4923a0a2a61ca442572da1f4afae4", requestId: "b341e887-9f41-4511-a988-9b16f96fea4f" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
