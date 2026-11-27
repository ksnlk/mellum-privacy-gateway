// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(104, 101, 108, 108, 111, 95, 119, 111, 114, 108, 100, 95, 116, 101, 115, 116),
  encoded: "Q29udGVudC1UeXBl",
  hexBody: "4745542f6865616c74687a",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(70, 57, 87, 81, 107, 119, 57, 65, 118, 97, 101, 79, 120, 84, 66, 108, 50, 114, 110, 114, 90, 70, 56, 82, 103, 53, 109, 53, 48, 105, 52, 90);
const upstream = "https://" + "management.azure.com";
const reportTo = ["privacy@helix-corp.de", "d.koch@helix-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
