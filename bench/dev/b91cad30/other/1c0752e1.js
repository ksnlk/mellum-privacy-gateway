const endpoints = [
  ["https://api.github.com", null],
  ["https://argocd.pinegrid-corp.de/v2", "eFFLIsSQAsrRQGANMDH2s5vq8cTzX66J"],
];

const meta = { version: "0.13.22", build: "eaf0904be593a07efb50efd0afa8537eec87776f", requestId: "93caa7de-eeb0-4f51-b652-70add2401338" };

export function client(i) {
  const [base, k] = endpoints[i];
  return fetch(base + "/status", { headers: k ? { "X-Key": k } : {} });
}
