// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101),
  encoded: "dXNlci1hZ2VudDpjdXJsLzguNC4w",
  hexBody: "4745542f6865616c74687a",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(120, 110, 71, 70, 66, 106, 65, 103, 79, 110, 79, 118, 98, 103, 112, 65, 109, 79, 84, 57, 81, 122, 55, 100, 98, 113, 56, 57, 108, 48, 53, 122);
const upstream = "https://" + "management.azure.com";
const reportTo = ["privacy@brightloop-corp.de", "felix89@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
