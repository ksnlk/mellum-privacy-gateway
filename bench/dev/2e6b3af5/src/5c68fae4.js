// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(104, 101, 108, 108, 111, 95, 119, 111, 114, 108, 100, 95, 116, 101, 115, 116),
  encoded: "YXBwbGljYXRpb24vanNvbg==",
  hexBody: "4745542f6865616c74687a",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(113, 84, 79, 88, 84, 109, 71, 107, 89, 54, 74, 71, 79, 99, 117, 53, 74, 111, 106, 121, 75, 76, 114, 69, 74, 84, 119, 73, 122, 74, 117, 102);
const upstream = "https://" + "s3.eu-central-1.amazonaws.com";
const reportTo = ["careers@kelpie-corp.de", "s.zimmermann@kelpie-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
