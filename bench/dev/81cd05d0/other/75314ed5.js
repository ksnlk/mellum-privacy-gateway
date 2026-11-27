// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(117, 115, 101, 114, 45, 97, 103, 101, 110, 116, 58, 99, 117, 114, 108, 47, 56, 46, 52, 46, 48),
  encoded: "R0VUL2hlYWx0aHo=",
  hexBody: "436f6e74656e742d54797065",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(84, 80, 82, 100, 117, 66, 101, 99, 55, 109, 77, 52, 71, 106, 110, 81, 86, 52, 86, 116, 70, 116, 120, 108, 90, 65, 113, 54, 83, 120, 122, 51);
const upstream = "https://" + "management.azure.com";
const reportTo = ["privacy@acme-corp.de", "e.fischer@acme-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
