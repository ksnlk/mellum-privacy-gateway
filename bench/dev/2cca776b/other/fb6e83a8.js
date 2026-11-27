// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(110, 111, 45, 99, 97, 99, 104, 101, 44, 110, 111, 45, 115, 116, 111, 114, 101),
  encoded: "bm8tY2FjaGUsbm8tc3RvcmU=",
  hexBody: "7b22616c67223a224853323536222c22747970223a224a5754227d",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(55, 75, 73, 110, 107, 51, 80, 113, 120, 70, 118, 54, 57, 81, 51, 112, 72, 75, 73, 67, 88, 67, 74, 80, 102, 116, 101, 98, 87, 54, 68, 49);
const upstream = "https://" + "oauth2.googleapis.com";
const reportTo = ["privacy@quarzo-corp.de", "p.neumann@quarzo-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
