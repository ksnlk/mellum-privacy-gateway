	a.GetID()
}

func TestAuthorization_GetNote(tt *testing.T) {
	var zeroValue string
	a := &Authorization{Note: &zeroValue}
	a.GetNote()
	a = &Authorization{}
	a.GetNote()
	a = nil
	a.GetNote()
}

func TestAuthorization_GetNoteURL(tt *testing.T) {
	var zeroValue string
	a := &Authorization{NoteURL: &zeroValue}
	a.GetNoteURL()
	a = &Authorization{}
	a.GetNoteURL()
	a = nil
	a.GetNoteURL()
}

func TestAuthorization_GetToken(tt *testing.T) {
	var zeroValue string
	a := &Authorization{Token: &zeroValue}
	a.GetToken()
	a = &Authorization{}
	a.GetToken()
	a = nil
	a.GetToken()
}

func TestAuthorization_GetTokenLastEight(tt *testing.T) {
	var zeroValue string
	a := &Authorization{TokenLastEight: &zeroValue}
	a.GetTokenLastEight()
	a = &Authorization{}
	a.GetTokenLastEight()
	a = nil
	a.GetTokenLastEight()
}

func TestAuthorization_GetUpdatedAt(tt *testing.T) {
	var zeroValue Timestamp
	a := &Authorization{UpdatedAt: &zeroValue}
	a.GetUpdatedAt()
	a = &Authorization{}
	a.GetUpdatedAt()
	a = nil
	a.GetUpdatedAt()
	c := &CheckSuitePreferenceResults{}
	c.GetPreferences()
	c = nil
	c.GetPreferences()
}

func TestCheckSuitePreferenceResults_GetRepository(tt *testing.T) {
	c := &CheckSuitePreferenceResults{}
	c.GetRepository()
	c = nil
	c.GetRepository()
}

func TestCodeOfConduct_GetBody(tt *testing.T) {
	var zeroValue string
	c := &CodeOfConduct{Body: &zeroValue}
	c.GetBody()
	c = &CodeOfConduct{}
	c.GetBody()
	c = nil
	c.GetBody()
}

func TestCodeOfConduct_GetKey(tt *testing.T) {
	var zeroValue string
	c := &CodeOfConduct{Key: &zeroValue}
	c.GetKey()
	c = &CodeOfConduct{}
	c.GetKey()
	c = nil
	c.GetKey()
}

func TestCodeOfConduct_GetName(tt *testing.T) {
	var zeroValue string
	c := &CodeOfConduct{Name: &zeroValue}
	c.GetName()
	c = &CodeOfConduct{}
	c.GetName()
	c = nil
	c.GetName()
}

func TestCodeOfConduct_GetURL(tt *testing.T) {
	var zeroValue string
	c := &CodeOfConduct{URL: &zeroValue}
	c.GetURL()
	c = &CodeOfConduct{}
	c.GetURL()
	c = nil
	c.GetURL()
	c = &Commit{}
	c.GetURL()
	c = nil
	c.GetURL()
}

func TestCommit_GetVerification(tt *testing.T) {
	c := &Commit{}
	c.GetVerification()
	c = nil
	c.GetVerification()
}

func TestCommitAuthor_GetDate(tt *testing.T) {
	var zeroValue time.Time
	c := &CommitAuthor{Date: &zeroValue}
	c.GetDate()
	c = &CommitAuthor{}
	c.GetDate()
	c = nil
	c.GetDate()
}

func TestCommitAuthor_GetEmail(tt *testing.T) {
	var zeroValue string
	c := &CommitAuthor{Email: &zeroValue}
	c.GetEmail()
	c = &CommitAuthor{}
	c.GetEmail()
	c = nil
	c.GetEmail()
}

func TestCommitAuthor_GetLogin(tt *testing.T) {
	var zeroValue string
	c := &CommitAuthor{Login: &zeroValue}
	c.GetLogin()
	c = &CommitAuthor{}
	c.GetLogin()
	c = nil
	c.GetLogin()
}

func TestCommitAuthor_GetName(tt *testing.T) {
	var zeroValue string
	c := &CommitAuthor{Name: &zeroValue}
	c.GetName()
	c = &CommitAuthor{}
	c.GetName()
	c = nil
	c.GetName()
}

func TestCommitCommentEvent_GetAction(tt *testing.T) {
	var zeroValue string
	c := &CommitCommentEvent{Action: &zeroValue}
	c.GetAction()
	c = &CommitCommentEvent{}
	c.GetAction()
	c = nil
	c.GetAction()
	i.GetURL()
}

func TestImport_GetUseLFS(tt *testing.T) {
	var zeroValue string
	i := &Import{UseLFS: &zeroValue}
	i.GetUseLFS()
	i = &Import{}
	i.GetUseLFS()
	i = nil
	i.GetUseLFS()
}

func TestImport_GetVCS(tt *testing.T) {
	var zeroValue string
	i := &Import{VCS: &zeroValue}
	i.GetVCS()
	i = &Import{}
	i.GetVCS()
	i = nil
	i.GetVCS()
}

func TestImport_GetVCSPassword(tt *testing.T) {
	var zeroValue string
	i := &Import{VCSPassword: &zeroValue}
	i.GetVCSPassword()
	i = &Import{}
	i.GetVCSPassword()
	i = nil
	i.GetVCSPassword()
}

func TestImport_GetVCSURL(tt *testing.T) {
	var zeroValue string
	i := &Import{VCSURL: &zeroValue}
	i.GetVCSURL()
	i = &Import{}
	i.GetVCSURL()
	i = nil
	i.GetVCSURL()
}

func TestImport_GetVCSUsername(tt *testing.T) {
	var zeroValue string
	i := &Import{VCSUsername: &zeroValue}
	i.GetVCSUsername()
	i = &Import{}
	i.GetVCSUsername()
	i = nil
	i.GetVCSUsername()
	i := &InstallationRepositoriesEvent{}
	i.GetSender()
	i = nil
	i.GetSender()
}

func TestInstallationToken_GetExpiresAt(tt *testing.T) {
	var zeroValue time.Time
	i := &InstallationToken{ExpiresAt: &zeroValue}
	i.GetExpiresAt()
	i = &InstallationToken{}
	i.GetExpiresAt()
	i = nil
	i.GetExpiresAt()
}

func TestInstallationToken_GetPermissions(tt *testing.T) {
	i := &InstallationToken{}
	i.GetPermissions()
	i = nil
	i.GetPermissions()
}

func TestInstallationToken_GetToken(tt *testing.T) {
	var zeroValue string
	i := &InstallationToken{Token: &zeroValue}
	i.GetToken()
	i = &InstallationToken{}
	i.GetToken()
	i = nil
	i.GetToken()
}

func TestInstallationTokenOptions_GetPermissions(tt *testing.T) {
	i := &InstallationTokenOptions{}
	i.GetPermissions()
	i = nil
	i.GetPermissions()
}

func TestInteractionRestriction_GetExpiresAt(tt *testing.T) {
	var zeroValue Timestamp
	i := &InteractionRestriction{ExpiresAt: &zeroValue}
	i.GetExpiresAt()
	i = &InteractionRestriction{}
	i.GetExpiresAt()
	i = nil
	i.GetExpiresAt()
}

func TestInteractionRestriction_GetLimit(tt *testing.T) {
	j.GetTotalCount()
}

func TestKey_GetCreatedAt(tt *testing.T) {
	var zeroValue Timestamp
	k := &Key{CreatedAt: &zeroValue}
	k.GetCreatedAt()
	k = &Key{}
	k.GetCreatedAt()
	k = nil
	k.GetCreatedAt()
}

func TestKey_GetID(tt *testing.T) {
	var zeroValue int64
	k := &Key{ID: &zeroValue}
	k.GetID()
	k = &Key{}
	k.GetID()
	k = nil
	k.GetID()
}

func TestKey_GetKey(tt *testing.T) {
	var zeroValue string
	k := &Key{Key: &zeroValue}
	k.GetKey()
	k = &Key{}
	k.GetKey()
	k = nil
	k.GetKey()
}

func TestKey_GetReadOnly(tt *testing.T) {
	var zeroValue bool
	k := &Key{ReadOnly: &zeroValue}
	k.GetReadOnly()
	k = &Key{}
	k.GetReadOnly()
	k = nil
	k.GetReadOnly()
}

func TestKey_GetTitle(tt *testing.T) {
	var zeroValue string
	k := &Key{Title: &zeroValue}
	k.GetTitle()
	k = &Key{}
	k.GetTitle()
	k = nil
	k.GetTitle()
	l.GetFeatured()
}

func TestLicense_GetHTMLURL(tt *testing.T) {
	var zeroValue string
	l := &License{HTMLURL: &zeroValue}
	l.GetHTMLURL()
	l = &License{}
	l.GetHTMLURL()
	l = nil
	l.GetHTMLURL()
}

func TestLicense_GetImplementation(tt *testing.T) {
	var zeroValue string
	l := &License{Implementation: &zeroValue}
	l.GetImplementation()
	l = &License{}
	l.GetImplementation()
	l = nil
	l.GetImplementation()
}

func TestLicense_GetKey(tt *testing.T) {
	var zeroValue string
	l := &License{Key: &zeroValue}
	l.GetKey()
	l = &License{}
	l.GetKey()
	l = nil
	l.GetKey()
}

func TestLicense_GetLimitations(tt *testing.T) {
	var zeroValue []string
	l := &License{Limitations: &zeroValue}
	l.GetLimitations()
	l = &License{}
	l.GetLimitations()
	l = nil
	l.GetLimitations()
}

func TestLicense_GetName(tt *testing.T) {
	var zeroValue string
	l := &License{Name: &zeroValue}
	l.GetName()
	l = &License{}
	l.GetName()
	l = nil
	l.GetName()
	m.GetHook()
}

func TestMetaEvent_GetHookID(tt *testing.T) {
	var zeroValue int64
	m := &MetaEvent{HookID: &zeroValue}
	m.GetHookID()
	m = &MetaEvent{}
	m.GetHookID()
	m = nil
	m.GetHookID()
}

func TestMetric_GetHTMLURL(tt *testing.T) {
	var zeroValue string
	m := &Metric{HTMLURL: &zeroValue}
	m.GetHTMLURL()
	m = &Metric{}
	m.GetHTMLURL()
	m = nil
	m.GetHTMLURL()
}

func TestMetric_GetKey(tt *testing.T) {
	var zeroValue string
	m := &Metric{Key: &zeroValue}
	m.GetKey()
	m = &Metric{}
	m.GetKey()
	m = nil
	m.GetKey()
}

func TestMetric_GetName(tt *testing.T) {
	var zeroValue string
	m := &Metric{Name: &zeroValue}
	m.GetName()
	m = &Metric{}
	m.GetName()
	m = nil
	m.GetName()
}

func TestMetric_GetURL(tt *testing.T) {
	var zeroValue string
	m := &Metric{URL: &zeroValue}
	m.GetURL()
	m = &Metric{}
	m.GetURL()
	m = nil
	m.GetURL()
}

func TestPublicEvent_GetInstallation(tt *testing.T) {
	p := &PublicEvent{}
	p.GetInstallation()
	p = nil
	p.GetInstallation()
}

func TestPublicEvent_GetRepo(tt *testing.T) {
	p := &PublicEvent{}
	p.GetRepo()
	p = nil
	p.GetRepo()
}

func TestPublicEvent_GetSender(tt *testing.T) {
	p := &PublicEvent{}
	p.GetSender()
	p = nil
	p.GetSender()
}

func TestPublicKey_GetKey(tt *testing.T) {
	var zeroValue string
	p := &PublicKey{Key: &zeroValue}
	p.GetKey()
	p = &PublicKey{}
	p.GetKey()
	p = nil
	p.GetKey()
}

func TestPublicKey_GetKeyID(tt *testing.T) {
	var zeroValue string
	p := &PublicKey{KeyID: &zeroValue}
	p.GetKeyID()
	p = &PublicKey{}
	p.GetKeyID()
	p = nil
	p.GetKeyID()
}

func TestPullRequest_GetActiveLockReason(tt *testing.T) {
	var zeroValue string
	p := &PullRequest{ActiveLockReason: &zeroValue}
	p.GetActiveLockReason()
	p = &PullRequest{}
	p.GetActiveLockReason()
	p = nil
	p.GetActiveLockReason()
	r.GetRef()
}

func TestReference_GetURL(tt *testing.T) {
	var zeroValue string
	r := &Reference{URL: &zeroValue}
	r.GetURL()
	r = &Reference{}
	r.GetURL()
	r = nil
	r.GetURL()
}

func TestRegistrationToken_GetExpiresAt(tt *testing.T) {
	var zeroValue Timestamp
	r := &RegistrationToken{ExpiresAt: &zeroValue}
	r.GetExpiresAt()
	r = &RegistrationToken{}
	r.GetExpiresAt()
	r = nil
	r.GetExpiresAt()
}

func TestRegistrationToken_GetToken(tt *testing.T) {
	var zeroValue string
	r := &RegistrationToken{Token: &zeroValue}
	r.GetToken()
	r = &RegistrationToken{}
	r.GetToken()
	r = nil
	r.GetToken()
}

func TestReleaseAsset_GetBrowserDownloadURL(tt *testing.T) {
	var zeroValue string
	r := &ReleaseAsset{BrowserDownloadURL: &zeroValue}
	r.GetBrowserDownloadURL()
	r = &ReleaseAsset{}
	r.GetBrowserDownloadURL()
	r = nil
	r.GetBrowserDownloadURL()
}

func TestReleaseAsset_GetContentType(tt *testing.T) {
	var zeroValue string
	r := &ReleaseAsset{ContentType: &zeroValue}
	r.GetContentType()
	r = &ReleaseAsset{}
	r.GetContentType()
	r = nil
	r.GetContentType()
	r := &ReleaseEvent{}
	r.GetRepo()
	r = nil
	r.GetRepo()
}

func TestReleaseEvent_GetSender(tt *testing.T) {
	r := &ReleaseEvent{}
	r.GetSender()
	r = nil
	r.GetSender()
}

func TestRemoveToken_GetExpiresAt(tt *testing.T) {
	var zeroValue Timestamp
	r := &RemoveToken{ExpiresAt: &zeroValue}
	r.GetExpiresAt()
	r = &RemoveToken{}
	r.GetExpiresAt()
	r = nil
	r.GetExpiresAt()
}

func TestRemoveToken_GetToken(tt *testing.T) {
	var zeroValue string
	r := &RemoveToken{Token: &zeroValue}
	r.GetToken()
	r = &RemoveToken{}
	r.GetToken()
	r = nil
	r.GetToken()
}

func TestRename_GetFrom(tt *testing.T) {
	var zeroValue string
	r := &Rename{From: &zeroValue}
	r.GetFrom()
	r = &Rename{}
	r.GetFrom()
	r = nil
	r.GetFrom()
}

func TestRename_GetTo(tt *testing.T) {
	var zeroValue string
	r := &Rename{To: &zeroValue}
	r.GetTo()
	r = &Rename{}
	r.GetTo()
	r = nil
	r.GetTo()
	u.GetID()
}

func TestUserAuthorization_GetNote(tt *testing.T) {
	var zeroValue string
	u := &UserAuthorization{Note: &zeroValue}
	u.GetNote()
	u = &UserAuthorization{}
	u.GetNote()
	u = nil
	u.GetNote()
}

func TestUserAuthorization_GetNoteURL(tt *testing.T) {
	var zeroValue string
	u := &UserAuthorization{NoteURL: &zeroValue}
	u.GetNoteURL()
	u = &UserAuthorization{}
	u.GetNoteURL()
	u = nil
	u.GetNoteURL()
}

func TestUserAuthorization_GetToken(tt *testing.T) {
	var zeroValue string
	u := &UserAuthorization{Token: &zeroValue}
	u.GetToken()
	u = &UserAuthorization{}
	u.GetToken()
	u = nil
	u.GetToken()
}

func TestUserAuthorization_GetTokenLastEight(tt *testing.T) {
	var zeroValue string
	u := &UserAuthorization{TokenLastEight: &zeroValue}
	u.GetTokenLastEight()
	u = &UserAuthorization{}
	u.GetTokenLastEight()
	u = nil
	u.GetTokenLastEight()
}

func TestUserAuthorization_GetUpdatedAt(tt *testing.T) {
	var zeroValue Timestamp
	u := &UserAuthorization{UpdatedAt: &zeroValue}
	u.GetUpdatedAt()
	u = &UserAuthorization{}
	u.GetUpdatedAt()
	u = nil
	u.GetUpdatedAt()
}

func TestWatchEvent_GetInstallation(tt *testing.T) {
	w := &WatchEvent{}
	w.GetInstallation()
	w = nil
	w.GetInstallation()
}

func TestWatchEvent_GetRepo(tt *testing.T) {
	w := &WatchEvent{}
	w.GetRepo()
	w = nil
	w.GetRepo()
}

func TestWatchEvent_GetSender(tt *testing.T) {
	w := &WatchEvent{}
	w.GetSender()
	w = nil
	w.GetSender()
}

func TestWebHookAuthor_GetEmail(tt *testing.T) {
	var zeroValue string
	w := &WebHookAuthor{Email: &zeroValue}
	w.GetEmail()
	w = &WebHookAuthor{}
	w.GetEmail()
	w = nil
	w.GetEmail()
}

func TestWebHookAuthor_GetName(tt *testing.T) {
	var zeroValue string
	w := &WebHookAuthor{Name: &zeroValue}
	w.GetName()
	w = &WebHookAuthor{}
	w.GetName()
	w = nil
	w.GetName()
}

func TestWebHookAuthor_GetUsername(tt *testing.T) {
	var zeroValue string
	w := &WebHookAuthor{Username: &zeroValue}
	w.GetUsername()
	w = &WebHookAuthor{}
	w.GetUsername()
	w = nil
	w.GetUsername()
}

func TestWebHookCommit_GetAuthor(tt *testing.T) {
	w := &WebHookCommit{}
	w.GetAuthor()
	w = nil
	w.GetAuthor()
}

func TestWebHookCommit_GetCommitter(tt *testing.T) {
	w := &WebHookCommit{}
	w.GetCommitter()
	w = nil
	w.GetCommitter()
}

func TestWebHookCommit_GetDistinct(tt *testing.T) {
	var zeroValue bool
	w := &WebHookCommit{Distinct: &zeroValue}
	w.GetDistinct()