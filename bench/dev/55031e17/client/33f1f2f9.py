                                       help='String with extra arguments for "git clone"')
        install_subparser.add_argument("-sf", "--source-folder",
                                       help='Install files only from a source subfolder from the '
                                       'specified origin')
        install_subparser.add_argument("-tf", "--target-folder",
                                       help='Install to that path in the conan cache')
        install_subparser.add_argument("-l", "--list", default=False, action='store_true',
                                       help='List stored configuration origins')
        install_subparser.add_argument("-r", "--remove", type=int,
                                       help='Remove configuration origin by index in list (index '
                                            'provided by --list argument)')
        rm_subparser.add_argument("item", help="Item to remove")
        set_subparser.add_argument("item", help="'item=value' to set")
        init_subparser.add_argument('-f', '--force', default=False, action='store_true',
                                    help='Overwrite existing Conan configuration files')

        args = parser.parse_args(*args)

        if args.subcommand == "set":
            try:
                key, value = args.item.split("=", 1)
            except ValueError:
                if "hooks." in args.item:
                    key, value = args.item.split("=", 1)[0], None
                else:
                    raise ConanException("Please specify 'key=value'")
            return self._conan.config_set(key, value)
        elif args.subcommand == "get":
            return self._conan.config_get(args.item)
        elif args.subcommand == "rm":
            return self._conan.config_rm(args.item)
        elif args.subcommand == "home":
            conan_home = self._conan.config_home()
            self._out.info(conan_home)
            if args.json:
                self._outputer.json_output({"home": conan_home}, args.json, os.getcwd())
            return conan_home
        elif args.subcommand == "install":
            if args.list:
                configs = self._conan.config_install_list()
                for index, config in enumerate(configs):
                    self._out.writeln("%s: %s" % (index, config))
                return
            elif args.remove is not None:
                self._conan.config_install_remove(index=args.remove)
                return
            verify_ssl = get_bool_from_text(args.verify_ssl)
            return self._conan.config_install(args.item, verify_ssl, args.type, args.args,
                                              source_folder=args.source_folder,
                                              target_folder=args.target_folder)
        elif args.subcommand == 'init':
        return self._conan.copy(reference=reference, user_channel=args.user_channel,
                                force=args.force, packages=packages_list or args.all)

    def user(self, *args):
        """
        Authenticates against a remote with user/pass, caching the auth token.

        Useful to avoid the user and password being requested later. e.g. while
        you're uploading a package.  You can have one user for each remote.
        Changing the user, or introducing the password is only necessary to
        perform changes in remote packages.
        """
        # FIXME: Difficult and confusing CLI. Better with:
        # - conan user clean -> clean users
        # - conan user list ('remote') -> list users (of a remote)
        # - conan user auth 'remote' ('user') ('password') -> login a remote (w/o user or pass)
        # - conan user set 'user' 'remote' -> set user for a remote (not login) necessary??
        parser = argparse.ArgumentParser(description=self.user.__doc__,
                                         prog="conan user",
                                         formatter_class=SmartFormatter)
        parser.add_argument("name", nargs='?', default=None,
                            help='Username you want to use. If no name is provided it will show the'
                            ' current user')
        parser.add_argument('-c', '--clean', default=False, action='store_true',
                            help='Remove user and tokens for all remotes')
        parser.add_argument("-p", "--password", nargs='?', const="", type=str, action=OnceArgument,
                            help='User password. Use double quotes if password with spacing, '
                                 'and escape quotes if existing. If empty, the password is '
                                 'requested interactively (not exposed)')
        parser.add_argument("-r", "--remote", help='Use the specified remote server',
                            action=OnceArgument)
        parser.add_argument("-j", "--json", default=None, action=OnceArgument,
                            help='json file path where the user list will be written to')
        parser.add_argument("-s", "--skip-auth", default=False, action='store_true',
                            help='Skips the authentication with the server if there are local '
                                 'stored credentials. It doesn\'t check if the '
                                 'current credentials are valid or not')
        args = parser.parse_args(*args)

        if args.clean and any((args.name, args.remote, args.password, args.json, args.skip_auth)):
            raise ConanException("'--clean' argument cannot be used together with 'name', "
                                 "'--password', '--remote', '--json' or '--skip.auth'")
        elif args.json and any((args.name, args.password)):
            raise ConanException("'--json' cannot be used together with 'name' or '--password'")

        cwd = os.getcwd()
        info = None

        try:
            if args.clean:  # clean users
                self._conan.users_clean()
            elif not args.name and args.password is None:  # list users
                info = self._conan.users_list(args.remote)
                self._outputer.print_user_list(info)
            elif args.password is None:  # set user for remote (no password indicated)
                remote_name, prev_user, user = self._conan.user_set(args.name, args.remote)
                self._outputer.print_user_set(remote_name, prev_user, user)
            else:  # login a remote
                remote_name = args.remote or self._conan.get_default_remote().name
                name = args.name
                password = args.password
                remote_name, prev_user, user = self._conan.authenticate(name,
                                                                        remote_name=remote_name,
                                                                        password=password,
                                                                        skip_auth=args.skip_auth)

                self._outputer.print_user_set(remote_name, prev_user, user)
                                                "folder or path to a profile file")

        parser_remove = subparsers.add_parser('remove', help='Remove a profile key')
        parser_remove.add_argument('item', help='key, e.g.: settings.compiler')
        parser_remove.add_argument('profile', help="Name of the profile in the '.conan/profiles' "
                                                   "folder or path to a profile file")

        args = parser.parse_args(*args)

        profile = args.profile if hasattr(args, 'profile') else None

        if args.subcommand == "list":
            profiles = self._conan.profile_list()
            self._outputer.profile_list(profiles)
            if args.json:
                self._outputer.json_output(profiles, args.json, os.getcwd())
        elif args.subcommand == "show":
            profile_text = self._conan.read_profile(profile)
            self._outputer.print_profile(profile, profile_text)
        elif args.subcommand == "new":
            self._conan.create_profile(profile, args.detect, args.force)
        elif args.subcommand == "update":
            try:
                key, value = args.item.split("=", 1)
            except ValueError:
                raise ConanException("Please specify key=value")
            self._conan.update_profile(profile, key, value)
        elif args.subcommand == "get":
            key = args.item
            self._out.writeln(self._conan.get_profile_key(profile, key))
        elif args.subcommand == "remove":
            self._conan.delete_profile_key(profile, args.item)

    def get(self, *args):
        """
        Gets a file or list a directory of a given reference or package.
        """
        parser = argparse.ArgumentParser(description=self.get.__doc__,
                                         prog="conan get",
                                         formatter_class=SmartFormatter)
        parser.add_argument('reference', help=_REF_OR_PREF_HELP)
        parser.add_argument('path',
                            help='Path to the file or directory. If not specified will get the '
                                 'conanfile if only a reference is specified and a conaninfo.txt '
                                 'file contents if the package is also specified',
                            default=None, nargs="?")
        parser.add_argument("-p", "--package", default=None,
                            help="Package ID [DEPRECATED: use full reference instead]",
                            action=OnceArgument)
        parser.add_argument("-r", "--remote", action=OnceArgument,
                            help='Get from this specific remote')