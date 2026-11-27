
         Given that the server only sends its version number after the client has already done so, the above
         seems to be suggesting that v3 should be the default version.  This makes sense given that v3 is the
         most popular.

         <http://tools.ietf.org/html/draft-ietf-secsh-filexfer-13#section-5.5> states the following;

         "If the server did not send the "versions" extension, or the version-from-list was not included, the
          server MAY send a status response describing the failure, but MUST then close the channel without
          processing any further requests."

         So what do you do if you have a client whose initial SSH_FXP_INIT packet says it implements v3 and
         a server whose initial SSH_FXP_VERSION reply says it implements v4 and only v4?  If it only implements
         v4, the "versions" extension is likely not going to have been sent so version re-negotiation as discussed
         in draft-ietf-secsh-filexfer-13 would be quite impossible.  As such, what \phpseclib3\Net\SFTP would do is close the
         channel and reopen it with a new and updated SSH_FXP_INIT packet.
        */
        switch ($this->version) {
            case 2:
            case 3:
                break;
            default:
                return false;
        }

        $this->pwd = $this->realpath('.');

        $this->update_stat_cache($this->pwd, []);

        return true;
    }

    /**
     * Disable the stat cache
     *
     * @access public
     */
    public function disableStatCache()
    {
        $this->use_stat_cache = false;
    }

    /**
     * Enable the stat cache
     *
     * @access public
     */
    public function enableStatCache()
    {
        $this->use_stat_cache = true;
    }
    {
        $files = $this->readlist($dir, true);
        if (!$recursive || $files === false) {
            return $files;
        }

        static $depth = 0;

        foreach ($files as $key => $value) {
            if ($depth != 0 && $key == '..') {
                unset($files[$key]);
                continue;
            }
            $is_directory = false;
            if ($key != '.' && $key != '..') {
                if ($this->use_stat_cache) {
                    $is_directory = is_array($this->query_stat_cache($this->realpath($dir . '/' . $key)));
                } else {
                    $stat = $this->lstat($dir . '/' . $key);
                    $is_directory = $stat && $stat['type'] === NET_SFTP_TYPE_DIRECTORY;
                }
            }

            if ($is_directory) {
                $depth++;
                $files[$key] = $this->rawlist($dir . '/' . $key, true);
                $depth--;
            } else {
                $files[$key] = (object) $value;
            }
        }

        return $files;
    }

    /**
     * Reads a list, be it detailed or not, of files in the given directory
     *
     * @param string $dir
     * @param bool $raw
     * @return mixed
     * @throws \UnexpectedValueException on receipt of unexpected packets
     * @access private
     */
    private function readlist($dir, $raw = true)
    {
        if (!($this->bitmap & SSH2::MASK_LOGIN)) {
            return false;
        }

        $dir = $this->realpath($dir . '/');