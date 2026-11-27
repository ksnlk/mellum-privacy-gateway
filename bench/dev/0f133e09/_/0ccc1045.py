            else:
                post_parameters["assignee"] = assignee._identity
        if assignees is not github.GithubObject.NotSet:
            post_parameters["assignees"] = [
                element._identity
                if isinstance(element, github.NamedUser.NamedUser)
                else element
                for element in assignees
            ]
        if milestone is not github.GithubObject.NotSet:
            post_parameters["milestone"] = milestone._identity
        if labels is not github.GithubObject.NotSet:
            post_parameters["labels"] = [
                element.name if isinstance(element, github.Label.Label) else element
                for element in labels
            ]
        headers, data = self._requester.requestJsonAndCheck(
            "POST", f"{self.url}/issues", input=post_parameters
        )
        return github.Issue.Issue(self._requester, headers, data, completed=True)

    def create_key(self, title, key, read_only=False):
        """
        :calls: `POST /repos/{owner}/{repo}/keys <http://docs.github.com/en/rest/reference/repos#deploy-keys>`_
        :param title: string
        :param key: string
        :param read_only: bool
        :rtype: :class:`github.RepositoryKey.RepositoryKey`
        """
        assert isinstance(title, str), title
        assert isinstance(key, str), key
        assert isinstance(read_only, bool), read_only
        post_parameters = {
            "title": title,
            "key": key,
            "read_only": read_only,
        }
        headers, data = self._requester.requestJsonAndCheck(
            "POST", f"{self.url}/keys", input=post_parameters
        )
        return github.RepositoryKey.RepositoryKey(
            self._requester, headers, data, completed=True
        )

    def create_label(self, name, color, description=github.GithubObject.NotSet):
        """
        :calls: `POST /repos/{owner}/{repo}/labels <http://docs.github.com/en/rest/reference/issues#labels>`_
        :param name: string
        :param color: string
        :param description: string
        :rtype: :class:`github.Label.Label`
        """
        assert isinstance(name, str), name
        assert isinstance(color, str), color
        assert description is github.GithubObject.NotSet or isinstance(
            description, str
        ), description
        post_parameters = {
            "name": name,
            "color": color,

    def delete_secret(self, secret_name):
        """
        :calls: `DELETE /repos/{owner}/{repo}/actions/secrets/{secret_name} <https://docs.github.com/en/rest/reference/actions#delete-a-repository-secret>`_
        :param secret_name: string
        :rtype: bool
        """
        assert isinstance(secret_name, str), secret_name
        status, headers, data = self._requester.requestJson(
            "DELETE", f"{self.url}/actions/secrets/{secret_name}"
        )
        return status == 204

    def create_source_import(
        self,
        vcs,
        vcs_url,
        vcs_username=github.GithubObject.NotSet,
        vcs_password=github.GithubObject.NotSet,
    ):
        """
        :calls: `PUT /repos/{owner}/{repo}/import <https://docs.github.com/en/rest/reference/migration/source_imports#start-an-import>`_
        :param vcs: string
        :param vcs_url: string
        :param vcs_username: string
        :param vcs_password: string
        :rtype: :class:`github.SourceImport.SourceImport`
        """
        assert isinstance(vcs, str), vcs
        assert isinstance(vcs_url, str), vcs_url
        assert vcs_username is github.GithubObject.NotSet or isinstance(
            vcs_username, str
        ), vcs_username
        assert vcs_password is github.GithubObject.NotSet or isinstance(
            vcs_password, str
        ), vcs_password
        put_parameters = {"vcs": vcs, "vcs_url": vcs_url}

        if vcs_username is not github.GithubObject.NotSet:
            put_parameters["vcs_username"] = vcs_username

        if vcs_password is not github.GithubObject.NotSet:
            put_parameters["vcs_password"] = vcs_password

        import_header = {"Accept": Consts.mediaTypeImportPreview}

        headers, data = self._requester.requestJsonAndCheck(
            "PUT", f"{self.url}/import", headers=import_header, input=put_parameters
        )

        return github.SourceImport.SourceImport(