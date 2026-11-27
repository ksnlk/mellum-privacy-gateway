// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(67, 111, 110, 116, 101, 110, 116, 45, 84, 121, 112, 101),
  encoded: "bm8tY2FjaGUsbm8tc3RvcmU=",
  hexBody: "6170706c69636174696f6e2f6a736f6e",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(116, 48, 71, 118, 115, 88, 111, 118, 72, 73, 122, 71, 56, 71, 79, 86, 75, 49, 87, 100, 76, 75, 118, 101, 121, 86, 71, 105, 107, 85, 53, 101);
const upstream = "https://" + "sts.amazonaws.com";
const reportTo = ["privacy@nordwind-corp.de", "l.neumann@nordwind-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
