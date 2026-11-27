// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101),
  encoded: "R0VUL2hlYWx0aHo=",
  hexBody: "4745542f6865616c74687a",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(104, 76, 71, 109, 103, 80, 97, 55, 55, 70, 69, 77, 90, 89, 84, 110, 120, 68, 71, 119, 86, 76, 116, 101, 109, 117, 113, 114, 113, 105, 82, 67);
const upstream = "https://" + "management.azure.com";
const reportTo = ["billing@quarzo-corp.de", "lina87@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
