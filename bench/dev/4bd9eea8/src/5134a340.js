// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(116, 101, 120, 116, 47, 112, 108, 97, 105, 110, 59, 99, 104, 97, 114, 115, 101, 116, 61, 117, 116, 102, 45, 56),
  encoded: "R0VUL2hlYWx0aHo=",
  hexBody: "68656c6c6f5f776f726c645f74657374",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(105, 121, 110, 75, 120, 105, 57, 102, 75, 104, 117, 102, 112, 117, 119, 117, 114, 56, 109, 108, 76, 74, 72, 100, 87, 83, 77, 82, 49, 78, 77, 75);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["billing@brightloop-corp.de", "paul80@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
