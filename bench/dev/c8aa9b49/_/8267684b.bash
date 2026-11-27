fi

__gitcomp_builtin_add_default=" --dry-run --verbose --interactive --patch --edit --force --update --renormalize --intent-to-add --all --ignore-removal --refresh --ignore-errors --ignore-missing --chmod= --pathspec-from-file= --pathspec-file-nul --no-dry-run -- --no-verbose --no-interactive --no-patch --no-edit --no-force --no-update --no-renormalize --no-intent-to-add --no-all --no-ignore-removal --no-refresh --no-ignore-errors --no-ignore-missing --no-chmod --no-pathspec-from-file --no-pathspec-file-nul"
__gitcomp_builtin_am_default=" --interactive --3way --quiet --signoff --utf8 --keep --keep-non-patch --message-id --keep-cr --no-keep-cr --scissors --whitespace= --ignore-space-change --ignore-whitespace --directory= --exclude= --include= --patch-format= --reject --resolvemsg= --continue --resolved --skip --abort --quit --show-current-patch --committer-date-is-author-date --ignore-date --rerere-autoupdate --gpg-sign -- --no-interactive --no-3way --no-quiet --no-signoff --no-utf8 --no-keep --no-keep-non-patch --no-message-id --no-scissors --no-whitespace --no-ignore-space-change --no-ignore-whitespace --no-directory --no-exclude --no-include --no-patch-format --no-reject --no-resolvemsg --no-committer-date-is-author-date --no-ignore-date --no-rerere-autoupdate --no-gpg-sign"
__gitcomp_builtin_apply_default=" --exclude= --include= --no-add --stat --numstat --summary --check --index --intent-to-add --cached --apply --3way --build-fake-ancestor= --whitespace= --ignore-space-change --ignore-whitespace --reverse --unidiff-zero --reject --allow-overlap --verbose --inaccurate-eof --recount --directory= --add -- --no-stat --no-numstat --no-summary --no-check --no-index --no-intent-to-add --no-cached --no-apply --no-3way --no-build-fake-ancestor --no-whitespace --no-ignore-space-change --no-ignore-whitespace --no-reverse --no-unidiff-zero --no-reject --no-allow-overlap --no-verbose --no-inaccurate-eof --no-recount --no-directory"
__gitcomp_builtin_archive_default=" --output= --remote= --exec= --no-output -- --no-remote --no-exec"
__gitcomp_builtin_bisect__helper_default=" --next-all --write-terms --bisect-clean-state --check-expected-revs --bisect-reset --bisect-write --check-and-set-terms --bisect-next-check --bisect-terms --bisect-start --bisect-next --bisect-auto-next --bisect-autostart --no-log --log"
__gitcomp_builtin_blame_default=" --incremental --root --show-stats --progress --score-debug --show-name --show-number --porcelain --line-porcelain --show-email --ignore-rev= --ignore-revs-file= --color-lines --color-by-age --minimal --contents= --abbrev --no-incremental -- --no-root --no-show-stats --no-progress --no-score-debug --no-show-name --no-show-number --no-porcelain --no-line-porcelain --no-show-email --no-ignore-rev --no-ignore-revs-file --no-color-lines --no-color-by-age --no-minimal --no-contents --no-abbrev"
__gitcomp_builtin_branch_default=" --verbose --quiet --track --set-upstream-to= --unset-upstream --color --remotes --contains --no-contains --abbrev --all --delete --move --copy --list --show-current --create-reflog --edit-description --merged --no-merged --column --sort= --points-at= --ignore-case --format= -- --no-verbose --no-quiet --no-track --no-set-upstream-to --no-unset-upstream --no-color --no-remotes --no-abbrev --no-all --no-delete --no-move --no-copy --no-list --no-show-current --no-create-reflog --no-edit-description --no-column --no-points-at --no-ignore-case --no-format"
__gitcomp_builtin_bugreport_default=" --output-directory= --suffix= --no-output-directory -- --no-suffix"
__gitcomp_builtin_cat_file_default=" --textconv --filters --path= --allow-unknown-type --buffer --batch --batch-check --follow-symlinks --batch-all-objects --unordered --no-path -- --no-allow-unknown-type --no-buffer --no-follow-symlinks --no-batch-all-objects --no-unordered"
__gitcomp_builtin_check_attr_default=" --all --cached --stdin --no-all -- --no-cached --no-stdin"
__gitcomp_builtin_check_ignore_default=" --quiet --verbose --stdin --non-matching --no-index --index -- --no-quiet --no-verbose --no-stdin --no-non-matching"
__gitcomp_builtin_check_mailmap_default=" --stdin --no-stdin"
__gitcomp_builtin_checkout_default=" --guess --overlay --quiet --recurse-submodules --progress --merge --conflict= --detach --track --orphan= --ignore-other-worktrees --ours --theirs --patch --ignore-skip-worktree-bits --pathspec-from-file= --pathspec-file-nul --no-guess -- --no-overlay --no-quiet --no-recurse-submodules --no-progress --no-merge --no-conflict --no-detach --no-track --no-orphan --no-ignore-other-worktrees --no-patch --no-ignore-skip-worktree-bits --no-pathspec-from-file --no-pathspec-file-nul"
__gitcomp_builtin_checkout_index_default=" --all --force --quiet --no-create --index --stdin --temp --prefix= --stage= --create -- --no-all --no-force --no-quiet --no-index --no-stdin --no-temp --no-prefix"
__gitcomp_builtin_cherry_default=" --abbrev --verbose --no-abbrev -- --no-verbose"
__gitcomp_builtin_cherry_pick_default=" --quit --continue --abort --skip --cleanup= --no-commit --edit --signoff --mainline= --rerere-autoupdate --strategy= --strategy-option= --gpg-sign --ff --allow-empty --allow-empty-message --keep-redundant-commits --commit -- --no-cleanup --no-edit --no-signoff --no-mainline --no-rerere-autoupdate --no-strategy --no-strategy-option --no-gpg-sign --no-ff --no-allow-empty --no-allow-empty-message --no-keep-redundant-commits"
__gitcomp_builtin_clean_default=" --quiet --dry-run --interactive --exclude= --no-quiet -- --no-dry-run --no-interactive"
__gitcomp_builtin_clone_default=" --verbose --quiet --progress --no-checkout --bare --mirror --local --no-hardlinks --shared --recurse-submodules --recursive --jobs= --template= --reference= --reference-if-able= --dissociate --origin= --branch= --upload-pack= --depth= --shallow-since= --shallow-exclude= --single-branch --no-tags --shallow-submodules --separate-git-dir= --config= --server-option= --ipv4 --ipv6 --filter= --remote-submodules --sparse --checkout --hardlinks --tags -- --no-verbose --no-quiet --no-progress --no-bare --no-mirror --no-local --no-shared --no-recurse-submodules --no-recursive --no-jobs --no-template --no-reference --no-reference-if-able --no-dissociate --no-origin --no-branch --no-upload-pack --no-depth --no-shallow-since --no-shallow-exclude --no-single-branch --no-shallow-submodules --no-separate-git-dir --no-config --no-server-option --no-ipv4 --no-ipv6 --no-filter --no-remote-submodules --no-sparse"
__gitcomp_builtin_column_default=" --command= --mode --raw-mode= --width= --indent= --nl= --padding= --no-command -- --no-mode --no-raw-mode --no-width --no-indent --no-nl --no-padding"
__gitcomp_builtin_commit_default=" --quiet --verbose --file= --author= --date= --message= --reedit-message= --reuse-message= --fixup= --squash= --reset-author --signoff --template= --edit --cleanup= --status --gpg-sign --all --include --interactive --patch --only --no-verify --dry-run --short --branch --ahead-behind --porcelain --long --null --amend --no-post-rewrite --untracked-files --pathspec-from-file= --pathspec-file-nul --verify --post-rewrite -- --no-quiet --no-verbose --no-file --no-author --no-date --no-message --no-reedit-message --no-reuse-message --no-fixup --no-squash --no-reset-author --no-signoff --no-template --no-edit --no-cleanup --no-status --no-gpg-sign --no-all --no-include --no-interactive --no-patch --no-only --no-dry-run --no-short --no-branch --no-ahead-behind --no-porcelain --no-long --no-null --no-amend --no-untracked-files --no-pathspec-from-file --no-pathspec-file-nul"
__gitcomp_builtin_commit_graph_default=" --object-dir= --no-object-dir"
__gitcomp_builtin_config_default=" --global --system --local --worktree --file= --blob= --get --get-all --get-regexp --get-urlmatch --replace-all --add --unset --unset-all --rename-section --remove-section --list --edit --get-color --get-colorbool --type= --bool --int --bool-or-int --bool-or-str --path --expiry-date --null --name-only --includes --show-origin --show-scope --default= --no-global -- --no-system --no-local --no-worktree --no-file --no-blob --no-get --no-get-all --no-get-regexp --no-get-urlmatch --no-replace-all --no-add --no-unset --no-unset-all --no-rename-section --no-remove-section --no-list --no-edit --no-get-color --no-get-colorbool --no-type --no-null --no-name-only --no-includes --no-show-origin --no-show-scope --no-default"
__gitcomp_builtin_count_objects_default=" --verbose --human-readable --no-verbose -- --no-human-readable"
__gitcomp_builtin_credential_cache_default=" --timeout= --socket= --no-timeout -- --no-socket"
__gitcomp_builtin_credential_cache__daemon_default=" --debug --no-debug"
__gitcomp_builtin_credential_store_default=" --file= --no-file"
__gitcomp_builtin_describe_default=" --contains --debug --all --tags --long --first-parent --abbrev --exact-match --candidates= --match= --exclude= --always --dirty --broken --no-contains -- --no-debug --no-all --no-tags --no-long --no-first-parent --no-abbrev --no-exact-match --no-candidates --no-match --no-exclude --no-always --no-dirty --no-broken"
__gitcomp_builtin_difftool_default=" --gui --dir-diff --no-prompt --symlinks --tool= --tool-help --trust-exit-code --extcmd= --no-index -- --no-gui --no-dir-diff --no-symlinks --no-tool --no-tool-help --no-trust-exit-code --no-extcmd"
__gitcomp_builtin_env__helper_default=" --type= --default= --exit-code --no-default -- --no-exit-code"
__gitcomp_builtin_fast_export_default=" --progress= --signed-tags= --tag-of-filtered-object= --reencode= --export-marks= --import-marks= --import-marks-if-exists= --fake-missing-tagger --full-tree --use-done-feature --no-data --refspec= --anonymize --anonymize-map= --reference-excluded-parents --show-original-ids --mark-tags --data -- --no-progress --no-signed-tags --no-tag-of-filtered-object --no-reencode --no-export-marks --no-import-marks --no-import-marks-if-exists --no-fake-missing-tagger --no-full-tree --no-use-done-feature --no-refspec --no-anonymize --no-reference-excluded-parents --no-show-original-ids --no-mark-tags"
__gitcomp_builtin_fetch_default=" --verbose --quiet --all --set-upstream --append --upload-pack= --force --multiple --tags --jobs= --prune --prune-tags --recurse-submodules --dry-run --write-fetch-head --keep --update-head-ok --progress --depth= --shallow-since= --shallow-exclude= --deepen= --unshallow --update-shallow --refmap= --server-option= --ipv4 --ipv6 --negotiation-tip= --filter= --auto-maintenance --auto-gc --show-forced-updates --write-commit-graph --stdin --no-verbose -- --no-quiet --no-all --no-set-upstream --no-append --no-upload-pack --no-force --no-multiple --no-tags --no-jobs --no-prune --no-prune-tags --no-recurse-submodules --no-dry-run --no-write-fetch-head --no-keep --no-update-head-ok --no-progress --no-depth --no-shallow-since --no-shallow-exclude --no-deepen --no-update-shallow --no-server-option --no-ipv4 --no-ipv6 --no-negotiation-tip --no-filter --no-auto-maintenance --no-auto-gc --no-show-forced-updates --no-write-commit-graph --no-stdin"
__gitcomp_builtin_fmt_merge_msg_default=" --log --message= --file= --no-log -- --no-message --no-file"
__gitcomp_builtin_for_each_ref_default=" --shell --perl --python --tcl --count= --format= --color --sort= --points-at= --merged --no-merged --contains --no-contains --ignore-case -- --no-shell --no-perl --no-python --no-tcl --no-count --no-format --no-color --no-points-at --no-ignore-case"
__gitcomp_builtin_format_patch_default=" --numbered --no-numbered --signoff --stdout --cover-letter --numbered-files --suffix= --start-number= --reroll-count= --rfc --cover-from-description= --subject-prefix= --output-directory= --keep-subject --no-binary --zero-commit --ignore-if-in-upstream --no-stat --add-header= --to= --cc= --from --in-reply-to= --attach --inline --thread --signature= --base= --signature-file= --quiet --progress --interdiff= --range-diff= --creation-factor= --binary -- --no-numbered --no-signoff --no-stdout --no-cover-letter --no-numbered-files --no-suffix --no-start-number --no-reroll-count --no-cover-from-description --no-zero-commit --no-ignore-if-in-upstream --no-add-header --no-to --no-cc --no-from --no-in-reply-to --no-attach --no-thread --no-signature --no-base --no-signature-file --no-quiet --no-progress --no-interdiff --no-range-diff --no-creation-factor"
__gitcomp_builtin_fsck_default=" --verbose --unreachable --dangling --tags --root --cache --reflogs --full --connectivity-only --strict --lost-found --progress --name-objects --no-verbose -- --no-unreachable --no-dangling --no-tags --no-root --no-cache --no-reflogs --no-full --no-connectivity-only --no-strict --no-lost-found --no-progress --no-name-objects"
__gitcomp_builtin_fsck_objects_default=" --verbose --unreachable --dangling --tags --root --cache --reflogs --full --connectivity-only --strict --lost-found --progress --name-objects --no-verbose -- --no-unreachable --no-dangling --no-tags --no-root --no-cache --no-reflogs --no-full --no-connectivity-only --no-strict --no-lost-found --no-progress --no-name-objects"
__gitcomp_builtin_gc_default=" --quiet --prune --aggressive --keep-largest-pack --no-quiet -- --no-prune --no-aggressive --no-keep-largest-pack"
__gitcomp_builtin_grep_default=" --cached --no-index --untracked --exclude-standard --recurse-submodules --invert-match --ignore-case --word-regexp --text --textconv --recursive --max-depth= --extended-regexp --basic-regexp --fixed-strings --perl-regexp --line-number --column --full-name --files-with-matches --name-only --files-without-match --only-matching --count --color --break --heading --context= --before-context= --after-context= --threads= --show-function --function-context --and --or --not --quiet --all-match --index -- --no-cached --no-untracked --no-exclude-standard --no-recurse-submodules --no-invert-match --no-ignore-case --no-word-regexp --no-text --no-textconv --no-recursive --no-extended-regexp --no-basic-regexp --no-fixed-strings --no-perl-regexp --no-line-number --no-column --no-full-name --no-files-with-matches --no-name-only --no-files-without-match --no-only-matching --no-count --no-color --no-break --no-heading --no-context --no-before-context --no-after-context --no-threads --no-show-function --no-function-context --no-or --no-quiet --no-all-match"
__gitcomp_builtin_hash_object_default=" --stdin --stdin-paths --no-filters --literally --path= --filters -- --no-stdin --no-stdin-paths --no-literally --no-path"
__gitcomp_builtin_help_default=" --all --guides --config --man --web --info --verbose --no-all -- --no-guides --no-config --no-man --no-web --no-info --no-verbose"
__gitcomp_builtin_init_default=" --template= --bare --shared --quiet --separate-git-dir= --initial-branch= --object-format= --no-template -- --no-bare --no-quiet --no-separate-git-dir --no-initial-branch --no-object-format"
__gitcomp_builtin_init_db_default=" --template= --bare --shared --quiet --separate-git-dir= --initial-branch= --object-format= --no-template -- --no-bare --no-quiet --no-separate-git-dir --no-initial-branch --no-object-format"
__gitcomp_builtin_interpret_trailers_default=" --in-place --trim-empty --where= --if-exists= --if-missing= --only-trailers --only-input --unfold --parse --no-divider --trailer= --divider -- --no-in-place --no-trim-empty --no-where --no-if-exists --no-if-missing --no-only-trailers --no-only-input --no-unfold --no-trailer"
__gitcomp_builtin_log_default=" --quiet --source --use-mailmap --mailmap --decorate-refs= --decorate-refs-exclude= --decorate --no-quiet -- --no-source --no-use-mailmap --no-mailmap --no-decorate-refs --no-decorate-refs-exclude --no-decorate"
__gitcomp_builtin_ls_files_default=" --cached --deleted --modified --others --ignored --stage --killed --directory --eol --empty-directory --unmerged --resolve-undo --exclude= --exclude-from= --exclude-per-directory= --exclude-standard --full-name --recurse-submodules --error-unmatch --with-tree= --abbrev --debug --no-cached -- --no-deleted --no-modified --no-others --no-ignored --no-stage --no-killed --no-directory --no-eol --no-empty-directory --no-unmerged --no-resolve-undo --no-exclude-per-directory --no-recurse-submodules --no-error-unmatch --no-with-tree --no-abbrev --no-debug"
__gitcomp_builtin_ls_remote_default=" --quiet --upload-pack= --tags --heads --refs --get-url --sort= --symref --server-option= --no-quiet -- --no-upload-pack --no-tags --no-heads --no-refs --no-get-url --no-symref --no-server-option"
__gitcomp_builtin_ls_tree_default=" --long --name-only --name-status --full-name --full-tree --abbrev --no-long -- --no-name-only --no-name-status --no-full-name --no-full-tree --no-abbrev"
__gitcomp_builtin_merge_default=" --stat --summary --log --squash --commit --edit --cleanup= --ff --ff-only --rerere-autoupdate --verify-signatures --strategy= --strategy-option= --message= --file --verbose --quiet --abort --quit --continue --allow-unrelated-histories --progress --gpg-sign --autostash --overwrite-ignore --signoff --no-verify --verify -- --no-stat --no-summary --no-log --no-squash --no-commit --no-edit --no-cleanup --no-ff --no-rerere-autoupdate --no-verify-signatures --no-strategy --no-strategy-option --no-message --no-verbose --no-quiet --no-abort --no-quit --no-continue --no-allow-unrelated-histories --no-progress --no-gpg-sign --no-autostash --no-overwrite-ignore --no-signoff"
__gitcomp_builtin_merge_base_default=" --all --octopus --independent --is-ancestor --fork-point --no-all"
__gitcomp_builtin_merge_file_default=" --stdout --diff3 --ours --theirs --union --marker-size= --quiet --no-stdout -- --no-diff3 --no-ours --no-theirs --no-union --no-marker-size --no-quiet"

	# use a hack to enable file mode in bash < 4
	compopt -o filenames +o nospace 2>/dev/null ||
	compgen -f /non-existing-dir/ >/dev/null ||
	true
}

# Execute 'git ls-files', unless the --committable option is specified, in
# which case it runs 'git diff-index' to find out the files that can be
# committed.  It return paths relative to the directory specified in the first
# argument, and using the options specified in the second argument.
__git_ls_files_helper ()
{
	if [ "$2" == "--committable" ]; then
		__git -C "$1" -c core.quotePath=false diff-index \
			--name-only --relative HEAD -- "${3//\\/\\\\}*"
	else
		# NOTE: $2 is not quoted in order to support multiple options
		__git -C "$1" -c core.quotePath=false ls-files \
			--exclude-standard $2 -- "${3//\\/\\\\}*"
	fi
}


# __git_index_files accepts 1 or 2 arguments:
# 1: Options to pass to ls-files (required).
# 2: A directory path (optional).
#    If provided, only files within the specified directory are listed.
#    Sub directories are never recursed.  Path must have a trailing
#    slash.
# 3: List only paths matching this path component (optional).
__git_index_files ()
{
	local root="$2" match="$3"

	__git_ls_files_helper "$root" "$1" "${match:-?}" |
	awk -F / -v pfx="${2//\\/\\\\}" '{
		paths[$1] = 1
	}
	END {
		for (p in paths) {
			if (substr(p, 1, 1) != "\"") {
				# No special characters, easy!
				print pfx p
				continue
			}

			# The path is quoted.
			p = dequote(p)
			if (p == "")
				continue
		__gitcomp "--default --branch"
		;;
	summary,--*)
		__gitcomp "--cached --files --summary-limit"
		;;
	foreach,--*|sync,--*)
		__gitcomp "--recursive"
		;;
	*)
		;;
	esac
}

_git_svn ()
{
	local subcommands="
		init fetch clone rebase dcommit log find-rev
		set-tree commit-diff info create-ignore propget
		proplist show-ignore show-externals branch tag blame
		migrate mkdirs reset gc
		"
	local subcommand="$(__git_find_on_cmdline "$subcommands")"
	if [ -z "$subcommand" ]; then
		__gitcomp "$subcommands"
	else
		local remote_opts="--username= --config-dir= --no-auth-cache"
		local fc_opts="
			--follow-parent --authors-file= --repack=
			--no-metadata --use-svm-props --use-svnsync-props
			--log-window-size= --no-checkout --quiet
			--repack-flags --use-log-author --localtime
			--add-author-from
			--recursive
			--ignore-paths= --include-paths= $remote_opts
			"
		local init_opts="
			--template= --shared= --trunk= --tags=
			--branches= --stdlayout --minimize-url
			--no-metadata --use-svm-props --use-svnsync-props
			--rewrite-root= --prefix= $remote_opts
			"
		local cmt_opts="
			--edit --rmdir --find-copies-harder --copy-similarity=
			"

		case "$subcommand,$cur" in
		fetch,--*)
			__gitcomp "--revision= --fetch-all $fc_opts"
			;;
		clone,--*)
			__gitcomp "--revision= $fc_opts $init_opts"