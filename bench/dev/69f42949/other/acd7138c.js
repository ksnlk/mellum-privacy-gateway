// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(101, 110, 45, 85, 83, 44, 101, 110, 59, 113, 61, 48, 46, 57),
  encoded: "bm8tY2FjaGUsbm8tc3RvcmU=",
  hexBody: "68656c6c6f5f776f726c645f74657374",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(77, 89, 70, 68, 120, 57, 97, 98, 104, 101, 112, 117, 53, 112, 101, 117, 106, 77, 119, 76, 107, 89, 48, 78, 112, 70, 97, 50, 86, 86, 51, 120);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["security@nordwind-corp.de", "elias.klein@nordwind-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
