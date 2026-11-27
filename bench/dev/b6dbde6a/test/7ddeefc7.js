// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(116, 101, 120, 116, 47, 112, 108, 97, 105, 110, 59, 99, 104, 97, 114, 115, 101, 116, 61, 117, 116, 102, 45, 56),
  encoded: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
  hexBody: "746578742f706c61696e3b636861727365743d7574662d38",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(75, 66, 51, 85, 72, 72, 75, 119, 54, 76, 99, 108, 101, 113, 119, 97, 105, 117, 113, 108, 80, 49, 122, 113, 66, 74, 78, 50, 105, 119, 119, 70);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["careers@quarzo-corp.de", "e.weber@quarzo-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
