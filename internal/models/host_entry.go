package models

type HostEntry struct {
    Name         string
    Hostname     string
    User         string
    IdentityFile string
    Port         int
}

func NewHostEntry(name, hostname, user, identityFile string, port int) *HostEntry {
    return &HostEntry{
        Name:         name,
        Hostname:     hostname,
        User:         user,
        IdentityFile: identityFile,
        Port:         port,
    }
}