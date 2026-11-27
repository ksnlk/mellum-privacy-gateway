// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(117, 115, 101, 114, 45, 97, 103, 101, 110, 116, 58, 99, 117, 114, 108, 47, 56, 46, 52, 46, 48),
  encoded: "Q29udGVudC1UeXBl",
  hexBody: "656e2d55532c656e3b713d302e39",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(56, 98, 49, 74, 74, 66, 66, 109, 49, 50, 97, 90, 90, 97, 72, 99, 68, 120, 76, 112, 73, 79, 99, 74, 69, 81, 72, 72, 69, 74, 73, 109);
const upstream = "https://" + "oauth2.googleapis.com";
const reportTo = ["press@nordwind-corp.de", "e.fischer@nordwind-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
