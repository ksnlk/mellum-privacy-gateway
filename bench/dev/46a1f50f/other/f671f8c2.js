const endpoints = [
  ["https://registry.npmjs.org", null],
  ["https://jenkins.brightloop-corp.de/v2", "0tuz55PxhAjph5i8Yv47TIyJuwgAjwOa"],
];

const meta = { version: "0.26.50", build: "5e2f854060006154bfd7e55eed6ffbcc34b468b9", requestId: "da3ae6b6-c24b-4692-ac56-ab0eaa070041" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
