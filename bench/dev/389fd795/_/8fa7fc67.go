func (a *Alert) GetRuleSeverity() string {
	if a == nil || a.RuleSeverity == nil {
		return ""
	}
	return *a.RuleSeverity
}

// GetTool returns the Tool field if it's non-nil, zero value otherwise.
func (a *Alert) GetTool() string {
	if a == nil || a.Tool == nil {
		return ""
	}
	return *a.Tool
}

// GetURL returns the URL field if it's non-nil, zero value otherwise.
func (a *Alert) GetURL() string {
	if a == nil || a.URL == nil {
		return ""
	}
	return *a.URL
}

// GetVerifiablePasswordAuthentication returns the VerifiablePasswordAuthentication field if it's non-nil, zero value otherwise.
func (a *APIMeta) GetVerifiablePasswordAuthentication() bool {
	if a == nil || a.VerifiablePasswordAuthentication == nil {
		return false
	}
	return *a.VerifiablePasswordAuthentication
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (a *App) GetCreatedAt() Timestamp {
	if a == nil || a.CreatedAt == nil {
		return Timestamp{}
	}
	return *a.CreatedAt
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (a *App) GetDescription() string {
	if a == nil || a.Description == nil {
		return ""
	}
	return *a.Description
}

// GetExternalURL returns the ExternalURL field if it's non-nil, zero value otherwise.
func (a *App) GetExternalURL() string {
	if a == nil || a.ExternalURL == nil {
		return ""
func (a *App) GetSlug() string {
	if a == nil || a.Slug == nil {
		return ""
	}
	return *a.Slug
}

// GetUpdatedAt returns the UpdatedAt field if it's non-nil, zero value otherwise.
func (a *App) GetUpdatedAt() Timestamp {
	if a == nil || a.UpdatedAt == nil {
		return Timestamp{}
	}
	return *a.UpdatedAt
}

// GetClientID returns the ClientID field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetClientID() string {
	if a == nil || a.ClientID == nil {
		return ""
	}
	return *a.ClientID
}

// GetClientSecret returns the ClientSecret field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetClientSecret() string {
	if a == nil || a.ClientSecret == nil {
		return ""
	}
	return *a.ClientSecret
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetCreatedAt() Timestamp {
	if a == nil || a.CreatedAt == nil {
		return Timestamp{}
	}
	return *a.CreatedAt
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetDescription() string {
	if a == nil || a.Description == nil {
		return ""
	}
	return *a.Description
}

// GetExternalURL returns the ExternalURL field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetExternalURL() string {
	if a == nil || a.ExternalURL == nil {
		return ""
func (a *AppConfig) GetOwner() *User {
	if a == nil {
		return nil
	}
	return a.Owner
}

// GetPEM returns the PEM field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetPEM() string {
	if a == nil || a.PEM == nil {
		return ""
	}
	return *a.PEM
}

// GetUpdatedAt returns the UpdatedAt field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetUpdatedAt() Timestamp {
	if a == nil || a.UpdatedAt == nil {
		return Timestamp{}
	}
	return *a.UpdatedAt
}

// GetWebhookSecret returns the WebhookSecret field if it's non-nil, zero value otherwise.
func (a *AppConfig) GetWebhookSecret() string {
	if a == nil || a.WebhookSecret == nil {
		return ""
	}
	return *a.WebhookSecret
}

// GetArchiveDownloadURL returns the ArchiveDownloadURL field if it's non-nil, zero value otherwise.
func (a *Artifact) GetArchiveDownloadURL() string {
	if a == nil || a.ArchiveDownloadURL == nil {
		return ""
	}
	return *a.ArchiveDownloadURL
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (a *Artifact) GetCreatedAt() Timestamp {
	if a == nil || a.CreatedAt == nil {
		return Timestamp{}
	}
	return *a.CreatedAt
}

// GetExpired returns the Expired field if it's non-nil, zero value otherwise.
func (a *Artifact) GetExpired() bool {
	if a == nil || a.Expired == nil {
		return false
func (a *AuthorizationApp) GetName() string {
	if a == nil || a.Name == nil {
		return ""
	}
	return *a.Name
}

// GetURL returns the URL field if it's non-nil, zero value otherwise.
func (a *AuthorizationApp) GetURL() string {
	if a == nil || a.URL == nil {
		return ""
	}
	return *a.URL
}

// GetClientID returns the ClientID field if it's non-nil, zero value otherwise.
func (a *AuthorizationRequest) GetClientID() string {
	if a == nil || a.ClientID == nil {
		return ""
	}
	return *a.ClientID
}

// GetClientSecret returns the ClientSecret field if it's non-nil, zero value otherwise.
func (a *AuthorizationRequest) GetClientSecret() string {
	if a == nil || a.ClientSecret == nil {
		return ""
	}
	return *a.ClientSecret
}

// GetFingerprint returns the Fingerprint field if it's non-nil, zero value otherwise.
func (a *AuthorizationRequest) GetFingerprint() string {
	if a == nil || a.Fingerprint == nil {
		return ""
	}
	return *a.Fingerprint
}

// GetNote returns the Note field if it's non-nil, zero value otherwise.
func (a *AuthorizationRequest) GetNote() string {
	if a == nil || a.Note == nil {
		return ""
	}
	return *a.Note
}

// GetNoteURL returns the NoteURL field if it's non-nil, zero value otherwise.
func (a *AuthorizationRequest) GetNoteURL() string {
	if a == nil || a.NoteURL == nil {
		return ""
func (i *Import) GetURL() string {
	if i == nil || i.URL == nil {
		return ""
	}
	return *i.URL
}

// GetUseLFS returns the UseLFS field if it's non-nil, zero value otherwise.
func (i *Import) GetUseLFS() string {
	if i == nil || i.UseLFS == nil {
		return ""
	}
	return *i.UseLFS
}

// GetVCS returns the VCS field if it's non-nil, zero value otherwise.
func (i *Import) GetVCS() string {
	if i == nil || i.VCS == nil {
		return ""
	}
	return *i.VCS
}

// GetVCSPassword returns the VCSPassword field if it's non-nil, zero value otherwise.
func (i *Import) GetVCSPassword() string {
	if i == nil || i.VCSPassword == nil {
		return ""
	}
	return *i.VCSPassword
}

// GetVCSURL returns the VCSURL field if it's non-nil, zero value otherwise.
func (i *Import) GetVCSURL() string {
	if i == nil || i.VCSURL == nil {
		return ""
	}
	return *i.VCSURL
}

// GetVCSUsername returns the VCSUsername field if it's non-nil, zero value otherwise.
func (i *Import) GetVCSUsername() string {
	if i == nil || i.VCSUsername == nil {
		return ""
	}
	return *i.VCSUsername
}

// GetAccessTokensURL returns the AccessTokensURL field if it's non-nil, zero value otherwise.
func (i *Installation) GetAccessTokensURL() string {
	if i == nil || i.AccessTokensURL == nil {
		return ""
func (i *InstallationPermissions) GetOrganizationPlan() string {
	if i == nil || i.OrganizationPlan == nil {
		return ""
	}
	return *i.OrganizationPlan
}

// GetOrganizationPreReceiveHooks returns the OrganizationPreReceiveHooks field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetOrganizationPreReceiveHooks() string {
	if i == nil || i.OrganizationPreReceiveHooks == nil {
		return ""
	}
	return *i.OrganizationPreReceiveHooks
}

// GetOrganizationProjects returns the OrganizationProjects field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetOrganizationProjects() string {
	if i == nil || i.OrganizationProjects == nil {
		return ""
	}
	return *i.OrganizationProjects
}

// GetOrganizationSecrets returns the OrganizationSecrets field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetOrganizationSecrets() string {
	if i == nil || i.OrganizationSecrets == nil {
		return ""
	}
	return *i.OrganizationSecrets
}

// GetOrganizationSelfHostedRunners returns the OrganizationSelfHostedRunners field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetOrganizationSelfHostedRunners() string {
	if i == nil || i.OrganizationSelfHostedRunners == nil {
		return ""
	}
	return *i.OrganizationSelfHostedRunners
}

// GetOrganizationUserBlocking returns the OrganizationUserBlocking field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetOrganizationUserBlocking() string {
	if i == nil || i.OrganizationUserBlocking == nil {
		return ""
	}
	return *i.OrganizationUserBlocking
}

// GetPackages returns the Packages field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetPackages() string {
	if i == nil || i.Packages == nil {
		return ""
func (i *InstallationPermissions) GetRepositoryHooks() string {
	if i == nil || i.RepositoryHooks == nil {
		return ""
	}
	return *i.RepositoryHooks
}

// GetRepositoryPreReceiveHooks returns the RepositoryPreReceiveHooks field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetRepositoryPreReceiveHooks() string {
	if i == nil || i.RepositoryPreReceiveHooks == nil {
		return ""
	}
	return *i.RepositoryPreReceiveHooks
}

// GetRepositoryProjects returns the RepositoryProjects field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetRepositoryProjects() string {
	if i == nil || i.RepositoryProjects == nil {
		return ""
	}
	return *i.RepositoryProjects
}

// GetSecrets returns the Secrets field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetSecrets() string {
	if i == nil || i.Secrets == nil {
		return ""
	}
	return *i.Secrets
}

// GetSecretScanningAlerts returns the SecretScanningAlerts field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetSecretScanningAlerts() string {
	if i == nil || i.SecretScanningAlerts == nil {
		return ""
	}
	return *i.SecretScanningAlerts
}

// GetSecurityEvents returns the SecurityEvents field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetSecurityEvents() string {
	if i == nil || i.SecurityEvents == nil {
		return ""
	}
	return *i.SecurityEvents
}

// GetSingleFile returns the SingleFile field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetSingleFile() string {
	if i == nil || i.SingleFile == nil {
		return ""
	}
	return *i.SingleFile
}

// GetStatuses returns the Statuses field if it's non-nil, zero value otherwise.
func (i *InstallationPermissions) GetStatuses() string {
	if i == nil || i.Statuses == nil {
		return ""