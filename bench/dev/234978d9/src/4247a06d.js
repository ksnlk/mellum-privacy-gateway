// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(110, 111, 45, 99, 97, 99, 104, 101, 44, 110, 111, 45, 115, 116, 111, 114, 101),
  encoded: "dXNlci1hZ2VudDpjdXJsLzguNC4w",
  hexBody: "656e2d55532c656e3b713d302e39",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(70, 84, 70, 86, 102, 120, 48, 87, 110, 67, 89, 115, 109, 86, 68, 121, 57, 106, 73, 77, 112, 122, 48, 86, 101, 108, 79, 114, 106, 122, 86, 84);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["hello@acme-corp.de", "emma.braun@acme-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
