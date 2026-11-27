const endpoints = [
  ["https://api.github.com", null],
  ["https://etl-api-02.int.vantor-internal.net/v2", "prS7iwKwburqHIdQLXHJw56fIgpVD5HG"],
];

const meta = { version: "4.26.13", build: "5ff7ea388c60b88c2ac23ae501b4451dd584eb7d", requestId: "4a91f92c-b31e-47a4-b73d-db213ed42ed4" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
