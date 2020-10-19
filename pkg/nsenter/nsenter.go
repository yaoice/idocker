package nsenter

/*
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *idocker_pid;
	idocker_pid = getenv("idocker_pid");
	if (idocker_pid) {
		//fprintf(stdout, "got idocker_pid=%s\n", idocker_pid);
	} else {
		//fprintf(stdout, "missing idocker_pid env skip nsenter");
		return;
	}
	char *idocker_cmd;
	idocker_cmd = getenv("idocker_cmd");
	if (idocker_cmd) {
		//fprintf(stdout, "got idocker_cmd=%s\n", idocker_cmd);
	} else {
		//fprintf(stdout, "missing idocker_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", idocker_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(idocker_cmd);
	exit(0);
	return;
}
*/
import "C"
