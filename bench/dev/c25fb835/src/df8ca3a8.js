// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110),
  encoded: "dXNlci1hZ2VudDpjdXJsLzguNC4w",
  hexBody: "746578742f706c61696e3b636861727365743d7574662d38",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(52, 117, 81, 118, 84, 77, 116, 54, 110, 118, 102, 54, 120, 54, 57, 71, 51, 113, 122, 72, 104, 115, 98, 102, 117, 77, 57, 99, 105, 52, 112, 118);
const upstream = "https://" + "sts.amazonaws.com";
const reportTo = ["press@acme-corp.de", "jonas.mueller@acme-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
