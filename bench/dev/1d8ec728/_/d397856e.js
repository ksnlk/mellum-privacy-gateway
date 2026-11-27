// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(71, 69, 84, 47, 104, 101, 97, 108, 116, 104, 122),
  encoded: "dXNlci1hZ2VudDpjdXJsLzguNC4w",
  hexBody: "656e2d55532c656e3b713d302e39",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(88, 76, 120, 106, 75, 49, 48, 50, 117, 111, 114, 89, 103, 97, 51, 48, 103, 55, 54, 80, 73, 88, 103, 55, 70, 73, 79, 107, 76, 115, 121, 83);
const upstream = "https://" + "s3.eu-central-1.amazonaws.com";
const reportTo = ["billing@quarzo-corp.de", "lukas.hoffmann@quarzo-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
