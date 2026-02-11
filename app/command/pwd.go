package command

func PwdCmd(s Shell, args []string) (stdout []string, stderr []string) {
	if len(args) != 0 {
		return stdout, []string{("pwd: too many arguments")}
	}
	return []string{s.WorkingDir()}, stderr
}
