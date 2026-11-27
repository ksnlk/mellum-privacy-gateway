const endpoints = [
  ["https://s3.amazonaws.com", null],
  ["https://crm.uhdcqy.ng.0001.euc1.cache.amazonaws.com/v2", "x2FRFy8sruUJWtlQSgNirIHJMdOPqWme"],
];

const meta = { version: "8.12.8", build: "2ee7ebfc7ba027b79237cb7eafddf285ad5a96f1", requestId: "2ccd22a2-e008-421d-af60-0361c7243060" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
