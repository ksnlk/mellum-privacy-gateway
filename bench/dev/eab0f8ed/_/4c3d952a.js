const endpoints = [
  ["https://api.github.com", null],
  ["https://vault.vantor-corp.de/v2", "hgbV6W0l6lAmvOojiprBiRt1lHyEW7ep"],
];

const meta = { version: "5.12.13", build: "ab6f2376e1b65c53cb5ae176d293817415dd523f", requestId: "a3fb3e11-37fa-40b7-8066-cf10b1a42c65" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
