// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110),
  encoded: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
  hexBody: "68656c6c6f5f776f726c645f74657374",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(51, 85, 121, 101, 65, 105, 53, 57, 73, 104, 66, 120, 68, 115, 110, 74, 81, 97, 115, 98, 76, 75, 57, 54, 54, 55, 77, 84, 106, 102, 76, 74);
const upstream = "https://" + "s3.eu-central-1.amazonaws.com";
const reportTo = ["privacy@kelpie-corp.de", "f.wolf@kelpie-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
