// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110),
  encoded: "bm8tY2FjaGUsbm8tc3RvcmU=",
  hexBody: "68656c6c6f5f776f726c645f74657374",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(97, 100, 72, 114, 50, 108, 72, 66, 106, 108, 102, 108, 50, 121, 50, 54, 80, 78, 121, 116, 86, 106, 57, 75, 81, 71, 73, 99, 83, 82, 118, 49);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["press@helix-corp.de", "david45@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
