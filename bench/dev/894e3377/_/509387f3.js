// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(101, 110, 45, 85, 83, 44, 101, 110, 59, 113, 61, 48, 46, 57),
  encoded: "dXNlci1hZ2VudDpjdXJsLzguNC4w",
  hexBody: "4745542f6865616c74687a",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(86, 69, 67, 73, 108, 114, 74, 53, 98, 71, 82, 73, 69, 65, 66, 51, 98, 49, 79, 68, 118, 102, 98, 81, 118, 106, 76, 117, 77, 74, 112, 81);
const upstream = "https://" + "management.azure.com";
const reportTo = ["privacy@helix-corp.de", "noah65@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
